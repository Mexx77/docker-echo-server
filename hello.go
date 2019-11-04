package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	uri := strings.TrimPrefix(r.URL.Path, "/")
	var answer string
	if uri == "" {
		answer = "Hello, World!"
	} else {
		answer = "Hello, " + uri + "!"
	}
	w.Write([]byte(answer))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Please configure environment variable PORT")
		os.Exit(1)
	}

	http.HandleFunc("/", sayHello)
	fmt.Printf("Starting http echo server on port %s\n", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		panic(err)
	}
}