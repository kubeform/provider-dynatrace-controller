package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"reflect"
	"strings"

	"github.com/dtcookie/dynatrace/rest/credentials"
)

// Verbose allows to get output for HTTP communcation via logging
var Verbose = false

var jar = createJar()

func createJar() *cookiejar.Jar {
	jar, _ := cookiejar.New(nil)
	return jar
}

// Client TODO: documentation
type Client struct {
	config      *Config
	credentials credentials.Credentials
	httpClient  *http.Client
	apiBaseURL  string
}

// NewClient TODO: documentation
func NewClient(config *Config, apiBaseURL string, credentials credentials.Credentials) *Client {
	client := Client{}
	client.credentials = credentials
	client.config = config
	client.apiBaseURL = apiBaseURL
	client.httpClient = createHTTPClient(config)
	return &client
}

func createHTTPClient(config *Config) *http.Client {
	var httpClient *http.Client
	if config.NoProxy {
		if config.Insecure {
			httpClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					Proxy:           http.ProxyURL(nil)}}
		} else {
			httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(nil)}}
		}
	} else {
		if config.Insecure {
			httpClient = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		} else {
			httpClient = &http.Client{}
		}
	}
	httpClient.Jar = jar
	return httpClient
}

func (client *Client) getURL(path string) string {
	apiBaseURL := client.apiBaseURL
	if !strings.HasSuffix(apiBaseURL, "/") {
		apiBaseURL = apiBaseURL + "/"
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	return apiBaseURL + path
}

// GET TODO: documentation
func (client *Client) GET(path string, expectedStatusCode int) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request

	url := client.getURL(path)
	if Verbose {
		log.Println(fmt.Sprintf("GET %s", url))
	}
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return readHTTPResponse(httpResponse, http.MethodGet, url, expectedStatusCode, nil, nil)
}

// NewPOST TODO: documentation
func (client *Client) NewPOST(path string, payload interface{}) *Post {
	return newPost(client, path, payload)
}

// NewPUT TODO: documentation
func (client *Client) NewPUT(path string, payload interface{}) *Put {
	return newPut(client, path, payload)
}

// POST TODO: documentation
func (client *Client) POST(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPost, payload, expectedStatusCode, nil, nil)
}

// DELETE TODO: documentation
func (client *Client) DELETE(path string, expectedStatusCode int) ([]byte, error) {
	var err error
	var httpResponse *http.Response
	var request *http.Request

	url := client.getURL(path)
	if request, err = http.NewRequest(http.MethodDelete, url, nil); err != nil {
		return make([]byte, 0), err
	}
	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}
	if Verbose {
		log.Println(fmt.Sprintf("%s %s", strings.ToUpper("DELETE"), url))
	}

	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return make([]byte, 0), err
	}
	return readHTTPResponse(httpResponse, http.MethodDelete, url, expectedStatusCode, nil, nil)
}

// PUT TODO: documentation
func (client *Client) PUT(path string, payload interface{}, expectedStatusCode int) ([]byte, error) {
	return client.send(path, http.MethodPut, payload, expectedStatusCode, nil, nil)
}

func (client *Client) send(path string, method string, payload interface{}, expectedStatusCode int, onResponse func(int) error, customize func(*http.Response)) ([]byte, error) {
	var err error
	var request *http.Request
	var httpResponse *http.Response
	var requestbody []byte

	if requestbody, err = json.Marshal(payload); err != nil {
		return nil, err
	}

	url := client.getURL(path)
	if Verbose {
		log.Println(fmt.Sprintf("%s %s", strings.ToUpper(method), url))
		log.Println("  Request Body: " + string(requestbody))
	}
	if request, err = http.NewRequest(method, url, bytes.NewReader(requestbody)); err != nil {
		return nil, err
	}

	if err = client.credentials.Authenticate(request); err != nil {
		return make([]byte, 0), err
	}

	request.Header.Add("Content-Type", "application/json")
	if httpResponse, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}
	return readHTTPResponse(httpResponse, method, url, expectedStatusCode, onResponse, customize)
}

func readHTTPResponse(httpResponse *http.Response, method string, url string, expectedStatusCode int, onResponse func(int) error, customize func(*http.Response)) ([]byte, error) {
	var err error
	var body []byte
	defer httpResponse.Body.Close()

	if Verbose {
		log.Println(fmt.Sprintf("  %d %s", httpResponse.StatusCode, http.StatusText(httpResponse.StatusCode)))
	}

	if onResponse != nil {
		if err = onResponse(httpResponse.StatusCode); err != nil {
			return nil, err
		}
	}

	if httpResponse.StatusCode != expectedStatusCode {
		finalError := fmt.Errorf("%s (%s) %s", http.StatusText(httpResponse.StatusCode), method, url)
		if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, finalError
		}
		if (body != nil) && len(body) > 0 {
			if Verbose {
				log.Println("  Response Body: " + string(body))
			}
			var env ErrorEnvelope
			if err = json.Unmarshal(body, &env); err == nil {
				finalError = &env.Error
			}
		}
		return body, finalError
	}
	if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
		return nil, err
	}
	if Verbose && (body != nil) && len(body) > 0 {
		log.Println("  Response Body: " + string(body))
	}

	if (body != nil) && len(body) > 0 {
		m := map[string]interface{}{}
		if err = json.Unmarshal(body, &m); err == nil {
			clean(m)
			var cleanBody []byte
			if cleanBody, err = json.Marshal(m); err == nil {
				if Verbose && (cleanBody != nil) && len(cleanBody) > 0 {
					log.Println("  Clean Response Body: " + string(cleanBody))
				}
				return cleanBody, nil
			}
		}
	}

	if customize != nil {
		customize(httpResponse)
	}

	return body, nil
}

func clean(v interface{}) {
	if v == nil {
		return
	}
	switch rv := v.(type) {
	case map[string]interface{}:
		for k, v := range rv {
			if v == nil {
				delete(rv, k)
			} else {
				clean(v)
			}
		}
	case []interface{}:
		for _, e := range rv {
			clean(e)
		}
	case string, float64, bool:
		return
	default:
		panic(reflect.TypeOf(v))
	}
}
