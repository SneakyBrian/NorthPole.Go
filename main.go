package main

//go:generate esc -o static.go static

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	useLocalPtr := flag.Bool("useLocal", false, "use local html files for debug purposes")
	portNumPtr := flag.Uint("port", 8181, "port number for webserver")

	flag.Parse()

	fmt.Printf("Starting NorthPole.Go\n")

	fmt.Printf("useLocal = %s\n", strconv.FormatBool(*useLocalPtr))
	fmt.Printf("port = %d\n", *portNumPtr)

	http.HandleFunc("/elf/", elfHandler)
	http.HandleFunc("/santa/", santaHandler)

	http.Handle("/static/", http.FileServer(FS(*useLocalPtr)))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portNumPtr), nil))

	fmt.Printf("Exiting NorthPole.Go\n")
}
