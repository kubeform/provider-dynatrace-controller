/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	dynatrace "github.com/dynatrace/terraform-provider-dynatrace/dynatrace"
	"github.com/gobuffalo/flect"
	auditlib "go.bytebuilders.dev/audit/lib"
	arv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	dnsv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/dns/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/firewall/v1alpha1"
	instancev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/instance/v1alpha1"
	kubernetesv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/kubernetes/v1alpha1"
	loadbalancerv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/loadbalancer/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/network/v1alpha1"
	snapshotv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/snapshot/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/ssh/v1alpha1"
	templatev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/template/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/volume/v1alpha1"
	controllersdns "kubeform.dev/provider-dynatrace-controller/controllers/dns"
	controllersfirewall "kubeform.dev/provider-dynatrace-controller/controllers/firewall"
	controllersinstance "kubeform.dev/provider-dynatrace-controller/controllers/instance"
	controllerskubernetes "kubeform.dev/provider-dynatrace-controller/controllers/kubernetes"
	controllersloadbalancer "kubeform.dev/provider-dynatrace-controller/controllers/loadbalancer"
	controllersnetwork "kubeform.dev/provider-dynatrace-controller/controllers/network"
	controllerssnapshot "kubeform.dev/provider-dynatrace-controller/controllers/snapshot"
	controllersssh "kubeform.dev/provider-dynatrace-controller/controllers/ssh"
	controllerstemplate "kubeform.dev/provider-dynatrace-controller/controllers/template"
	controllersvolume "kubeform.dev/provider-dynatrace-controller/controllers/volume"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var runningControllers = struct {
	sync.RWMutex
	mp map[schema.GroupVersionKind]bool
}{mp: make(map[schema.GroupVersionKind]bool)}

func watchCRD(ctx context.Context, crdClient *clientset.Clientset, vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, stopCh <-chan struct{}, mgr manager.Manager, auditor *auditlib.EventPublisher, watchOnlyDefault bool) error {
	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				klog.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				klog.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				klog.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "dynatrace.kubeform.com") {
				gvk := schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: crd.Spec.Versions[0].Name,
					Kind:    crd.Spec.Names.Kind,
				}

				// check whether this gvk came before, if no then start the controller
				runningControllers.RLock()
				_, ok := runningControllers.mp[gvk]
				runningControllers.RUnlock()

				if !ok {
					runningControllers.Lock()
					runningControllers.mp[gvk] = true
					runningControllers.Unlock()

					if enableValidatingWebhook {
						// add dynamic ValidatingWebhookConfiguration

						// create empty VWC if the group has come for the first time
						err := createEmptyVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						// update
						err = updateVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						err = SetupWebhook(mgr, gvk)
						if err != nil {
							setupLog.Error(err, "unable to enable webhook")
							os.Exit(1)
						}
					}

					err = SetupManager(ctx, mgr, gvk, auditor, watchOnlyDefault)
					if err != nil {
						setupLog.Error(err, "unable to start manager")
						os.Exit(1)
					}
				}
			}
		},
	})

	informerFactory.Start(stopCh)

	return nil
}

func createEmptyVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	_, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err == nil || !(errors.IsNotFound(err)) {
		return err
	}

	emptyVWC := &arv1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ValidatingWebhookConfiguration",
			APIVersion: "admissionregistration.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-"),
			Labels: map[string]string{
				"app.kubernetes.io/instance": "dynatrace.kubeform.com",
				"app.kubernetes.io/part-of":  "kubeform.com",
			},
		},
	}
	_, err = vwcClient.ValidatingWebhookConfigurations().Create(context.TODO(), emptyVWC, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func updateVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	vwc, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	path := "/validate-" + strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-") + "-v1alpha1-" + strings.ToLower(gvk.Kind)
	fail := arv1.Fail
	sideEffects := arv1.SideEffectClassNone
	admissionReviewVersions := []string{"v1beta1"}

	rules := []arv1.RuleWithOperations{
		{
			Operations: []arv1.OperationType{
				arv1.Delete,
			},
			Rule: arv1.Rule{
				APIGroups:   []string{strings.ToLower(gvk.Group)},
				APIVersions: []string{gvk.Version},
				Resources:   []string{strings.ToLower(flect.Pluralize(gvk.Kind))},
			},
		},
	}

	data, err := ioutil.ReadFile("/tmp/k8s-webhook-server/serving-certs/ca.crt")
	if err != nil {
		return err
	}

	name := strings.ToLower(gvk.Kind) + "." + gvk.Group
	for _, webhook := range vwc.Webhooks {
		if webhook.Name == name {
			return nil
		}
	}

	newWebhook := arv1.ValidatingWebhook{
		Name: name,
		ClientConfig: arv1.WebhookClientConfig{
			Service: &arv1.ServiceReference{
				Namespace: webhookNamespace,
				Name:      webhookName,
				Path:      &path,
			},
			CABundle: data,
		},
		Rules:                   rules,
		FailurePolicy:           &fail,
		SideEffects:             &sideEffects,
		AdmissionReviewVersions: admissionReviewVersions,
	}

	vwc.Webhooks = append(vwc.Webhooks, newWebhook)

	_, err = vwcClient.ValidatingWebhookConfigurations().Update(context.TODO(), vwc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetupManager(ctx context.Context, mgr manager.Manager, gvk schema.GroupVersionKind, auditor *auditlib.EventPublisher, watchOnlyDefault bool) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "dns.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "DomainName",
	}:
		if err := (&controllersdns.DomainNameReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("DomainName"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_dns_domain_name"],
			TypeName:         "dynatrace_dns_domain_name",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "DomainName")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "dns.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "DomainRecord",
	}:
		if err := (&controllersdns.DomainRecordReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("DomainRecord"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_dns_domain_record"],
			TypeName:         "dynatrace_dns_domain_record",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "DomainRecord")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Firewall",
	}:
		if err := (&controllersfirewall.FirewallReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Firewall"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_firewall"],
			TypeName:         "dynatrace_firewall",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Firewall")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&controllersfirewall.RuleReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Rule"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_firewall_rule"],
			TypeName:         "dynatrace_firewall_rule",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Instance",
	}:
		if err := (&controllersinstance.InstanceReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Instance"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_instance"],
			TypeName:         "dynatrace_instance",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Instance")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "kubernetes.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Cluster",
	}:
		if err := (&controllerskubernetes.ClusterReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Cluster"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_kubernetes_cluster"],
			TypeName:         "dynatrace_kubernetes_cluster",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Cluster")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "kubernetes.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NodePool",
	}:
		if err := (&controllerskubernetes.NodePoolReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("NodePool"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_kubernetes_node_pool"],
			TypeName:         "dynatrace_kubernetes_node_pool",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "NodePool")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "loadbalancer.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Loadbalancer",
	}:
		if err := (&controllersloadbalancer.LoadbalancerReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Loadbalancer"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_loadbalancer"],
			TypeName:         "dynatrace_loadbalancer",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Loadbalancer")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "network.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Network",
	}:
		if err := (&controllersnetwork.NetworkReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Network"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_network"],
			TypeName:         "dynatrace_network",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Network")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "snapshot.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Snapshot",
	}:
		if err := (&controllerssnapshot.SnapshotReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Snapshot"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_snapshot"],
			TypeName:         "dynatrace_snapshot",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Snapshot")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ssh.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Key",
	}:
		if err := (&controllersssh.KeyReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Key"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_ssh_key"],
			TypeName:         "dynatrace_ssh_key",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Key")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "template.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Template",
	}:
		if err := (&controllerstemplate.TemplateReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Template"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_template"],
			TypeName:         "dynatrace_template",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Template")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&controllersvolume.VolumeReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Volume"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_volume"],
			TypeName:         "dynatrace_volume",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Volume")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&controllersvolume.AttachmentReconciler{
			Client:           mgr.GetClient(),
			Log:              ctrl.Log.WithName("controllers").WithName("Attachment"),
			Scheme:           mgr.GetScheme(),
			Gvk:              gvk,
			Provider:         dynatrace.Provider(),
			Resource:         dynatrace.Provider().ResourcesMap["dynatrace_volume_attachment"],
			TypeName:         "dynatrace_volume_attachment",
			WatchOnlyDefault: watchOnlyDefault,
		}).SetupWithManager(ctx, mgr, auditor); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Attachment")
			return err
		}

	default:
		return fmt.Errorf("Invalid CRD")
	}

	return nil
}

func SetupWebhook(mgr manager.Manager, gvk schema.GroupVersionKind) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "dns.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "DomainName",
	}:
		if err := (&dnsv1alpha1.DomainName{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "DomainName")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "dns.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "DomainRecord",
	}:
		if err := (&dnsv1alpha1.DomainRecord{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "DomainRecord")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Firewall",
	}:
		if err := (&firewallv1alpha1.Firewall{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Firewall")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "firewall.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&firewallv1alpha1.Rule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "instance.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Instance",
	}:
		if err := (&instancev1alpha1.Instance{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Instance")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "kubernetes.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Cluster",
	}:
		if err := (&kubernetesv1alpha1.Cluster{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Cluster")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "kubernetes.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NodePool",
	}:
		if err := (&kubernetesv1alpha1.NodePool{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "NodePool")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "loadbalancer.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Loadbalancer",
	}:
		if err := (&loadbalancerv1alpha1.Loadbalancer{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Loadbalancer")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "network.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Network",
	}:
		if err := (&networkv1alpha1.Network{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Network")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "snapshot.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Snapshot",
	}:
		if err := (&snapshotv1alpha1.Snapshot{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Snapshot")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ssh.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Key",
	}:
		if err := (&sshv1alpha1.Key{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Key")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "template.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Template",
	}:
		if err := (&templatev1alpha1.Template{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Template")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Volume",
	}:
		if err := (&volumev1alpha1.Volume{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Volume")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "volume.dynatrace.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Attachment",
	}:
		if err := (&volumev1alpha1.Attachment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Attachment")
			return err
		}

	default:
		return fmt.Errorf("Invalid Webhook")
	}

	return nil
}
