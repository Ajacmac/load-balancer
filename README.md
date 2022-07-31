# load-balancer
A toy load balancer in go

Implemented starting from the tut here:
https://kasvith.me/posts/lets-create-a-simple-lb-go/

Going to work on adding features mentioned at the end

Add way to check whether a given backend actually works
instead of just responding to a ping

config file should include a request to use to see if
the server is responding
- setup default tcp request using this or similar 

    host := "example.com"
    port := "80"
    timeout := time.Duration(1 * time.Second)
    _, err := net.DialTimeout("tcp", host+":"+port, timeout)
    if err != nil {
        fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
    } else {
        fmt.Printf("%s %s %s\n", host, "responding on port:", port)
    }
    
- time how long it takes to respond and log that
- prioritize faster backends instead of round robin?
- any point in being able to parse api spec to automatically determine endpoint to check with?


TODO:
- start testing
    - do more research on go testing best practices
- implement health checking in independent goroutines
- figure out how to make those actually run in parallel xD
- implement a distributed consensus system 
    - pick something easy like raft or a newer leaderless consensus system
    - create a log of uptime stats and configuration details
    - in order for distributed consensus for configuration to make sense thee will need to be a way to update config without a restart