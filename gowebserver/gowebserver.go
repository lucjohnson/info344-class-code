package main

// to run the server in the background do 'nohup <executable name> &'
// can now exit the shell and it will stay running, use ps -e | grep <executable name> to list what is running
// can then kill a process by id if necessary

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"log"
	"runtime"
)

// HelloResponse represents a response from the hello route
type HelloResponse struct {
	// adding the `json:"varName"` allows you to change how you want the JSON to be returned (ex. as lowercase)
	Name string `json:"name"`
	Message string `json:"message"`
	GeneratedAt time.Time `json:"generatedAt"`
}

var memstats = new(runtime.MemStats) 

func getMemStats(w http.ResponseWriter, r *http.Request)  {
	runtime.ReadMemStats(memstats)
	allocstats := make(map[string]uint64)
	allocstats["alloc"] = memstats.Alloc
	allocstats["totalAlloc"] = memstats.TotalAlloc
	j, err := json.Marshal(allocstats)
	if nil != err {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	// get whatever follows /api/v1/hello/ on the requested URL
	name := r.URL.Path[len("/api/v1/hello/"):]
	// create and initialize the response struct
	resp := HelloResponse{Name: name, Message: "Hello " + name, GeneratedAt: time.Now()}
	
	// convert struct to JSON
	j, err := json.Marshal(resp)
	if nil != err {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		// tell the client we are sending back JSON
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	}
	
	// w.Write([]byte("Hello " + name))
}

func main() {
	// http.HandleFunc("/", sayHello)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// adding the trailing slash makes it so any call with this syntax plus anything else after the
	// slash will call the sayHello function. Without a trailing slash the request has to match it exactly
	http.HandleFunc("/api/v1/hello/", sayHello)
	http.HandleFunc("/api/v1/memstats", getMemStats)
		
	fmt.Println("Server listening on port 9000...")
	http.ListenAndServe(":9000", nil)
}