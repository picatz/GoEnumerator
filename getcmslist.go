package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func getCMSList(url string) []string {
	var CMSList []string
	fmt.Printf("\n+ Getting a new copy of list of known CMS's from wikipedia\n")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if i <= 13 {
			s.Find("tr").Each(func(a int, b *goquery.Selection) {
				b.Find("td").Each(func(j int, s2 *goquery.Selection) {
					if j == 0 {

						if len(strings.TrimSpace(s2.Text())) != 0 {
							//fmt.Println(s2.Text())
							CMSList = append(CMSList, s2.Text())
						}
					}
				})
			})
		}
	})
	return CMSList
}
