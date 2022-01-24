package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Exit(1)
	}

	SECRET := "Bearer " + os.Getenv("NOTION_SECRET")
	req, _ := http.NewRequest("GET", "https://api.notion.com/v1/databases", nil)
	req.Header.Add("Notion-Version", "2021-08-16")
	req.Header.Set("Authorization", SECRET)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
