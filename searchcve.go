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
				for _, cveitems := range Data.(CVEParse).CVEItems {
					for _, Vendor := range cveitems.Cve.Affects.Vendor.VendorData {
						for _, vname := range Vendor.Product.ProductData {

							fmt.Println(vname.ProductName)
							//}
							//if Data.(CVEParse).CVEItems[item].Cve.Affects.Vendor[word] {
							//fmt.Println(Vendor)
						}
					}
				}
			}
		}
	}

}
