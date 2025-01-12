package examples_test

import (
	"context"
	"fmt"
	"log"

	"github.com/amberpixels/notion-sdk-go"
)

func ExamplePagesService_Get() {
	ctx := context.Background()

	envCfg := loadEnv()

	client := notion.New(envCfg.Token)

	page, err := client.Pages.Get(ctx, envCfg.PageID)
	if err != nil {
		log.Fatal(err)
	}

	// Output: page
	fmt.Println(page.GetObject())
}
