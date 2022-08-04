package demo_backend

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
	-setup basic static file server
	-make docker image
	-setup docket containers
	-setup test function to invoke the docker containers
	-have test function create the index.html file for each container,
	each html file only containing a number for identifying which backend responded

	-test demo server to ensure it works

*/
