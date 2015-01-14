package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type ReceivedRequest struct {
	url    string
	method string
	data   string
}

type SentRequest struct {
	headers []string
	url     string
	method  string
	data    string
}

type ReceivedResponse struct {
	headers []string
	url     string
	method  string
	data    string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var rr ReceivedRequest

	// Save parameters on rr
	rr.method = r.FormValue("method")
	rr.url = r.FormValue("url")

	if rr.url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rr.method == "POST" {
		fmt.Println("The method was post.")
		rr.method = "POST"
		var b bytes.Buffer
		_, err := io.Copy(b, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		rr.data = b.String()
		fmt.Println(rr.method)
	} else {
		fmt.Println("The method was not post. Defaulting to GET.")
		rr.method = "GET"
	}

	fmt.Println(rr)

	//fmt.Fprintf(w,"%s", m.method)
	//fmt.Fprintf(w,"%s", m.headers)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}
