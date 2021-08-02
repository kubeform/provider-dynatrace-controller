package database_topology

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
	Cluster       Value
	Embedded      Value
	Failover      Value
	Ipc           Value
	LoadBalancing Value
	SingleServer  Value
	Unspecified   Value
}{
	"CLUSTER",
	"EMBEDDED",
	"FAILOVER",
	"IPC",
	"LOAD_BALANCING",
	"SINGLE_SERVER",
	"UNSPECIFIED",
}
