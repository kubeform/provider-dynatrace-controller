package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// ScopeConditions Conditions for data capturing.
type ScopeConditions struct {
	HostGroup         *string                    `json:"hostGroup,omitempty"`         // Only applies to this host group.
	ProcessGroup      *string                    `json:"processGroup,omitempty"`      // Only applies to this process group. Note that this can't be transferred between different clusters or environments.
	ServiceTechnology *ServiceTechnology         `json:"serviceTechnology,omitempty"` // Only applies to this service technology.
	TagOfProcessGroup *string                    `json:"tagOfProcessGroup,omitempty"` // Only apply to process groups matching this tag.
	Unknowns          map[string]json.RawMessage `json:"-"`
}

func (me *ScopeConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"host_group": {
			Type:        hcl.TypeString,
			Description: "Only applies to this host group",
			Optional:    true,
		},
		"process_group": {
			Type:        hcl.TypeString,
			Description: "Only applies to this process group. Note that this can't be transferred between different clusters or environments",
			Optional:    true,
		},
		"service_technology": {
			Type:        hcl.TypeString,
			Description: "Only applies to this service technology",
			Optional:    true,
		},
		"tag_of_process_group": {
			Type:        hcl.TypeString,
			Description: "Only apply to process groups matching this tag",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ScopeConditions) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.HostGroup != nil {
		result["host_group"] = *me.HostGroup
	}
	if me.ProcessGroup != nil {
		result["process_group"] = *me.ProcessGroup
	}
	if me.ServiceTechnology != nil {
		result["service_technology"] = string(*me.ServiceTechnology)
	}
	if me.TagOfProcessGroup != nil {
		result["tag_of_process_group"] = *me.TagOfProcessGroup
	}
	return result, nil
}

