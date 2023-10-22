package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		fmt.Fprintf(w, "%v %v\n", name, header)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, strconv.Itoa(len(header))+"\n")
		for _, h := range header {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
