package main

import (
	"fmt"
	"log"
	"net/http"
)

var port string = ":4000"

func home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
