package test_backend

import "net/http"

func main() {
	http.Handle("/", http)
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
