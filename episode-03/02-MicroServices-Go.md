# Introduction to Microservices with Go

## Overview
In this episode, following content are covered:
- Separate handler into a new package. Kind of decouple the handler from the main package.
- Dive into HTTP handler.
  - How a handler is created. 
    - Through a struct and the "ServeHttp" method.
    ```go
    type Goodbye struct {
	    l *log.Logger
    }

    func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	    // sleep 30 seconds before execute the shutdown
	    time.Sleep(20 * time.Second)
	    rw.Write([]byte("Bye!"))
    }
    ```
- Logging
  - Logging is implemented using depdenency injection. The handler object will take a logger object as a parameter. While you create the handler, you can pass the logger object to it. This make the parameter flexible and easy to test. (When you have different kind of logger, you can pass different logger to the handler object.)
- Golang Server 
  - Create a new server object.
  - Why we need to create a new server object? 
    - With the default server, the server is easy to be attacked with denial of service attack. 
    - When there is a blocking request (e.g. sleep 20 seconds, upload a large file), the server will not be able to handle other requests.
  - How it got implemented?
    ```go
    s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 25 * time.Second,
		ReadTimeout:  25 * time.Second,
	}
    ```
- Graceful Shutdown
  - We want to shutdown the server gracefully. With the graceful shutdown, the server will not accept new requests, but it will finish the existing requests before it shutdown.
  - However, the example in the video is not working. The server will not shutdown gracefully.

## Graceful Shutdown

### How to test out the example in video?




## Extended reading
- How Go Channel go implemented and how it works?
- How the Go app stopped accepting new requests while it is in the graceful shutdown waiting time?
    