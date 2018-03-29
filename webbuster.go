package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	"os"
)

// perform an HTTP HEAD and see if the path exists.
// If the path returns a 200 OK print out the path
// and acumulate the string slice to use later
func checkIfURLExists(checkIfbaseURL, filePath string, doneChannel chan bool) {
	// Create URL object from raw string
	targetURL, err := url.Parse(checkIfbaseURL)
	if err != nil {
		log.Println("Error parsing base URL. ", err)
	}

	// Set the part of the URL after the host name
	targetURL.Path = filePath

	// Perform a HEAD only, checking status without
	// downloading the entire file
	response, err := http.Head(targetURL.String())
	if err != nil {
		log.Println("Error fetching ", targetURL.String())
	}

	// Added this to avoid a random bug "panic: runtime error: invalid memory address or nil pointer dereference"
	defer response.Body.Close()
	// If server returns 200 OK file can be downloaded
	if response.StatusCode == 200 {
		log.Println(targetURL.String())
		// increment slice with 200 result
		webBusterResult = append(webBusterResult, targetURL.String())

	}
	// Signal completion so next thread can start
	doneChannel <- true
	return
}

func webBuster(url string, dicWeb string, Threads int, webBusterResult []string) {
	// Load command line arguments
	wordlistFilename := dicWeb
	checkIfbaseURL := url
	maxThreads := Threads

	// Track how many threads are active to avoid
	// flooding a web server
	activeThreads := 0
	doneChannel := make(chan bool)

	// Open word list file for reading
	wordlistFile, err := os.Open(wordlistFilename)
	if err != nil {
		log.Fatal("Error opening wordlist file. ", err)
	}

	// Read each line and do an HTTP HEAD
	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		go checkIfURLExists(checkIfbaseURL, scanner.Text(), doneChannel)
		activeThreads++

		// Wait until a done signal before next if max threads reached
		if activeThreads >= maxThreads {
			<-doneChannel
			activeThreads--
		}
	}

	// Wait for all threads before repeating and fetching a new batch
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}

	// Scanner errors must be checked manually
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading wordlist file. ", err)
	}
}
