package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpClient_GetRequestHeaders(t *testing.T) {
	/*
		1) White box testings is on the same package while black box testing is on a different package.
		2) one test case for every return in your code

	*/

	//Initialization (Optional)
	client := httpClient{}
	commonHeaders := make(http.Header) //headers is a map of string/string
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	//Execution (Mandatory)
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.GetRequestHeaders(requestHeaders) // takes request specific headers and returns both custom and common headers

	//Validation (Mandatory)
	if len(finalHeaders) != 3 {
		t.Errorf("We expect three headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Errorf("Invalid content-type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Errorf("Invalid User-type received")
	}
	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Errorf("Invalid Request-Id received")
	}

}

//TO TEST THE BODY YOU NEED FOUR TEST CASES, ONE PER RETURN STATEMENT
func TestRequestBody(t *testing.T) {

	//Set up only once
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		//Execution
		body, err := client.getRequestBody("", nil)

		//Validation
		if err != nil {
			t.Error("no error expected when passing in nil body")
		}

		if body != nil {
			t.Error("no body expected whrn passing a nil body")
		}

	})

	t.Run("BodyWithJson", func(t *testing.T) {

		// exection
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		fmt.Println(string(body))

		//validation
		if err != nil {
			t.Error(" no error expected while marshaling json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid body retrieved")
		}

	})

	t.Run("BodyWithXml", func(t *testing.T) {

		//
	})

	t.Run("BodyWithDefaultJson", func(t *testing.T) {

		//
	})

}
