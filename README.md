# load-balancer
A toy load balancer in go

Implemented starting from the tut here:
https://kasvith.me/posts/lets-create-a-simple-lb-go/

Going to work on adding features mentioned at the end

Add way to check whether a given backend actually works
instead of just responding to a ping

config file should include a request to use to see if
the server is responding
- time how long it takes to respond and log that
- prioritize faster backends instead of round robin?
