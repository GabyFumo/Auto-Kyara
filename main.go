package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"auto-kyara/models"
)

var baseUrl = "https://kyara.games/api"

var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
var WEBHOOK_URL = os.Getenv("WEBHOOK_URL")
var DISCORD_ID = os.Getenv("DISCORD_ID")

func main() {
	if AUTH_TOKEN == "" {
		log.Fatal("Error loading .env file")
	}
	FetchMe()
}

func FetchMe() models.User {
	url := baseUrl + "/me"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+AUTH_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal("Error unmarshaling response: ", err)
	}

	return user
}

func OpenPacks(amount int) []models.OpenOnePack {
	var url string
	if amount > 1 {
		url = baseUrl + "/packs/waifu/open-many"
	} else {
		url = baseUrl + "/packs/waifu/open"
	}

	body := map[string]int{
		"maxChakra": 0,
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Error marshaling body: ", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Add("Authorization", "Bearer "+AUTH_TOKEN)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}

	if amount > 1 {
		var result models.OpenManyPacks
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			log.Fatal("Error unmarshaling response: ", err)
		}
		return result.Results
	} else {
		var result models.OpenOnePack
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			log.Fatal("Error unmarshaling response: ", err)
		}
		return []models.OpenOnePack{result}
	}
}
