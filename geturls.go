package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func getURLS(url string) {

	// Fetch the URL
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL. ", err)
	}

	// Extract all links
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println("Error loading HTTP response body. ", err)
	}

	// Find and print all links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println(href)
			targetURLS = append(targetURLS, href)
		}
	})
}
