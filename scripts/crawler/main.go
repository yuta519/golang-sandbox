package main

import (
	"fmt"
	"os"

	"golang_sandbox/scripts/crawler/pkg"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://www.linkedin.com/jobs/search?keywords=&location=Vancouver%2C%20British%20Columbia%2C%20Canada&geoId=103366113&trk=guest_homepage-basic_jobs-search-bar_search-submit&position=1&pageNum=0")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("ul", "class", "jobs-search__results-list").FindAll("li")
	for _, link := range links {
		url := link.Find("a", "class", "base-card__full-link")
		if url.Pointer != nil {
			fmt.Println(url.Attrs()["href"])
			pkg.ExtractCompanyInfo(url.Attrs()["href"])
			fmt.Println()
		}
	}
}
