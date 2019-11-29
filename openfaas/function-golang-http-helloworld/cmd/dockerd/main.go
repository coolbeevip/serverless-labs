package main

import (
	"fmt"
	"function-golang-http-helloworld/internal/inspector"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	readTimeout := inspector.ParseIntOrDurationValue(os.Getenv("read_timeout"), 10*time.Second)
	writeTimeout := inspector.ParseIntOrDurationValue(os.Getenv("write_timeout"), 10*time.Second)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8082),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	http.HandleFunc("/", inspector.MakeRequestHandler())
	log.Fatal(s.ListenAndServe())
}
