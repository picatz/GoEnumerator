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
		re := regexp.MustCompile("[-,/,(,),_]")
		NewBanner := re.ReplaceAllString(strings.TrimSpace(banner), " ")
		BannerSlice := strings.Split(NewBanner, " ")
		for _, Data := range CVE {
			for _, cveitems := range Data.(CVEParse).CVEItems {
				for _, Vendor := range cveitems.Cve.Affects.Vendor.VendorData {
					for _, vname := range Vendor.Product.ProductData {

						if strings.Contains(vname.ProductName, BannerSlice[0]) {
							fmt.Println(BannerSlice[0])
							fmt.Println("Match " + vname.ProductName)
						}
						//}
						//if Data.(CVEParse).CVEItems[item].Cve.Affects.Vendor[word] {
						//fmt.Println(Vendor)
					}
				}
			}
		}

	}

}
