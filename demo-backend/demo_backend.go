package test_backend

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

/*
	TODO:
	-setup basic static file server
	-make docker image
	-setup docket containers
	-setup test function to invoke the docker containers
	-have test function create the index.html file for each container,
	each html file only containing a number for identifying which backend responded

*/
