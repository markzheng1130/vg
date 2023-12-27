package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandlder(writer http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	nameList, ok := q["name"] // nameList值，會是1個「[]string」

	var message []byte
	if !ok {
		message = []byte("hello, world!")
	} else {
		message = []byte(fmt.Sprintf("hello, world!, %s", nameList[0]))
	}

	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandlder)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// http://localhost:8080/hello
// http://localhost:8080/hello?name=peter
