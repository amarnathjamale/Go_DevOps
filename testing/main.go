package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

}
