package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	log.Printf("\nNew visitor from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Hello friend! Looks like you're visiting from %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
