package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러 발생")
		log.Fatalln(err)
	}
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello world")
}
