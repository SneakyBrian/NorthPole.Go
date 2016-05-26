package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type elfResponse struct {
	Hash      string
	Timestamp int64
}

func elfHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		query := r.URL.Query()
		key := query.Get("key")

		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		bodyString := string(body)

		hash, timestamp := getWrapping(bodyString, key)

		response := elfResponse{hash, timestamp}

		b, _ := json.Marshal(response)

		w.Write(b)

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}
