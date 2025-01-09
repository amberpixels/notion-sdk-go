package notion_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestBlocksUnmarshal(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-11-04T02:09:00Z")
	if err != nil {
		t.Fatal(err)
	}

	var emoji notion.Emoji = "📌"
	user := notion.NewPersonUser("some_id", "some_user@example.com")

	t.Run("BlockArray", func(t *testing.T) {
		tests := []struct {
			name     string
			filePath string
			want     notion.Blocks
			wantErr  bool
			err      error
		}{
			{
				name:     "unmarshal",
				filePath: "testdata/block_array_unmarshal.json",
				want: notion.Blocks{
					&notion.CalloutBlock{
						BaseBlock: notion.BaseBlock{
							AtomObject: notion.AtomObject{
								Object: "block",
							},
							AtomID: notion.AtomID{
								ID: "block1",
							},
							AtomCreated: notion.AtomCreated{
								CreatedTime: &timestamp,
								CreatedBy:   user,
							},
							AtomLastEdited: notion.AtomLastEdited{
								LastEditedTime: &timestamp,
								LastEditedBy:   user,
							},
							Type: notion.BlockTypeCallout,
						},
						Callout: notion.Callout{
							RichText: notion.RichTexts{
								notion.NewTextRichText("This page is designed to be shared with students on the web. Click ").
									WithColor(notion.ColorDefault),

								notion.NewTextRichText("Share").WithCode().WithColor(notion.ColorDefault),
							},
							Icon: notion.NewEmojiIcon(emoji),

							Color: notion.ColorBlue.String(),
						},
					},
					&notion.Heading1Block{
						BaseBlock: notion.BaseBlock{
							Object:         "block",
							ID:             "block2",
							Type:           "heading_1",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
						},
						Heading1: notion.Heading{
							RichText: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "History 340",
									},
									Annotations: &notion.Annotations{
										Color: "default",
									},
									PlainText: "History 340",
								},
							},
							Color: notion.ColorBrownBackground.String(),
						},
					},
					&notion.ChildDatabaseBlock{
						BaseBlock: notion.BaseBlock{
							Object:         "block",
							ID:             "block3",
							Type:           "child_database",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
						},
						ChildDatabase: struct {
							Title string "json:\"title\""
						}{
							Title: "Required Texts",
						},
					},
					&notion.ColumnListBlock{
						BaseBlock: notion.BaseBlock{
							Object:         "block",
							ID:             "block4",
							Type:           "column_list",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
							HasChildren:    true,
						},
					},
					&notion.Heading3Block{
						BaseBlock: notion.BaseBlock{
							Object:         "block",
							ID:             "block5",
							Type:           "heading_3",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
						},
						Heading3: notion.Heading{
							RichText: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "Assignment Submission",
									},
									Annotations: &notion.Annotations{
										Bold:  true,
										Color: "default",
									},
									PlainText: "Assignment Submission",
								},
							},
							Color: notion.ColorDefault.String(),
						},
					},
					&notion.ParagraphBlock{
						BaseBlock: notion.BaseBlock{
							Object:         "block",
							ID:             "block6",
							Type:           "paragraph",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
						},
						Paragraph: notion.Paragraph{
							RichText: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "All essays and papers are due in lecture (due dates are listed on the schedule). No electronic copies will be accepted!",
									},
									Annotations: &notion.Annotations{
										Color: "default",
									},
									PlainText: "All essays and papers are due in lecture (due dates are listed on the schedule). No electronic copies will be accepted!",
								},
							},
							Color: notion.ColorRed.String(),
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				data, err := os.ReadFile(tt.filePath)
				if err != nil {
					t.Fatal(err)
				}
				got := make(notion.Blocks, 0)
				err = json.Unmarshal(data, &got)
				if err != nil {
					t.Fatal(err)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
