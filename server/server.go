package main

import (
	"fmt"
	"net/http"
	"time"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {

	time.Sleep(time.Second * 2)
	fmt.Fprint(res, "Hello Peter")
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)

}
