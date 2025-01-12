package examples_test

import (
	"context"
	"log"

	"github.com/amberpixels/notion-sdk-go"
)

func ExampleBlocksService_GetChildren() {
	ctx := context.Background()

	envCfg := loadEnv()

	client := notion.New(envCfg.Token)

	blocks, err := client.Blocks.GetChildren(ctx, envCfg.PageID, nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = blocks
	// Output:
}
