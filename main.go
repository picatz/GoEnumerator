package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

//globa variables
var openPorts []int
var webBusterResult []string
var targetPorts = map[int]string{}
var targetComments []string
var targetURLS []string
var targetEmails []string
var targetHeaders []string
var targetRobots string
var targetCves []string
var targetCMS []string
var targetLogAdmin []string

// CVE map will join year and the content of that year file.
var CVE = make(map[int]interface{})

// Config type Here I create custom config type
type config struct {
	DicWeb    string
	DicPass   string
	PortStart int
	PortEnd   int
	Threads   int
	CVEPath   string
	Year      int
	YearStart int
}

// CVEParse from downloaded CVE json files
type CVEParse struct {
	CVEItems []struct {
		Cve struct {
			DataType    string `json:"data_type"`
			DataFormat  string `json:"data_format"`
			DataVersion string `json:"data_version"`
			CVEDataMeta struct {
				ID       string `json:"ID"`
				ASSIGNER string `json:"ASSIGNER"`
			} `json:"CVE_data_meta"`
			Affects struct {
				Vendor struct {
					VendorData []struct {
						VendorName string `json:"vendor_name"`
						Product    struct {
							ProductData []struct {
								ProductName string `json:"product_name"`
								Version     struct {
									VersionData []struct {
										VersionValue string `json:"version_value"`
									} `json:"version_data"`
								} `json:"version"`
							} `json:"product_data"`
						} `json:"product"`
					} `json:"vendor_data"`
				} `json:"vendor"`
			} `json:"affects"`
			Problemtype struct {
				ProblemtypeData []struct {
					Description []struct {
						Lang  string `json:"lang"`
						Value string `json:"value"`
					} `json:"description"`
				} `json:"problemtype_data"`
			} `json:"problemtype"`
			References struct {
				ReferenceData []struct {
					URL string `json:"url"`
				} `json:"reference_data"`
			} `json:"references"`
			Description struct {
				DescriptionData []struct {
					Lang  string `json:"lang"`
					Value string `json:"value"`
				} `json:"description_data"`
			} `json:"description"`
		} `json:"cve"`
		Configurations struct {
			CVEDataVersion string `json:"CVE_data_version"`
			Nodes          []struct {
				Operator string `json:"operator"`
				Cpe      []struct {
					Vulnerable          bool   `json:"vulnerable"`
					Cpe22URI            string `json:"cpe22Uri"`
					Cpe23URI            string `json:"cpe23Uri"`
					VersionEndIncluding string `json:"versionEndIncluding,omitempty"`
				} `json:"cpe"`
			} `json:"nodes"`
		} `json:"configurations"`
		Impact struct {
			BaseMetricV2 struct {
				CvssV2 struct {
					Version               string  `json:"version"`
					VectorString          string  `json:"vectorString"`
					AccessVector          string  `json:"accessVector"`
					AccessComplexity      string  `json:"accessComplexity"`
					Authentication        string  `json:"authentication"`
					ConfidentialityImpact string  `json:"confidentialityImpact"`
					IntegrityImpact       string  `json:"integrityImpact"`
					AvailabilityImpact    string  `json:"availabilityImpact"`
					BaseScore             float64 `json:"baseScore"`
				} `json:"cvssV2"`
				Severity                string  `json:"severity"`
				ExploitabilityScore     float64 `json:"exploitabilityScore"`
				ImpactScore             float64 `json:"impactScore"`
				ObtainAllPrivilege      bool    `json:"obtainAllPrivilege"`
				ObtainUserPrivilege     bool    `json:"obtainUserPrivilege"`
				ObtainOtherPrivilege    bool    `json:"obtainOtherPrivilege"`
				UserInteractionRequired bool    `json:"userInteractionRequired"`
			} `json:"baseMetricV2"`
		} `json:"impact"`
		PublishedDate    string `json:"publishedDate"`
		LastModifiedDate string `json:"lastModifiedDate"`
	} `json:"CVE_Items"`
}

var wg sync.WaitGroup

