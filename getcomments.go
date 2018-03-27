package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func getComments(url string) {

	// Fetch the URL and get response
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Look for HTML comments using a regular expression
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	matches := re.FindAllString(string(body), -1)
	if matches == nil {
		// Clean exit if no matches found
		fmt.Println("No HTML comments found.")
	}

	// Print all HTML comments found
	for _, match := range matches {
		fmt.Println(match)
	}
}
