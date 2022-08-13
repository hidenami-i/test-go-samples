package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("start")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	spreadSheetId := os.Getenv("SPREAD_SHEET_ID")
	fmt.Println(spreadSheetId)

	ctx := context.Background()
	srv, err := NewSpreadSheetService(ctx, spreadSheetId)

	ss, err := srv.Spreadsheets.Get(spreadSheetId).Do()
	if err != nil {
		log.Fatalf("Unable to get spread sheet: %v", err)
	}

	for _, sheet := range ss.Sheets {
		if strings.Contains(sheet.Properties.Title, "[T]") {
			fmt.Println(sheet.Properties.Title)
			if err != nil {
				log.Fatalf("Unable to get values: %v", err)
			}
		}
		if strings.Contains(sheet.Properties.Title, "[D]") {
			//fmt.Println(sheet.Properties.Title)
		}
	}
}

func NewSpreadSheetService(ctx context.Context, spreadSheetId string) (*sheets.Service, error) {
	b, err := os.ReadFile("secret.json")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	jwt, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to jwt config from json: %v", err)
	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(jwt.Client(ctx)))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return srv, nil
}
