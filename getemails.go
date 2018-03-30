package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func getEmails(url string) {

	// Fetch the URL
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL. ", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading HTTP body. ", err)
	}

	// Look for mailto: links using a regular expression
	re := regexp.MustCompile("\"mailto:.*?[?\"]")
	matches := re.FindAllString(string(body), -1)
	if matches == nil {
		// Clean exit if no matches found
		fmt.Println("No emails found.")
	}

	// Print all emails found
	for _, match := range matches {
		// Remove "mailto prefix and the trailing quote or question mark
		// by performing a slice operation to extract the substring
		cleanedMatch := match[8 : len(match)-1]
		fmt.Println(cleanedMatch)
	}
}
