package hypervisor_type

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (v *Value) String() string {
	return string(*v)
}

// Values offers the known enum values
var Values = struct {
	Ahv        Value
	HyperV     Value
	Kvm        Value
	Lpar       Value
	Qemu       Value
	VirtualBox Value
	Vmware     Value
	Wpar       Value
	Xen        Value
}{
	"AHV",
	"HYPER_V",
	"KVM",
	"LPAR",
	"QEMU",
	"VIRTUAL_BOX",
	"VMWARE",
	"WPAR",
	"XEN",
}
