package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func isHTTP(ipToScan string, openPorts []int, webServer []int) []int {

	for _, port := range openPorts {

		url := "http://" + ipToScan + ":" + strconv.Itoa(port)

		_, err := http.Get(url)

		if err == nil {
			webServer = append(webServer, port)
			fmt.Println("Found a webserver on: ", port)
		}

	}

	return webServer
}
