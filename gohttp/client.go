package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {

	//one pointer to client
	client *http.Client

	//Parameters to be set on client.
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration

	//holds all headers
	Headers http.Header
}

// make a singelton client
func New() *httpClient {

	httpClient := &httpClient{}
	return httpClient
}

type HttpClient interface {
	getRequestBody(headerType string, body interface{}) ([]byte, error)
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	SetConnectionTimeout(timeout time.Duration)
	SetMaxIdleConnections(connections int)
	SetRepsonseTimeout(timeout time.Duration)
}

//Set configuration of client
func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) SetMaxIdleConnections(connections int) {
	c.maxIdleConnections = connections
}

func (c *httpClient) SetRepsonseTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

//Set headers
func (c *httpClient) SetHeaders(headers http.Header) {

	c.Headers = headers

}

//Methods
func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
