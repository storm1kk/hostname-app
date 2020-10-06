package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	check(err)
	fmt.Fprintf(w, "%s\n", hostname)
}

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}
