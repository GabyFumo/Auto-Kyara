# Auto-Kyara

![Go](https://img.shields.io/badge/Go-1.21+-blue?style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/Docker-24.0.0+-blue?style=for-the-badge&logo=docker)
![Discord](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white)

## Introduction

This is a auto-claimer for a TCG site made with golang.

## Requirements

- Go 1.21+ (if you want to use Native version - [download](https://go.dev/dl/))
- Docker (recommanded version - [download](https://www.docker.com/products/docker-desktop/))

## How to use it

### Native
1. Clone the repository
2. Install the dependencies
3. Create a .env file with the following variables:
   - **AUTH_TOKEN** : Your auth token from the site (can be found in the storage tab)
   - **WEBHOOK_URL** : Your webhook url from discord
   - **DISCORD_ID** : Your discord id
> Check **.env.example** for an example of the .env file
4. Run the program with `go run main.go`

### Docker
1. Clone the repository
2. Create a .env file with the following variables:
   - **AUTH_TOKEN** : Your auth token from the site (can be found in the storage tab)
   - **WEBHOOK_URL** : Your webhook url from discord
   - **DISCORD_ID** : Your discord id
> Check **.env.example** for an example of the .env file
4. Run the program with `docker compose up`

## Next Features

- [ ] Ask which pack to open
- [ ] Better error handling

## How it works

The program will wake up every 15 to 90 minutes and check if you have any free packs. If you do, it will open them and send a message to your discord with the results.

## Disclaimer

I'm not responsible for any consequences that may come from using this program. Use it at your own risk.
