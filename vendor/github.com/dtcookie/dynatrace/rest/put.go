package rest

import "net/http"

// Put TODO: documentation
type Put struct {
	client             *Client
	path               string
	payload            interface{}
	expectedStatusCode int
	onResponse         func(int) error
}

func newPut(client *Client, path string, payload interface{}) *Put {
	return &Put{client: client, path: path, payload: payload, expectedStatusCode: 200}
}

// Expect TODO: documentation
func (put *Put) Expect(statusCode int) *Put {
	put.expectedStatusCode = statusCode
	return put
}

// OnResponse TODO: documentation
func (put *Put) OnResponse(fn func(int) error) *Put {
	put.onResponse = fn
	return put
}

// Send TODO: documentation
func (put *Put) Send() ([]byte, error) {
	return put.client.send(put.path, http.MethodPut, put.payload, put.expectedStatusCode, put.onResponse, nil)
}
