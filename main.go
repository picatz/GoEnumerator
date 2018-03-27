package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var openPorts []int

// Config type Here I create custom config type
type config struct {
	DicWeb  string
	DicPass string
	Threads int
}

func main() {

	var webServer []int

	if len(os.Args) != 2 {
		fmt.Println(os.Args[0] + " - Perform pasive enumeration to given ip.")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " 127.0.0.1")
		os.Exit(1)
	}

	ipToScan := os.Args[1]

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := config{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("About to portmap ip: ", ipToScan)
	fmt.Println(Config.DicWeb, Config.Threads)
	portScan(ipToScan, openPorts)
	webServer = isHTTP(ipToScan, openPorts, webServer)

	if len(webServer) > 0 {
		for _, port := range webServer {
			url := "http://" + ipToScan + ":" + strconv.Itoa(port)
			webBuster(url, Config.DicWeb, Config.Threads)
		}
	}
}
