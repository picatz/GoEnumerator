package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHeaders(url string) {

	// Perform HTTP HEAD
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	// Print out each header key and value pair
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value[0])
	}
}
