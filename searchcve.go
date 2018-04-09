package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func searchCVE(Banners map[int]string, CVE map[int]interface{}) {

	//fmt.Println(CVE[2013])
	var word string
	var n int
	for _, banner := range Banners {
		fmt.Printf("\nTesting for: %s\n", banner)
		re := regexp.MustCompile("[-,/,(,),_]")
		NewBanner := re.ReplaceAllString(strings.TrimSpace(banner), " ")
		BannerSlice := strings.Split(NewBanner, " ")
		for n, word = range BannerSlice {
			if GetWords(word) && len(strings.TrimSpace(word)) != 0 && VerContains(BannerSlice[n+1]) {
				fmt.Printf("checking %s - %s\n", word, BannerSlice[n+1])

				for _, Data := range CVE {
					for _, cveitems := range Data.(CVEParse).CVEItems {
						for _, Vendor := range cveitems.Cve.Affects.Vendor.VendorData {
							for _, vname := range Vendor.Product.ProductData {

								if CaseInContains(vname.ProductName, word) {
									for _, version := range vname.Version.VersionData {
										if strings.EqualFold(version.VersionValue, BannerSlice[n+1]) {
											fmt.Printf("%s %s  is vulnerable\n", vname.ProductName, version.VersionValue)
										}
									}
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

// GetWords will match only words and return true or false
//func GetWords(s string) bool {
//fmt.Printf("Testing for word: %s\n", s)
//	reg := regexp.MustCompile(`(?m)[a-zA-Z][a-zA-Z]+`)
//	return reg.MatchString(s)

//}

// GetWords will match only words and return true or false
func GetWords(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// VerContains here I check if the string has spaces
func VerContains(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
