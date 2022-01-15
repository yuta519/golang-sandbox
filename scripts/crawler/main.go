package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.linkedin.com/jobs/view/2800185827/?eBP=JOB_SEARCH_ORGANIC&recommendedFlavor=ACTIVELY_HIRING_COMPANY&refId=OLy91auxhev%2Bji3JoxTGMA%3D%3D&trackingId=ra4i8jiXL5tS4n2mI0oxVA%3D%3D&trk=flagship3_search_srp_jobs")
	// resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("HTTP Request Error!")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Parse Error!")
	}
	fmt.Println(string(body))
}
