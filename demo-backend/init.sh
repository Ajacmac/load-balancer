#!/usr/bin/env bash

echo $BNUM >> ./index.html
go run .

# this will keep adding the 1 or 2 to the index.html file every time it runs
# ditch the file server and just use a dynamically generated thing that references
# the environment variables directly