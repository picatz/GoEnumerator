package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//globa variables
var openPorts []int
var webBusterResult []string
var targetPorts = map[int]string{}
var targetComments []string
var targetURLS []string
var targetEmails []string
var targetHeaders []string

// Config type Here I create custom config type
type config struct {
	DicWeb    string
	DicPass   string
	PortStart int
	PortEnd   int
	Threads   int
	Year      int
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

	// Create CVE directory
	if _, err := os.Stat("CVE"); os.IsNotExist(err) {
		os.Mkdir("CVE", 0755)
	}

	for year := Config.Year; year >= 2010; year-- {
		if _, err := os.Stat("CVE/nvdcve-1.0-" + strconv.Itoa(year) + ".json.gz"); os.IsNotExist(err) {
			go getCVE(strconv.Itoa(year))
		}
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
			HFile := "HeadersPort-" + strconv.Itoa(port)
			writeResultsString(TargetToScan, targetHeaders, HFile)

			getEmails(url)
			EFile := "EmailsPort-" + strconv.Itoa(port)
			writeResultsString(TargetToScan, targetEmails, EFile)

			getURLS(url)
			UFile := "URLSPort-" + strconv.Itoa(port)
			writeResultsString(TargetToScan, targetURLS, UFile)

			getComments(url)
			CFile := "CommentsPort-" + strconv.Itoa(port)
			writeResultsString(TargetToScan, targetComments, CFile)

			webBuster(url, Config.DicWeb, Config.Threads, webBusterResult)
			WBFile := "webBustingResultsPort-" + strconv.Itoa(port)
			writeResultsString(TargetToScan, webBusterResult, WBFile)

		}
	}

	writeResultsMap(TargetToScan, targetPorts, "Banners")
	fmt.Println("Enumeration done!!")
	fmt.Printf("Check the output files inside directory: %s \n", TargetToScan)
	fmt.Println("GoEnumerator by ReK2 and the Hispagatos Hacker collective")
	fmt.Println("GPL v3.0, 2018 check the LICENSE file for details")

}
