package pkg

// package main

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

func ExtractCompanyInfo(url string) {
	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	parsed_html := soup.HTMLParse(resp)
	company_name_tag := parsed_html.Find("a", "class", "topcard__org-name-link")
	if company_name_tag.Pointer != nil {
		fmt.Println(company_name_tag.Text())
	} else {
		company_name_tag = parsed_html.Find("a", "class", "ember-view")
		if company_name_tag.Pointer != nil {
			fmt.Println(company_name_tag.Text())
		}
	}

}

// func main() {
// 	ExtractCompanyInfo("https://ca.linkedin.com/jobs/view/human-resources-assistant-at-the-university-of-british-columbia-2873936663?refId=jmf1s0SC410ue01SZVhwbA%3D%3D&trackingId=3texVlPFSg72e6HXWsR9bw%3D%3D&position=1&pageNum=0&trk=public_jobs_jserp-result_search-card")
// 	ExtractCompanyInfo("https://www.linkedin.com/jobs/view/client-services-representative-cervical-cancer-screening-bc-cancer-vancouver-at-bc-cancer-2868211100/?trackingId=UieVaLj6QDAG4NybfRhihQ%3D%3D&refId=s6G3VXyUgkzcqRKQphuLhA%3D%3D&pageNum=0&position=17&trk=public_jobs_jserp-result_search-card&originalSubdomain=ca")
// }
