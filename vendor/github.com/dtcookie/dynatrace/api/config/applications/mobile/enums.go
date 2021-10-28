package mobile

type PropertyType string

var PropertyTypes = struct {
	Double PropertyType
	Long   PropertyType
	String PropertyType
}{
	Double: PropertyType("DOUBLE"),
	Long:   PropertyType("LONG"),
	String: PropertyType("STRING"),
}

type Origin string

var Origins = struct {
	API                        Origin
	ServerSideRequestAttribute Origin
}{
	API:                        Origin("API"),
	ServerSideRequestAttribute: Origin("SERVER_SIDE_REQUEST_ATTRIBUTE"),
}

type Aggregation string

var Aggregations = struct {
	First Aggregation
	Last  Aggregation
	Max   Aggregation
	Min   Aggregation
	Sum   Aggregation
}{
	First: Aggregation("FIRST"),
	Last:  Aggregation("LAST"),
	Max:   Aggregation("MAX"),
	Min:   Aggregation("MIN"),
	Sum:   Aggregation("SUM"),
}
