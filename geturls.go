package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func getURLS(url string) {

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
	re := regexp.MustCompile("(http|https|ftp).*")
	matches := re.FindAllString(string(body), -1)
	if matches == nil {
		// Clean exit if no matches found
		fmt.Println("No URL's found")
	}

	// Print all emails found
	for _, match := range matches {
		fmt.Println(match)
		targetURLS = append(targetURLS, match)
	}
}
