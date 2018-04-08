package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getHeaders(url string, port int) {

	// Perform HTTP HEAD
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}

	// Print out each header key and value pair
	for key, value := range response.Header {
		fmt.Println("+ Got a header...")
		headerLine := key + " " + value[0]
		targetHeaders = append(targetHeaders, headerLine)
		//fmt.Printf("%s: %s\n", key, value[0])
		if strings.Contains(key, "Server") {
			targetPorts[port] = value[0]
		}
	}
}
