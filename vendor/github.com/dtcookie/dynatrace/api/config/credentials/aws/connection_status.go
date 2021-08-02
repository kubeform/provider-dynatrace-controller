package aws

// ConnectionStatus The status of the connection to the AWS environment.
//  * `CONNECTED`: There was a connection within last 10 minutes.
// * `DISCONNECTED`: A problem occurred with establishing connection using these credentials. Check whether the data is correct.
// * `UNINITIALIZED`: The successful connection has never been established for these credentials.
type ConnectionStatus string

// ConnectionStati offers the known enum values
var ConnectionStati = struct {
	Connected     ConnectionStatus
	Disconnected  ConnectionStatus
	Uninitialized ConnectionStatus
}{
	"CONNECTED",
	"DISCONNECTED",
	"UNINITIALIZED",
}
