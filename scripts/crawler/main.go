package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://www.linkedin.com/jobs/search?keywords=&location=Vancouver%2C%20British%20Columbia%2C%20Canada&geoId=103366113&trk=guest_homepage-basic_jobs-search-bar_search-submit&position=1&pageNum=0")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	fmt.Println(doc)
	links := doc.Find("ul", "class", "jobs-search__results-list").FindAll("li")
	for _, link := range links {
		url := link.Find("a", "class", "base-card__full-link")
		if url.Pointer != nil {
			response, err := soup.Get(url.Attrs()["href"])
			if err != nil {
				return
			}
			parsed_html := soup.HTMLParse(response)
			company_name := parsed_html.Find("a", "class", "topcard__org-name-link")
			if company_name.Pointer != nil {
				fmt.Println(company_name.Text())
			}
			fmt.Println(url.Attrs()["href"])
		}
	}
}
