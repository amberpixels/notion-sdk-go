package examples_test

import (
	"log"
	"os"

	"github.com/amberpixels/notion-sdk-go"

	"github.com/joho/godotenv"
)

type ExamplesConfig struct {
	Token  notion.Token
	PageID notion.ObjectID
}

func loadEnv() ExamplesConfig {
	err := godotenv.Load(".env")
	if os.IsNotExist(err) {
		// having `.env` is optional, so we're OK here
	} else if err != nil {
		log.Fatal("failed to read .env")
	}

	token := os.Getenv("NOTION_API_TOKEN")
	if token == "" {
		log.Fatal("NOTION_API_TOKEN env var is not set")
	}

	examplePageID := os.Getenv("NOTION_EXAMPLE_PAGE_ID")
	if examplePageID == "" {
		log.Fatal("NOTION_EXAMPLE_PAGE_ID env var is required")
	}

	return ExamplesConfig{
		Token:  notion.Token(token),
		PageID: notion.ObjectID(examplePageID),
	}
}
