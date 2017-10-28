package main

import (
	"fmt"
	"log"
	"net/http"
	"projects/mirandae.me/api"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	log.Printf("\nNew visitor from %s", r.RemoteAddr)
	api.LogLocation(fmt.Sprint(r.RemoteAddr))
	fmt.Fprintf(w, "Hello friend! Looks like you're visiting from %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
