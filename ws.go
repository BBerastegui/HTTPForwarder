package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReceivedRequest struct {
	Method string
	Url    string
	Header map[string][]string
	Body   string
}

type SentRequest struct {
	Method string
	Url    string
	Header map[string][]string
	Body   string
}

type ReceivedResponse struct {
	Status int
	Header map[string][]string
	Body   string
}

// TODO
// Check for consistency of the request
// Minimum values specified...

func handler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	// Create ReceivedRequest
	var rreq ReceivedRequest
	err := decoder.Decode(&rreq)
	// Error on decoder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[/!\\]Error in decoder:", err)
		return
	}

	// Perform checks mandatory parameters
	if rreq.Url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[i]URL is missing on JSON structure.")
		return
	}

	// Perform the request to the target and store the received response
	rresp := performRequest(rreq)

	// Create Encoder and set writer to write to (w)
	encoder := json.NewEncoder(w)
	// Encode received response and write to writer (w)
	err = encoder.Encode(&rresp)
	// Error in encoder
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("[/!\\]Error in encoder:", err)
		return
	}
}

func performRequest(rreq ReceivedRequest) ReceivedResponse {

	// Setup the request to the target
	req, err := http.NewRequest(rreq.Method, rreq.Url, bytes.NewBuffer([]byte(rreq.Body)))
	// Add the custom headers
	req.Header = rreq.Header

	client := &http.Client{}
	// Perform request and store response on "resp"
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle response (resp) and store in ReceivedResponse (rresp)
	var rresp ReceivedResponse

	// Store HTTP Status code
	rresp.Status = resp.StatusCode
	// Store Headers
	rresp.Header = resp.Header
	// Store Body
	body, _ := ioutil.ReadAll(resp.Body)
	rresp.Body = string(body)

	fmt.Println(rresp)
	return rresp
}

func main() {
	var url = "localhost:8080"
	fmt.Println("Running on:", url)
	http.HandleFunc("/", handler)
	http.ListenAndServe(url, nil)
}
