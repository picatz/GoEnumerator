package main

import (
	"fmt"
	"regexp"
	"strings"
)

func searchCVE(Banners map[int]string, CVE map[int]interface{}) {

	//fmt.Println(CVE[2013])
	//var word string

	for _, banner := range Banners {
		fmt.Printf("Testing for: %s\n", banner)
		re := regexp.MustCompile("[-,/,(,),_]")
		NewBanner := re.ReplaceAllString(strings.TrimSpace(banner), " ")
		BannerSlice := strings.Split(NewBanner, " ")
		for _, Data := range CVE {
			for _, cveitems := range Data.(CVEParse).CVEItems {
				for _, Vendor := range cveitems.Cve.Affects.Vendor.VendorData {
					for _, vname := range Vendor.Product.ProductData {

						if CaseInContains(vname.ProductName, BannerSlice[0]) {
							//fmt.Println(BannerSlice[0])
							//fmt.Println("Match " + vname.ProductName)
							for _, version := range vname.Version.VersionData {
								if CaseInContains(version.VersionValue, BannerSlice[1]) {
									fmt.Println(version.VersionValue + " Is vulnerable")
								}
							}
						}
					}
				}
			}
		}

	}

}

// CaseInContains small funtion to help with ignoring case on our search above
func CaseInContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}
