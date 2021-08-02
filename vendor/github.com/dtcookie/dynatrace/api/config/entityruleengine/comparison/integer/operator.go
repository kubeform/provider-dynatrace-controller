package integer

// Operator Operator of the comparison. You can reverse it by setting **negate** to `true`.
// Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	Equals             Operator
	Exists             Operator
	GreaterThan        Operator
	GreaterThanOrEqual Operator
	LowerThan          Operator
	LowerThanOrEqual   Operator
}{
	"EQUALS",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
