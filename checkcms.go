package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// open url and inspect html for any of the words in our CMS list

func checkCMS(url string, path string, port int) {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL. ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	CMSFile := path + "/CMSList"
	file, err := os.Open(CMSFile)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(string(body), scanner.Text()) {
			fmt.Printf("\n+ Found possible CMS: %s\n", scanner.Text())
			targetCMS = append(targetCMS, scanner.Text())
			targetPorts[port] = string(scanner.Text())
		}
	}

}
