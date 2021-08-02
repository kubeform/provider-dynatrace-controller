package match

type Comparison string

var Comparisons = struct {
	Equals           Comparison
	Contains         Comparison
	StartsWith       Comparison
	EndsWith         Comparison
	DoesNotEqual     Comparison
	DoesNotContain   Comparison
	DoesNotStartWith Comparison
	DoesNotEndWith   Comparison
}{
	Equals:           Comparison("EQUALS"),
	Contains:         Comparison("CONTAINS"),
	StartsWith:       Comparison("STARTS_WITH"),
	EndsWith:         Comparison("ENDS_WITH"),
	DoesNotEqual:     Comparison("DOES_NOT_EQUAL"),
	DoesNotContain:   Comparison("DOES_NOT_CONTAIN"),
	DoesNotStartWith: Comparison("DOES_NOT_START_WITH"),
	DoesNotEndWith:   Comparison("DOES_NOT_END_WITH"),
}
