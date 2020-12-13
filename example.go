package main

import (
	"fmt"
	"github.com/bluedolfin/httpclient/gohttp"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	//make a singleton and reuse for every call
	HttpClient       = gohttp.New()
	GitHubHttpClient = gohttp.New()
)

func main() {
	getURLS()
}

func getGitHubHttpClient() gohttp.HttpClient {

	// create our common headers.
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-ACV")
	GitHubHttpClient.SetHeaders(headers)

	return GitHubHttpClient
}

func getHttpClient() gohttp.HttpClient {

	HttpClient.SetConnectionTimeout(2 * time.Second)
	HttpClient.SetRepsonseTimeout(50 * time.Millisecond)

	return HttpClient
}

func getURLS() {

	response, err := HttpClient.Get("http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	// create our common headers.
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-ACV")
	HttpClient.SetHeaders(headers)

	response, err = GitHubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}
	bytes, err = ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
