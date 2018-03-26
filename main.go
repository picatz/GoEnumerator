package main

import (
	"fmt"
	"os"
	"strconv"
)

var openPorts []int

func main() {

	var webServer []int

	if len(os.Args) != 2 {
		fmt.Println(os.Args[0] + " - Perform pasive enumeration to given ip.")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " 127.0.0.1")
		os.Exit(1)
	}

	ipToScan := os.Args[1]
	fmt.Println("About to portmap ip: ", ipToScan)
	portScan(ipToScan, openPorts)
	webServer = isHTTP(ipToScan, openPorts, webServer)

	if len(webServer) > 0 {
		for _, port := range webServer {
			url := "http://" + ipToScan + ":" + strconv.Itoa(port)
			webScan(url)
		}
	}
}
