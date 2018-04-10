package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRobots(url string) {

	// Fetch the URL
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL. ", err)
	}

	defer response.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading HTTP body. ", err)
	}

	fmt.Printf("\n+ Found robots.txt\n %s\n", body)
	targetRobots = string(body)
}
