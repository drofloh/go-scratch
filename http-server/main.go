/*
Pretty basic example of a webserver
Taken mainly from:
    https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// couple of variables to demo use of incrementing a counter and using a mutex
var counter int
var mutex = &sync.Mutex{}

// test handler
func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// handler to use to increment a counter and write it to the response
func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	//http.HandleFunc("/", echoString)

	// call the incrementCounter function on /increment
	http.HandleFunc("/increment", incrementCounter)

	// inline function
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hiya")
	})

	// example to server a static file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// run the server and log if any issue.
	log.Fatal(http.ListenAndServe(":8080", nil))

}
