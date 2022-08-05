package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
	TODO:
	-make docker image
	-setup docker containers
	-setup test function to invoke the docker containers
	-have test function concatenate the index to the index.html file for each container

*/
