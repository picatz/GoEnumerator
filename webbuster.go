package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IsLoginAdmin here we check if found folder or file is a login or admin portal
func IsLoginAdmin(word string) bool {
	adminlogin := []string{"admin", "login", "Login", "Admin", "control", "panel", "Control", "Panel"}

	for _, list := range adminlogin {
		if strings.Contains(word, list) {
			return true
		}
	}
	return false
}

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
	// This is to ignore self made TLS certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// increment timeout to avoid url fetch errors
	timeout := time.Duration(5 * time.Second)
	//keepAliveTimeout := 600 * time.Second

	client := http.Client{
		Timeout: timeout,
	}

	// Perform a HEAD only, checking status without
	// downloading the entire file
	response, err := client.Head(targetURL.String())
	if err != nil {
		log.Println("Error fetching ", targetURL.String())
		return
	}

	// Added this to avoid a random bug "panic: runtime error: invalid memory address or nil pointer dereference"
	if response != nil {
		defer response.Body.Close()
	}
	// If server returns 200 OK file can be downloaded
	if response.StatusCode == 200 {
		//log.Println(targetURL.String())
		fmt.Printf("\n+ %s\n", targetURL.String())
		// Check if dir found or file is a login or admin
		if IsLoginAdmin(targetURL.String()) {
			fmt.Printf("\n+ Found possible admin/login entrance: %s\n", targetURL.String())
			targetLogAdmin = append(targetLogAdmin, targetURL.String())
		}
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

	fmt.Println("Starting DirBusting:  ")
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
