package main

// TODO
// - Flag to run in stdin - stdout mode

import (
	//	"bytes"
	"fmt"
	//	"io"
	"net/http"
	"encoding/json"
)

type ReceivedRequest struct {
	Method string
	Url string
	Headers map[string]string
	Data   string
}

type SentRequest struct {
	Method string
	Url string
	Headers map[string]string
	data    string
}

type ReceivedResponse struct {
	Status int
	Headers map[string]string
	Method  string
	Data    string
}

// The expected json to be received:
// {method:"",headers:{"name":"value"},body:""}

// Check for consistency of the request
// Minimum values specified...

func handler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var rr ReceivedRequest
	err := decoder.Decode(&rr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Perform checks
	// Minimum parameters

	if (rr.Method == "" || rr.Url == "" || len(rr.Headers) == 0 || rr.Data == ""){
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Something is missing")
		return
	}

	performRequest(rr)

	if rr.Url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*
	var b bytes.Buffer
	_, err := io.Copy(&b, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	*/
}

func performRequest(rr ReceivedRequest){
	var st SentRequest
	st.Headers = make(map[string]string)

	//	Build response
	for key, value := range rr.Headers {
		fmt.Println("Key:", key, "Value:", value)
		st.Headers[key] = "AAA"
		fmt.Println(st.Headers[key])
	}

	if (rr.Method == "GET"){
		resp, err := http.Get(rr.Url)
		if err != nil {
			fmt.Println("Ooops:",err)
			return
		}
		fmt.Println(resp)
	}
}

/*
func sendRequest() {
	resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
}
*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}
