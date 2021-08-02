package rest

import "net/http"

// Post TODO: documentation
type Post struct {
	client             *Client
	path               string
	payload            interface{}
	expectedStatusCode int
	onResponse         func(int) error
	customize          func(*http.Response)
}

func newPost(client *Client, path string, payload interface{}) *Post {
	return &Post{client: client, path: path, payload: payload, expectedStatusCode: 200}
}

// Expect TODO: documentation
func (post *Post) Customize(customize func(*http.Response)) *Post {
	post.customize = customize
	return post
}

// Expect TODO: documentation
func (post *Post) Expect(statusCode int) *Post {
	post.expectedStatusCode = statusCode
	return post
}

// OnResponse TODO: documentation
func (post *Post) OnResponse(fn func(int) error) *Post {
	post.onResponse = fn
	return post
}

// Send TODO: documentation
func (post *Post) Send() ([]byte, error) {
	return post.client.send(post.path, http.MethodPost, post.payload, post.expectedStatusCode, post.onResponse, post.customize)
}
