package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"time"

	discordwebhook "github.com/dozerokz/discord-webhook-go"
	_ "github.com/joho/godotenv/autoload"

	"auto-kyara/models"
)

var baseUrl = "https://kyara.games/api"

var AUTH_TOKEN = os.Getenv("AUTH_TOKEN")
var WEBHOOK_URL = os.Getenv("WEBHOOK_URL")
var DISCORD_ID = os.Getenv("DISCORD_ID")

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
		url = baseUrl + "/packs/all/open-many"
	} else {
		url = baseUrl + "/packs/all/open"
	}

	body := map[string]int{
		"maxChakra": 0,
	}
	if amount > 1 {
		body["count"] = amount
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

func CreateCardEmbed(card models.Card) *discordwebhook.Embed {
	notifReason := "New card"
	notifColor, _ := strconv.ParseInt("5a5662", 16, 64)
	rariryName := "Common"
	if card.Rarity == "U" {
		notifColor, _ = strconv.ParseInt("3fbf7f", 16, 64)
		rariryName = "Uncommon"
	}
	if card.Rarity == "R" {
		notifColor, _ = strconv.ParseInt("4d9bff", 16, 64)
		rariryName = "Rare"
	}
	if card.Rarity == "E" {
		notifReason = "Epic drop"
		notifColor, _ = strconv.ParseInt("b072ff", 16, 64)
		rariryName = "Epic"
	}
	if card.Rarity == "L" {
		notifReason = "Legendary drop"
		notifColor, _ = strconv.ParseInt("f5c518", 16, 64)
		rariryName = "Legendary"
	}
	if card.Rarity == "M" {
		notifReason = "Mythical drop"
		notifColor, _ = strconv.ParseInt("ff4d2e", 16, 64)
		rariryName = "Mythical"
	}

	status := "Already owned"
	if card.IsNew {
		status = "New card"
	}

	return &discordwebhook.Embed{
		Title:       card.Name,
		Description: notifReason,
		Image:       discordwebhook.CreateImage(card.Image, "", 0, 0),
		Color:       int(notifColor),
		Fields: []discordwebhook.Field{
			{
				Name:   "Id",
				Value:  strconv.Itoa(card.ID),
				Inline: true,
			},
			{
				Name:   "Rarity",
				Value:  rariryName,
				Inline: true,
			},
			{
				Name:   "Series",
				Value:  card.Series,
				Inline: true,
			},
			{
				Name:   "Rank",
				Value:  rariryName,
				Inline: true,
			},
			{
				Name:   "Power / Aura",
				Value:  strconv.Itoa(card.Pow) + " / " + strconv.Itoa(card.Aura),
				Inline: true,
			},
			{
				Name:   "Status",
				Value:  status,
				Inline: true,
			},
		},
	}
}

func PostEmbed(embed *discordwebhook.Embed, ping bool) {
	webhook, err := discordwebhook.CreateWebhook("", "Kyara", "")
	if err != nil {
		log.Fatal("Error creating webhook: ", err)
	}
	if ping {
		pingWebhook, err := discordwebhook.CreateWebhook("<@"+DISCORD_ID+">", "Kyara", "")
		if err != nil {
			log.Fatal("Error creating pingWebhook: ", err)
		}
		discordwebhook.SendWebhook(WEBHOOK_URL, pingWebhook)
	}
	webhook.Embeds = append(webhook.Embeds, *embed)
	err = discordwebhook.SendWebhook(WEBHOOK_URL, webhook)
	if err != nil {
		log.Fatal("Error sending webhook: ", err)
	}
}

func PostNextOpeningWebhook(nextTime time.Time) {
	timeWebhook, err := discordwebhook.CreateWebhook("", "Kyara", "")
	if err != nil {
		log.Fatal("Error creating webhook: ", err)
	}
	notifColor, _ := strconv.ParseInt("2c0055", 16, 64)
	timeWebhook.Embeds = append(timeWebhook.Embeds, discordwebhook.Embed{
		Title:       "Next opening",
		Description: "Next opening at " + nextTime.Format("15h04") + " <t:" + strconv.FormatInt(nextTime.Unix(), 10) + ":R>",
		Color:       int(notifColor),
	})
	err = discordwebhook.SendWebhook(WEBHOOK_URL, timeWebhook)
	if err != nil {
		log.Fatal("Error sending webhook: ", err)
	}
}

func main() {
	meTest := FetchMe()
	fmt.Println("Name : ", meTest.Discord.Username)

	if AUTH_TOKEN == "" {
		log.Fatal("Error loading .env file")
	}

	for {
		sleepTime := 15 + rand.IntN(90-15+1)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "🕛 Sleeping for", sleepTime, "minutes... Next opening will be at", time.Now().Add(time.Minute*time.Duration(sleepTime)).Format("2006-01-02 15:04:05"))
		PostNextOpeningWebhook(time.Now().Add(time.Minute * time.Duration(sleepTime)))

		time.Sleep(time.Duration(sleepTime) * time.Minute)

		me := FetchMe()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Free packs : ", me.FreePacks)
		if me.FreePacks >= 1 {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Opening packs...")
			result := OpenPacks(me.FreePacks)
			var cards []models.Card
			for _, openPackResult := range result {
				cards = append(cards, openPackResult.Cards...)
			}

			alreadyPinged := false
			for _, card := range cards {
				if card.Rarity == "M" || card.Rarity == "L" || card.Rarity == "E" || card.IsNew {
					PostEmbed(CreateCardEmbed(card), !alreadyPinged)
					alreadyPinged = true
				}
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}
