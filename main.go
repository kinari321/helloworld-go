package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World,%s3度目からの４度目", request.URL.Path[1:])
}

func main() {
	key := ":" + os.Getenv("LOCALHELLO")
	fmt.Printf("type: %T, key:%v", key, key)
	http.HandleFunc("/", handler)
	http.ListenAndServe(key, nil)
}
