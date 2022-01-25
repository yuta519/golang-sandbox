package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Exit(1)
	}

	SECRET := "Bearer " + os.Getenv("NOTION_SECRET")
	fmt.Println(fetch_database_id(SECRET))
	fmt.Println(fetch_page_id(SECRET, "xxxxxxxxxxxxxxxxxxxxxxxxxx"))
}

func fetch_database_id(SECRET string) string {
	req, _ := http.NewRequest("GET", "https://api.notion.com/v1/databases", nil)
	req.Header.Add("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func fetch_page_id(SECRET string, db_id string) string {
	endpoint := "https://api.notion.com/v1/databases/" + db_id + "/query"
	fmt.Println(endpoint)
	payload := strings.NewReader("{\"page_size\":100}")
	req, _ := http.NewRequest("POST", endpoint, payload)
	req.Header.Add("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
