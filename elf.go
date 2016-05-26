package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ElfResponse struct {
	Hash      string
	Timestamp int64
}

func elfHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got request for %s\n", r.URL.Path[1:])

	if r.Method == "POST" {

		query := r.URL.Query()
		key := query.Get("key")

		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		bodyString := string(body)

		//fmt.Printf("got request for %s, body: %s\n", r.URL.Path[1:], bodyString)

		hash, timestamp := getWrapping(bodyString, key)

		response := ElfResponse{hash, timestamp}

		b, _ := json.Marshal(response)

		w.Write(b)

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}
