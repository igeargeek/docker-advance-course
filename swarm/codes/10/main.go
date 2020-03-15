// https://medium.com/faun/docker-multi-stage-build-for-go-a6821eabde1a

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverAddress = ":3030"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome")
	fmt.Fprintf(w, "Hello version 2.0, welcome %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on port ", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