func main() {

	var webServer []int

	if len(os.Args) != 2 {
		fmt.Println(os.Args[0] + " - Perform pasive enumeration to given ip.")
		fmt.Println("Usage: " + os.Args[0] + " <url>")
		fmt.Println("Example: " + os.Args[0] + " 127.0.0.1")
		os.Exit(1)
	}

	targetToScan := os.Args[1]
	var striphttp = strings.NewReplacer("http://", "", "https://", "")
	targetToScan = striphttp.Replace(targetToScan)

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := config{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Create output directory
	if _, err := os.Stat(targetToScan); os.IsNotExist(err) {
		os.Mkdir(targetToScan, 0750)
	}

	// Create CVE directory
	if _, err := os.Stat(Config.CVEPath); os.IsNotExist(err) {
		os.Mkdir(Config.CVEPath, 0755)
	}

	// If list of CMS's do not exist download a copy
	if _, err := os.Stat(Config.CVEPath + "/CMSList"); os.IsNotExist(err) {
		CMSUrl := "https://en.wikipedia.org/wiki/List_of_content_management_systems"
		GetCMSList := getCMSList(CMSUrl)
		CMSFile := Config.CVEPath + "/CMSList"
		file, err := os.Create(CMSFile)
		if err != nil {
			fmt.Printf("Error opening %s because of: %s", CMSFile, err)

		}

		defer file.Close()
		for _, line := range GetCMSList {
			file.WriteString(line + "\n")
		}

	}

	for year := Config.Year; year >= Config.YearStart; year-- {
		if _, err := os.Stat(Config.CVEPath + "/nvdcve-1.0-" + strconv.Itoa(year) + ".json.gz"); os.IsNotExist(err) {
			wg.Add(1)
			go getCVE(strconv.Itoa(year))
		}
	}

	wg.Wait()

	for year := Config.Year; year >= Config.YearStart; year-- {

		CVEParse := CVEParse{}
		JSONFile := Config.CVEPath + "/nvdcve-1.0-" + strconv.Itoa(year) + ".json"
		fmt.Println("Parsin: " + JSONFile)

		file, err := os.Open(JSONFile)
		if err != nil {
			fmt.Println("error:", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		decoder.Decode(&CVEParse)
		CVE[year] = CVEParse

	}

	fmt.Println("About to portmap target: ", targetToScan)
	portScan(targetToScan, Config.PortStart, Config.PortEnd)
	fmt.Println(openPorts)
	writeResultsInt(targetToScan, openPorts, "OpenPorts")
	webServer = isHTTP(targetToScan, openPorts, webServer)

	// if port is a webserver or talks http protocols then run this
	if len(webServer) > 0 {
		var url string
		for _, port := range webServer {
			if port == 443 {
				url = "https://" + targetToScan + ":" + strconv.Itoa(port)
			} else {
				url = "http://" + targetToScan + ":" + strconv.Itoa(port)
			}
			getHeaders(url, port)
			HFile := "HeadersPort-" + strconv.Itoa(port)
			writeResultsString(targetToScan, targetHeaders, HFile)

			getEmails(url)
			EFile := "EmailsPort-" + strconv.Itoa(port)
			writeResultsString(targetToScan, targetEmails, EFile)

			getURLS(url)
			UFile := "URLSPort-" + strconv.Itoa(port)
			writeResultsString(targetToScan, targetURLS, UFile)

			getComments(url)
			CFile := "CommentsPort-" + strconv.Itoa(port)
			writeResultsString(targetToScan, targetComments, CFile)

			getRobots(url + "/robots.txt")
			RFile := "robots.txt-" + strconv.Itoa(port)
			writeResultsSingle(targetToScan, targetRobots, RFile)

			checkCMS(url, Config.CVEPath, port)
			CMSFile := "PossibleCMS-port-" + strconv.Itoa(port)
			writeResultsString(targetToScan, targetCMS, CMSFile)

			webBuster(url, Config.DicWeb, Config.Threads)
			WBFile := "webBustingResultsPort-" + strconv.Itoa(port)
			writeResultsString(targetToScan, webBusterResult, WBFile)

		}
	}

	writeResultsMap(targetToScan, targetPorts, "Banners")

	searchCVE(targetPorts, CVE)
	writeResultsString(targetToScan, targetCves, "FoundVulnerabilities")

	fmt.Printf("\n ************************************************\n")
	fmt.Println("Enumeration done!!")
	fmt.Printf("\nCheck the output files inside directory: %s \n", targetToScan)
	fmt.Println("GoEnumerator by ReK2 and the Hispagatos Hacker collective")
	fmt.Println("GPL v3.0, 2018 check the LICENSE file for details")

}
