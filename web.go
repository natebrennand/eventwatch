package main

import (
	"log"
	"net/http"
	"os"
)

var (
	port string
)

func init() {
	if port = os.Getenv("PORT"); port == "" {
		log.Println("'PORT' can be set as an environment variable, we will listen on port 5000")
		port = "5000"
	}
}

func helloworld() {
	http.Handle("/", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte(`This is the UEM event watcher.
Please visit https://github.com/natebrennand/eventwatch for more information`,
		))
	}))

	panic(http.ListenAndServe(":"+port, nil))
}
