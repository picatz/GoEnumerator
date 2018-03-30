package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var openPorts []int
var webBusterResult []string
var targetPorts = map[int]string{}
var targetComments []string

// Config type Here I create custom config type
type config struct {
	DicWeb    string
	DicPass   string
	PortStart int
	PortEnd   int
	Threads   int
}

func main() {

	var webServer []int

	if len(os.Args) != 2 {
		fmt.Println(os.Args[0] + " - Perform pasive enumeration to given ip.")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " 127.0.0.1")
		os.Exit(1)
	}

	TargetToScan := os.Args[1]

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := config{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Create output directory
	if _, err := os.Stat(TargetToScan); os.IsNotExist(err) {
		os.Mkdir(TargetToScan, 0750)
	}

	fmt.Println("About to portmap target: ", TargetToScan)
	portScan(TargetToScan, Config.PortStart, Config.PortEnd, openPorts)
	//fmt.Println(targetPorts)
	writeResultsInt(TargetToScan, openPorts, "OpenPorts")
	webServer = isHTTP(TargetToScan, openPorts, webServer)

	if len(webServer) > 0 {
		var url string
		for _, port := range webServer {
			if port == 443 {
				url = "https://" + TargetToScan + ":" + strconv.Itoa(port)
			} else {
				url = "http://" + TargetToScan + ":" + strconv.Itoa(port)
			}
			getHeaders(url, port)
			getURLS(url)
			getComments(url)
			webBuster(url, Config.DicWeb, Config.Threads, webBusterResult)
			filename := "webBustingResultsPort" + strconv.Itoa(port)
			writeResultsString(TargetToScan, webBusterResult, filename)

		}
	}

	writeResultString(TargetToScan, targetComments, "targetComments")
	writeResultsMap(TargetToScan, targetPorts, "banners")
	fmt.Printf("Enumeration done!!")

}
