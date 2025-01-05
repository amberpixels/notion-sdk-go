package notion_test

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestBlockClient(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("GetChildren", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			len        int
			wantErr    bool
			err        error
		}{
			{
				name:       "returns blocks by id of parent block",
				id:         "some_id",
				statusCode: http.StatusOK,
				filePath:   "testdata/block_get_children.json",
				len:        2,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Block.GetChildren(context.Background(), tt.id, nil)

				if (err != nil) != tt.wantErr {
					t.Errorf("GetChildren() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if tt.len != len(got.Results) {
					t.Errorf("GetChildren got %d, want: %d", len(got.Results), tt.len)
				}
			})
		}
	})

	t.Run("AppendChildren", func(t *testing.T) {

		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			request    *notion.AppendBlockChildrenRequest
			want       *notion.AppendBlockChildrenResponse
			wantErr    bool
			err        error
		}{
			{
				name:       "return list object",
				id:         "some_id",
				filePath:   "testdata/block_append_children.json",
				statusCode: http.StatusOK,
				request: &notion.AppendBlockChildrenRequest{
					Children: []notion.Block{
						&notion.Heading2Block{
							BasicBlock: notion.BasicBlock{
								Object: notion.ObjectTypeBlock,
								Type:   notion.BlockTypeHeading2,
							},
							Heading2: struct {
								RichText     []notion.RichText `json:"rich_text"`
								Children     notion.Blocks     `json:"children,omitempty"`
								Color        string            `json:"color,omitempty"`
								IsToggleable bool              `json:"is_toggleable,omitempty"`
							}{[]notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "Hello"},
								},
							}, nil, "", false,
							},
						},
					},
				},
				want: &notion.AppendBlockChildrenResponse{
					Object: notion.ObjectTypeList,
					Results: []notion.Block{&notion.ParagraphBlock{
						BasicBlock: notion.BasicBlock{
							Object:         notion.ObjectTypeBlock,
							ID:             "some_id",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							Type:           notion.BlockTypeParagraph,
							CreatedBy: &notion.User{
								Object: "user",
								ID:     "some_id",
							},
							LastEditedBy: &notion.User{
								Object: "user",
								ID:     "some_id",
							},
						},
						Paragraph: notion.Paragraph{
							RichText: []notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "AAAAAA"},
									Annotations: &notion.Annotations{
										Bold:  true,
										Color: notion.ColorDefault,
									},
									PlainText: "AAAAAA",
								},
							},
							Color: "blue",
						},
					}},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Block.AppendChildren(context.Background(), tt.id, tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("AppendChildren() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				a, err := json.Marshal(got)
				if err != nil {
					t.Errorf("AppendChildren() marshal error = %v", err)
					return
				}
				b, err := json.Marshal(tt.want)
				if err != nil {
					t.Errorf("AppendChildren() marshal error = %v", err)
					return
				}

				if !(string(a) == string(b)) {
					t.Errorf("AppendChildren() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			want       notion.Block
			wantErr    bool
			err        error
		}{
			{
				name:       "returns block object",
				filePath:   "testdata/block_get.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				want: &notion.ChildPageBlock{
					BasicBlock: notion.BasicBlock{
						Object:         notion.ObjectTypeBlock,
						ID:             "some_id",
						Type:           notion.BlockTypeChildPage,
						CreatedTime:    &timestamp,
						LastEditedTime: &timestamp,
						CreatedBy: &notion.User{
							Object: "user",
							ID:     "some_id",
						},
						LastEditedBy: &notion.User{
							Object: "user",
							ID:     "some_id",
						},
						HasChildren: true,
						Parent: &notion.Parent{
							Type:   "page_id",
							PageID: "59833787-2cf9-4fdf-8782-e53db20768a5",
						},
					},
					ChildPage: struct {
						Title string `json:"title"`
					}{
						Title: "Hello",
					},
				},
				wantErr: false,
				err:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Block.Get(context.Background(), tt.id)

				if (err != nil) != tt.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			req        *notion.BlockUpdateRequest
			want       notion.Block
			wantErr    bool
			err        error
		}{
			{
				name:       "updates block and returns it",
				filePath:   "testdata/block_update.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				req: &notion.BlockUpdateRequest{
					Paragraph: &notion.Paragraph{
						RichText: []notion.RichText{
							{
								Text: &notion.Text{Content: "Hello"},
							},
						},
						Color: notion.ColorYellow.String(),
					},
				},
				want: &notion.ParagraphBlock{
					BasicBlock: notion.BasicBlock{
						Object:         notion.ObjectTypeBlock,
						ID:             "some_id",
						Type:           notion.BlockTypeParagraph,
						CreatedTime:    &timestamp,
						LastEditedTime: &timestamp,
					},
					Paragraph: notion.Paragraph{
						RichText: []notion.RichText{
							{
								Type: notion.RichTextTypeText,
								Text: &notion.Text{
									Content: "Hello",
								},
								Annotations: &notion.Annotations{Color: notion.ColorDefault},
								PlainText:   "Hello",
							},
						},
						Color: notion.ColorYellow.String(),
					},
				},
				wantErr: false,
				err:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Block.Update(context.Background(), tt.id, tt.req)

				if (err != nil) != tt.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestBlockArrayUnmarshal(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-11-04T02:09:00Z")
	if err != nil {
		t.Fatal(err)
	}

	var emoji notion.Emoji = "📌"
	var user *notion.User = &notion.User{
		Object: "user",
		ID:     "some_id",
	}
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
						BasicBlock: notion.BasicBlock{
							Object:         "block",
							ID:             "block1",
							Type:           "callout",
							CreatedTime:    &timestamp,
							LastEditedTime: &timestamp,
							CreatedBy:      user,
							LastEditedBy:   user,
						},
						Callout: notion.Callout{
							RichText: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "This page is designed to be shared with students on the web. Click ",
									},
									Annotations: &notion.Annotations{
										Color: "default",
									},
									PlainText: "This page is designed to be shared with students on the web. Click ",
								}, {
									Type: "text",
									Text: &notion.Text{
										Content: "Share",
									},
									Annotations: &notion.Annotations{
										Code:  true,
										Color: "default",
									},
									PlainText: "Share",
								},
							},
							Icon: &notion.Icon{
								Type:  "emoji",
								Emoji: &emoji,
							},
							Color: notion.ColorBlue.String(),
						},
					},
					&notion.Heading1Block{
						BasicBlock: notion.BasicBlock{
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
						BasicBlock: notion.BasicBlock{
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
						BasicBlock: notion.BasicBlock{
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
						BasicBlock: notion.BasicBlock{
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
						BasicBlock: notion.BasicBlock{
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

func TestBlockUpdateRequest_MarshallJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     *notion.BlockUpdateRequest
		want    []byte
		wantErr bool
	}{
		{
			name: "update todo checkbox",
			req: &notion.BlockUpdateRequest{
				ToDo: &notion.ToDo{Checked: false, RichText: make([]notion.RichText, 0)},
			},
			want: []byte(`{"to_do":{"rich_text":[],"checked":false}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}
