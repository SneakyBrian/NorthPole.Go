package main

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func santaHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got request for %s\n", r.URL.Path[1:])

	if r.Method == "POST" {

		query := r.URL.Query()

		timestamp, _ := strconv.ParseInt(query.Get("timestamp"), 10, 64)
		hash := query.Get("hash")
		key := query.Get("key")

		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		bodyString := string(body)

		//fmt.Printf("got request for %s, body: %s\n", r.URL.Path[1:], bodyString)

		age := getWrappingAge(bodyString, timestamp, hash, key)

		ageString := strconv.FormatInt(age, 10)

		ageBytes := []byte(ageString)

		w.Write(ageBytes)

	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

}
