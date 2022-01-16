package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	// resp, err := http.Get("https://www.linkedin.com/jobs/search/?geoId=103366113&location=Vancouver%2C%20British%20Columbia%2C%20Canada")
	// // resp, err := http.Get("http://example.com/")
	// if err != nil {
	// 	fmt.Println("HTTP Request Error!")
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Parse Error!")
	// }
	// fmt.Println(string(body))

	resp, err := soup.Get("https://www.linkedin.com/jobs/search?keywords=&location=Vancouver%2C%20British%20Columbia%2C%20Canada&geoId=103366113&trk=guest_homepage-basic_jobs-search-bar_search-submit&position=1&pageNum=0")

	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "base-card").Attrs()["data-search-id"]
	fmt.Println(links)

}
