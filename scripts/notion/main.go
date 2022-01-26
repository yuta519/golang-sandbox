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
	fmt.Println(fetch_page_id(SECRET, "xxxxxxxxxxxxxxxxxxx"))
	fmt.Println(updata_page(SECRET, "xxxxxxxxxxxxxxxxxxx"))
	fmt.Println(create_page(SECRET, "xxxxxxxxxxxxxxxxxxx"))
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
	payload := strings.NewReader("{\"page_size\":100}")
	req, _ := http.NewRequest("POST", endpoint, payload)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func create_page(SECRET string, db_id string) string {
	endpoint := "https://api.notion.com/v1/pages"
	payload := strings.NewReader("{\"parent\": {\"database_id\": \"" + db_id + "\"}, \"properties\": {\"Name\": {\"title\": [{\"text\": {\"content\": \"連休の予定\"}}]}}}")
	req, _ := http.NewRequest("POST", endpoint, payload)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func updata_page(SECRET string, page_id string) string {
	endpoint := "https://api.notion.com/v1/pages/" + page_id
	payload := strings.NewReader("{\"properties\": { \"Name\": {\"title\": [{\"text\": {\"content\": \"カナダでの連休の過ごし方\"}}]}, \"やること\": {\"rich_text\": [{\"text\": {\"content\":\"由莉が行きたいお店に行く\"}}]}}}")
	req, _ := http.NewRequest("PATCH", endpoint, payload)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
