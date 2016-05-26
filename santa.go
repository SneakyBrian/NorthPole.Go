package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func santaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		query := r.URL.Query()

		timestamp, _ := strconv.ParseInt(query.Get("timestamp"), 10, 64)
		hash := query.Get("hash")
		key := query.Get("key")

		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		bodyString := string(body)

		age := getWrappingAge(bodyString, timestamp, hash, key)

		ageString := strconv.FormatInt(age, 10)

		ageBytes := []byte(ageString)

		w.Write(ageBytes)

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}
