package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const BNUM = "BNUM"

func handler(w http.ResponseWriter, req *http.Request) {
	backendNum := os.Getenv(BNUM) // should be working but double check
	fmt.Fprintf(w, "Backend: %s", backendNum)

	log.Printf("working")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
	TODO:
*/
