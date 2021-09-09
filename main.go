package main

import (
	"fmt"
	"log"
	// "net/http"
	"os"

	"github.com/joho/godotenv"
)

// func handler(writer http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(writer, "Hello World !!%s\n", request.URL.Path[1:])
// }

func main() {
	// key := ":" + os.Getenv("LOCALHELLO")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	userName := os.Getenv("HOGE")

	// fmt.Printf("type: %T, key: %v\n", key, key)
	fmt.Printf("type: %T, hoge: %v\n", userName, userName)
	// fmt.Printf("type: %T, pass: %v\n", password, password)

	// http.HandleFunc("/", handler)
	// http.ListenAndServe(key, nil)
}
