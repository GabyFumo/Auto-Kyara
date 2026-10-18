package main

import (
	"encoding/json"
	"fmt"
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
	user := FetchMe()
	fmt.Println(user)
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
