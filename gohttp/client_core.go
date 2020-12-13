package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

// main function doing the request call and returning a response or a error
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	//Get all headers by sending in custom headers
	fullHeader := c.GetRequestHeaders(headers)

	//get and Parse body
	requestBody, err := c.getRequestBody(fullHeader.Get("Content-Type"), body)
	if err != nil {
		return nil, errors.New("unable to process body")
	}

	//make request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to process request")
	}

	//Set header to request
	req.Header = fullHeader

	client := c.getHttpClient()

	// call the method on the client of the httpclient
	return client.Do(req)
}

//Get client
func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnection(),
			ResponseHeaderTimeout: c.getResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeout(),
			}).DialContext,
		},
	}

	return c.client

}

func (c *httpClient) getMaxIdleConnection() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {

	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}
	return defaultConnectionTimeout
}

// will return all headers (custom and default)
func (c *httpClient) GetRequestHeaders(requestHeaders http.Header) http.Header {

	// is a map[string]string
	result := make(http.Header)

	// ADD COMMON HEADERS TO REQUEST (headers is a map[string]string)
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// ADD CUSTOM HEADERS TO REQUEST (headers is a map[string]string)
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}

	}
	return result
}