func (me *ScopeConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "host_group")
		delete(me.Unknowns, "process_group")
		delete(me.Unknowns, "service_technology")
		delete(me.Unknowns, "tag_of_process_group")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("host_group"); ok {
		me.HostGroup = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("process_group"); ok {
		me.ProcessGroup = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("service_technology"); ok {
		me.ServiceTechnology = ServiceTechnology(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("tag_of_process_group"); ok {
		me.TagOfProcessGroup = opt.NewString(value.(string))
	}
	return nil
}

func (me *ScopeConditions) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("hostGroup", me.HostGroup); err != nil {
		return nil, err
	}
	if err := m.Marshal("processGroup", me.ProcessGroup); err != nil {
		return nil, err
	}
	if err := m.Marshal("serviceTechnology", me.ServiceTechnology); err != nil {
		return nil, err
	}
	if err := m.Marshal("tagOfProcessGroup", me.TagOfProcessGroup); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *ScopeConditions) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("hostGroup", &me.HostGroup); err != nil {
		return err
	}
	if err := m.Unmarshal("processGroup", &me.ProcessGroup); err != nil {
		return err
	}
	if err := m.Unmarshal("serviceTechnology", &me.ServiceTechnology); err != nil {
		return err
	}
	if err := m.Unmarshal("tagOfProcessGroup", &me.TagOfProcessGroup); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// ServiceTechnology Only applies to this service technology.
type ServiceTechnology string

func (me ServiceTechnology) Ref() *ServiceTechnology {
	return &me
}

// ServiceTechnologys offers the known enum values
var ServiceTechnologys = struct {
	ActiveMq                             ServiceTechnology
	ActiveMqArtemis                      ServiceTechnology
	AdoNet                               ServiceTechnology
	AIX                                  ServiceTechnology
	Akka                                 ServiceTechnology
	AmazonRedshift                       ServiceTechnology
	Amqp                                 ServiceTechnology
	ApacheCamel                          ServiceTechnology
	ApacheCassandra                      ServiceTechnology
	ApacheCouchDB                        ServiceTechnology
	ApacheDerby                          ServiceTechnology
	ApacheHTTPClientAsync                ServiceTechnology
	ApacheHTTPClientSync                 ServiceTechnology
	ApacheHTTPServer                     ServiceTechnology
	ApacheKafka                          ServiceTechnology
	ApacheSolr                           ServiceTechnology
	ApacheStorm                          ServiceTechnology
	ApacheSynapse                        ServiceTechnology
	ApacheTomcat                         ServiceTechnology
	Apparmor                             ServiceTechnology
	ApplicationInsightsSdk               ServiceTechnology
	ASPDotNet                            ServiceTechnology
	ASPDotNetCore                        ServiceTechnology
	ASPDotNetCoreSignalr                 ServiceTechnology
	ASPDotNetSignalr                     ServiceTechnology
	AsyncHTTPClient                      ServiceTechnology
	AWSLambda                            ServiceTechnology
	AWSRds                               ServiceTechnology
	AWSService                           ServiceTechnology
	Axis                                 ServiceTechnology
	AzureFunctions                       ServiceTechnology
	AzureServiceBus                      ServiceTechnology
	AzureServiceFabric                   ServiceTechnology
	AzureStorage                         ServiceTechnology
	Boshbpm                              ServiceTechnology
	Citrix                               ServiceTechnology
	CitrixCommon                         ServiceTechnology
	CitrixDesktopDeliveryControllers     ServiceTechnology
	CitrixDirector                       ServiceTechnology
	CitrixLicenseServer                  ServiceTechnology
	CitrixProvisioningServices           ServiceTechnology
	CitrixStorefront                     ServiceTechnology
	CitrixVirtualDeliveryAgent           ServiceTechnology
	CitrixWorkspaceEnvironmentManagement ServiceTechnology
	CloudFoundry                         ServiceTechnology
	CloudFoundryAuctioneer               ServiceTechnology
	CloudFoundryBosh                     ServiceTechnology
	CloudFoundryGorouter                 ServiceTechnology
	Coldfusion                           ServiceTechnology
	Containerd                           ServiceTechnology
	CoreDNS                              ServiceTechnology
	Couchbase                            ServiceTechnology
	Crio                                 ServiceTechnology
	Cxf                                  ServiceTechnology
	Datastax                             ServiceTechnology
	DB2                                  ServiceTechnology
	DiegoCell                            ServiceTechnology
	Docker                               ServiceTechnology
	DotNet                               ServiceTechnology
	DotNetRemoting                       ServiceTechnology
	ElasticSearch                        ServiceTechnology
	Envoy                                ServiceTechnology
	Erlang                               ServiceTechnology
	Etcd                                 ServiceTechnology
	F5Ltm                                ServiceTechnology
	Fsharp                               ServiceTechnology
	Garden                               ServiceTechnology
	Glassfish                            ServiceTechnology
	Go                                   ServiceTechnology
	GraalTruffle                         ServiceTechnology
	Grpc                                 ServiceTechnology
	Grsecurity                           ServiceTechnology
	Hadoop                               ServiceTechnology
	HadoopHdfs                           ServiceTechnology
	HadoopYarn                           ServiceTechnology
	Haproxy                              ServiceTechnology
	Heat                                 ServiceTechnology
	Hessian                              ServiceTechnology
	HornetQ                              ServiceTechnology
	IBMCICSRegion                        ServiceTechnology
	IBMCICSTransactionGateway            ServiceTechnology
	IBMIMSConnectRegion                  ServiceTechnology
	IBMIMSControlRegion                  ServiceTechnology
	IBMIMSMessageProcessingRegion        ServiceTechnology
	IBMIMSSoapGateway                    ServiceTechnology
	IBMIntegrationBus                    ServiceTechnology
	IBMMq                                ServiceTechnology
	IBMMqClient                          ServiceTechnology
	IBMWebshprereApplicationServer       ServiceTechnology
	IBMWebshprereLiberty                 ServiceTechnology
	IIS                                  ServiceTechnology
	IISAppPool                           ServiceTechnology
	Istio                                ServiceTechnology
	Java                                 ServiceTechnology
	JaxWs                                ServiceTechnology
	JBoss                                ServiceTechnology
	JBossEap                             ServiceTechnology
	JdkHTTPServer                        ServiceTechnology
	Jersey                               ServiceTechnology
	Jetty                                ServiceTechnology
	Jruby                                ServiceTechnology
	Jython                               ServiceTechnology
	Kubernetes                           ServiceTechnology
	Libvirt                              ServiceTechnology
	Linkerd                              ServiceTechnology
	Mariadb                              ServiceTechnology
	Memcached                            ServiceTechnology
	MicrosoftSQLServer                   ServiceTechnology
	Mongodb                              ServiceTechnology
	MSSQLClient                          ServiceTechnology
	MuleEsb                              ServiceTechnology
	MySQL                                ServiceTechnology
	MySQLConnector                       ServiceTechnology
	NetflixServo                         ServiceTechnology
	Netty                                ServiceTechnology
	Nginx                                ServiceTechnology
	NodeJs                               ServiceTechnology
	OkHTTPClient                         ServiceTechnology
	OneAgentSdk                          ServiceTechnology
	Opencensus                           ServiceTechnology
	Openshift                            ServiceTechnology
	OpenStackCompute                     ServiceTechnology
	OpenStackController                  ServiceTechnology
	Opentelemetry                        ServiceTechnology
	Opentracing                          ServiceTechnology
	OpenLiberty                          ServiceTechnology
	OracleDatabase                       ServiceTechnology
	OracleWeblogic                       ServiceTechnology
	Owin                                 ServiceTechnology
	Perl                                 ServiceTechnology
	PHP                                  ServiceTechnology
	PHPFpm                               ServiceTechnology
	Play                                 ServiceTechnology
	PostgreSQL                           ServiceTechnology
	PostgreSQLDotNetDataProvider         ServiceTechnology
	PowerDNS                             ServiceTechnology
	Progress                             ServiceTechnology
	Python                               ServiceTechnology
	RabbitMq                             ServiceTechnology
	Redis                                ServiceTechnology
	Resteasy                             ServiceTechnology
	Restlet                              ServiceTechnology
	Riak                                 ServiceTechnology
	Ruby                                 ServiceTechnology
	SagWebmethodsIs                      ServiceTechnology
	SAP                                  ServiceTechnology
	SAPHanadb                            ServiceTechnology
	SAPHybris                            ServiceTechnology
	SAPMaxdb                             ServiceTechnology
	SAPSybase                            ServiceTechnology
	Scala                                ServiceTechnology
	Selinux                              ServiceTechnology
	Sharepoint                           ServiceTechnology
	Spark                                ServiceTechnology
	Spring                               ServiceTechnology
	Sqlite                               ServiceTechnology
	Thrift                               ServiceTechnology
	Tibco                                ServiceTechnology
	TibcoBusinessWorks                   ServiceTechnology
	TibcoEms                             ServiceTechnology
	VarnishCache                         ServiceTechnology
	Vim2                                 ServiceTechnology
	VirtualMachineKvm                    ServiceTechnology
	VirtualMachineQemu                   ServiceTechnology
	Wildfly                              ServiceTechnology
	WindowsContainers                    ServiceTechnology
	Wink                                 ServiceTechnology
	ZeroMq                               ServiceTechnology
}{
	"ACTIVE_MQ",
	"ACTIVE_MQ_ARTEMIS",
	"ADO_NET",
	"AIX",
	"AKKA",
	"AMAZON_REDSHIFT",
	"AMQP",
	"APACHE_CAMEL",
	"APACHE_CASSANDRA",
	"APACHE_COUCH_DB",
	"APACHE_DERBY",
	"APACHE_HTTP_CLIENT_ASYNC",
	"APACHE_HTTP_CLIENT_SYNC",
	"APACHE_HTTP_SERVER",
	"APACHE_KAFKA",
	"APACHE_SOLR",
	"APACHE_STORM",
	"APACHE_SYNAPSE",
	"APACHE_TOMCAT",
	"APPARMOR",
	"APPLICATION_INSIGHTS_SDK",
	"ASP_DOTNET",
	"ASP_DOTNET_CORE",
	"ASP_DOTNET_CORE_SIGNALR",
	"ASP_DOTNET_SIGNALR",
	"ASYNC_HTTP_CLIENT",
	"AWS_LAMBDA",
	"AWS_RDS",
	"AWS_SERVICE",
	"AXIS",
	"AZURE_FUNCTIONS",
	"AZURE_SERVICE_BUS",
	"AZURE_SERVICE_FABRIC",
	"AZURE_STORAGE",
	"BOSHBPM",
	"CITRIX",
	"CITRIX_COMMON",
	"CITRIX_DESKTOP_DELIVERY_CONTROLLERS",
	"CITRIX_DIRECTOR",
	"CITRIX_LICENSE_SERVER",
	"CITRIX_PROVISIONING_SERVICES",
	"CITRIX_STOREFRONT",
	"CITRIX_VIRTUAL_DELIVERY_AGENT",
	"CITRIX_WORKSPACE_ENVIRONMENT_MANAGEMENT",
	"CLOUDFOUNDRY",
	"CLOUDFOUNDRY_AUCTIONEER",
	"CLOUDFOUNDRY_BOSH",
	"CLOUDFOUNDRY_GOROUTER",
	"COLDFUSION",
	"CONTAINERD",
	"CORE_DNS",
	"COUCHBASE",
	"CRIO",
	"CXF",
	"DATASTAX",
	"DB2",
	"DIEGO_CELL",
	"DOCKER",
	"DOTNET",
	"DOTNET_REMOTING",
	"ELASTIC_SEARCH",
	"ENVOY",
	"ERLANG",
	"ETCD",
	"F5_LTM",
	"FSHARP",
	"GARDEN",
	"GLASSFISH",
	"GO",
	"GRAAL_TRUFFLE",
	"GRPC",
	"GRSECURITY",
	"HADOOP",
	"HADOOP_HDFS",
	"HADOOP_YARN",
	"HAPROXY",
	"HEAT",
	"HESSIAN",
	"HORNET_Q",
	"IBM_CICS_REGION",
	"IBM_CICS_TRANSACTION_GATEWAY",
	"IBM_IMS_CONNECT_REGION",
	"IBM_IMS_CONTROL_REGION",
	"IBM_IMS_MESSAGE_PROCESSING_REGION",
	"IBM_IMS_SOAP_GATEWAY",
	"IBM_INTEGRATION_BUS",
	"IBM_MQ",
	"IBM_MQ_CLIENT",
	"IBM_WEBSHPRERE_APPLICATION_SERVER",
	"IBM_WEBSHPRERE_LIBERTY",
	"IIS",
	"IIS_APP_POOL",
	"ISTIO",
	"JAVA",
	"JAX_WS",
	"JBOSS",
	"JBOSS_EAP",
	"JDK_HTTP_SERVER",
	"JERSEY",
	"JETTY",
	"JRUBY",
	"JYTHON",
	"KUBERNETES",
	"LIBVIRT",
	"LINKERD",
	"MARIADB",
	"MEMCACHED",
	"MICROSOFT_SQL_SERVER",
	"MONGODB",
	"MSSQL_CLIENT",
	"MULE_ESB",
	"MYSQL",
	"MYSQL_CONNECTOR",
	"NETFLIX_SERVO",
	"NETTY",
	"NGINX",
	"NODE_JS",
	"OK_HTTP_CLIENT",
	"ONEAGENT_SDK",
	"OPENCENSUS",
	"OPENSHIFT",
	"OPENSTACK_COMPUTE",
	"OPENSTACK_CONTROLLER",
	"OPENTELEMETRY",
	"OPENTRACING",
	"OPEN_LIBERTY",
	"ORACLE_DATABASE",
	"ORACLE_WEBLOGIC",
	"OWIN",
	"PERL",
	"PHP",
	"PHP_FPM",
	"PLAY",
	"POSTGRE_SQL",
	"POSTGRE_SQL_DOTNET_DATA_PROVIDER",
	"POWER_DNS",
	"PROGRESS",
	"PYTHON",
	"RABBIT_MQ",
	"REDIS",
	"RESTEASY",
	"RESTLET",
	"RIAK",
	"RUBY",
	"SAG_WEBMETHODS_IS",
	"SAP",
	"SAP_HANADB",
	"SAP_HYBRIS",
	"SAP_MAXDB",
	"SAP_SYBASE",
	"SCALA",
	"SELINUX",
	"SHAREPOINT",
	"SPARK",
	"SPRING",
	"SQLITE",
	"THRIFT",
	"TIBCO",
	"TIBCO_BUSINESS_WORKS",
	"TIBCO_EMS",
	"VARNISH_CACHE",
	"VIM2",
	"VIRTUAL_MACHINE_KVM",
	"VIRTUAL_MACHINE_QEMU",
	"WILDFLY",
	"WINDOWS_CONTAINERS",
	"WINK",
	"ZERO_MQ",
}
