package indexed_name

// Operator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	Contains Operator
	Equals   Operator
	Exists   Operator
}{
	"CONTAINS",
	"EQUALS",
	"EXISTS",
}
