package main

import (
	"fmt"
	"regexp"
	"strings"
)

func searchCVE(Banners map[int]string, CVE map[int]interface{}) {

	//fmt.Println(CVE[2013])

	for _, banner := range Banners {
		re := regexp.MustCompile("[-,/,(,),_]")
		NewBanner := re.ReplaceAllString(strings.TrimSpace(banner), " ")
		BannerSlice := strings.Split(NewBanner, " ")
		for _, word := range BannerSlice {
			fmt.Println(word)
			for _, Data := range CVE {
				for item := range Data.(CVEParse).CVEItems {
					for n, Vendor := range Data.(CVEParse).CVEItems[item].Cve.Affects.Vendor.VendorData {
						fmt.Println(Vendor[n].VendorName)
						//if Data.(CVEParse).CVEItems[item].Cve.Affects.Vendor[word] {
						//fmt.Println(Vendor)
						//}
					}
				}
			}
		}
	}

}
