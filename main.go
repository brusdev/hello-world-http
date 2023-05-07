package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const defaultHelloAddr = ":8080"
const defaultHelloText = "Hello, World!"

func main() {
	fmt.Println("Hello, World HTTP!")

	helloAddr := os.Getenv("HELLO_ADDR")
	if len(helloAddr) == 0 {
		helloAddr = defaultHelloAddr
		fmt.Println("The default serving address is " + defaultHelloAddr +
			", set HELLO_ADDR to change the it.")
	}

	helloText := os.Getenv("HELLO_TEXT")
	if len(helloText) == 0 {
		helloText = defaultHelloText
		fmt.Println("The default response text is " + defaultHelloText +
			", set HELLO_TEXT to change the it.")
	}

	fmt.Println("Serving " + helloAddr + " with the text " + helloText)
	http.ListenAndServe(helloAddr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%d %s > %s > %s\n", time.Now().Unix(),
				r.RemoteAddr, r.RequestURI, helloText)
			w.Write([]byte(helloText + "\n"))
		}))
}
