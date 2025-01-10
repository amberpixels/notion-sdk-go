package notion_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestBlocksUnmarshal(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-11-04T02:09:00Z")
	assert.NoError(t, err, "Failed to parse timestamp")

	icon := notion.NewEmojiIcon("📌")
	user := notion.NewPersonUser("some_id", "some_user@example.com")

	t.Run("BlockArray", func(t *testing.T) {
		tests := []struct {
			name     string
			filePath string
			want     notion.Blocks
			wantErr  bool
		}{
			{
				name:     "unmarshal",
				filePath: "testdata/block_array_unmarshal.json",
				want: notion.Blocks{
					decorateTestBasicBlock(notion.NewCalloutBlock(
						notion.Callout{
							RichText: notion.RichTexts{
								notion.NewTextRichText("This page is designed to be shared with students on the web. Click ").
									WithColor(notion.ColorDefault),
								notion.NewTextRichText("Share").WithCode().WithColor(notion.ColorDefault),
							},
							Icon:  icon,
							Color: notion.ColorBlue,
						},
					), "block1", &timestamp, user),

					decorateTestBasicBlock(notion.NewHeading1Block(
						notion.Heading{
							RichText: notion.RichTexts{
								notion.NewTextRichText("History 340"),
							},
							Color: notion.ColorBrownBackground,
						},
					), "block2", &timestamp, user),

					decorateTestBasicBlock(
						notion.NewChildDataBasicBlock("Required Texts"),
						"block3",
						&timestamp,
						user,
					),

					decorateTestBasicBlock(
						notion.NewColumnListBlock(notion.ColumnList{}),
						"block4",
						&timestamp,
						user,
					),

					decorateTestBasicBlock(
						notion.NewHeading3Block(notion.Heading{
							RichText: notion.RichTexts{
								notion.NewTextRichText("Assignment Submission"),
							},
							Color: notion.ColorDefault,
						}),
						"block5",
						&timestamp,
						user,
					),

					decorateTestBasicBlock(
						notion.NewParagraphBlock(notion.Paragraph{
							RichText: notion.RichTexts{
								notion.NewTextRichText("All essays and papers are due in lecture (due dates are listed on the schedule). No electronic copies will be accepted!"),
							},
							Color: notion.ColorRed,
						}),
						"block6",
						&timestamp,
						user,
					),
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				data, err := os.ReadFile(tt.filePath)
				assert.NoError(t, err, "Failed to read file")

				got := make(notion.Blocks, 0)
				err = json.Unmarshal(data, &got)
				assert.NoError(t, err, "Failed to unmarshal JSON")

				assert.Equal(t, tt.want, got, "Unmarshaled blocks do not match expected result")
			})
		}
	})
}
