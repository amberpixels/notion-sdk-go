package notion_test

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestPageClient(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.PageID
			want       *notion.Page
			wantErr    bool
			err        error
		}{
			{
				name:       "returns page by id",
				id:         "some_id",
				filePath:   "testdata/page_get.json",
				statusCode: http.StatusOK,
				want: &notion.Page{
					Object:         notion.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					CreatedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					LastEditedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					Parent: notion.Parent{
						Type:       notion.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
					URL:      "some_url",
					Properties: notion.Properties{
						"Tags": &notion.MultiSelectProperty{
							ID:   ";s|V",
							Type: "multi_select",
							MultiSelect: []notion.Option{
								{
									ID:    "some_id",
									Name:  "tag",
									Color: "blue",
								},
							},
						},
						"Some another column": &notion.PeopleProperty{
							ID:   "rJt\\",
							Type: "people",
							People: []notion.User{
								{
									Object:    "user",
									ID:        "some_id",
									Name:      "some name",
									AvatarURL: "some.url",
									Type:      "person",
									Person: &notion.Person{
										Email: "some@email.com",
									},
								},
							},
						},
						"SomeColumn": &notion.RichTextProperty{
							ID:   "~j_@",
							Type: "rich_text",
							RichText: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "some text",
									},
									Annotations: &notion.Annotations{
										Color: "default",
									},
									PlainText: "some text",
								},
							},
						},
						"Some Files": &notion.FilesProperty{
							ID:   "files",
							Type: "files",
							Files: []notion.File{
								{
									Name: "https://google.com",
									Type: "external",
									External: &notion.FileObject{
										URL: "https://google.com",
									},
								},
							},
						},
						"Name": &notion.TitleProperty{
							ID:   "title",
							Type: "title",
							Title: []notion.RichText{
								{
									Type: "text",
									Text: &notion.Text{
										Content: "Hello",
									},
									Annotations: &notion.Annotations{
										Color: "default",
									},
									PlainText: "Hello",
								},
							},
						},
						"RollupArray": &notion.RollupProperty{
							ID:   "abcd",
							Type: "rollup",
							Rollup: notion.Rollup{
								Type: "array",
								Array: notion.PropertyArray{
									&notion.NumberProperty{
										Type:   "number",
										Number: 42.2,
									},
									&notion.NumberProperty{
										Type:   "number",
										Number: 56,
									},
								},
							},
						},
					},
				},
			},
			{
				name:       "returns validation error for invalid request",
				id:         "some_id",
				filePath:   "testdata/validation_error.json",
				statusCode: http.StatusBadRequest,
				wantErr:    true,
				err: &notion.Error{
					Object:  notion.ObjectTypeError,
					Status:  http.StatusBadRequest,
					Code:    "validation_error",
					Message: "The provided page ID is not a valid Notion UUID: bla bla.",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := client.Page.Get(context.Background(), tt.id)
				if err != nil {
					if tt.wantErr {
						if !reflect.DeepEqual(err, tt.err) {
							t.Errorf("Get error() got = %v, want %v", err, tt.err)
						}
					} else {
						t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)

					}
					return
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.PageID
			request    *notion.PageCreateRequest
			want       *notion.Page
			wantErr    bool
			err        error
		}{
			{
				name:       "returns a new page",
				filePath:   "testdata/page_create.json",
				statusCode: http.StatusOK,
				request: &notion.PageCreateRequest{
					Parent: notion.Parent{
						Type:       notion.ParentTypeDatabaseID,
						DatabaseID: "f830be5eff534859932e5b81542b3c7b",
					},
					Properties: notion.Properties{
						"Name": notion.TitleProperty{
							Title: []notion.RichText{
								{Text: &notion.Text{Content: "hello"}},
							},
						},
					},
				},
				want: &notion.Page{
					Object:         notion.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					CreatedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					LastEditedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					Parent: notion.Parent{
						Type:       notion.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
					URL:      "some_url",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Page.Create(context.Background(), tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.PageID
			request    *notion.PageUpdateRequest
			want       *notion.Page
			wantErr    bool
			err        error
		}{
			{
				name:       "change requested properties and return the result",
				id:         "some_id",
				filePath:   "testdata/page_update.json",
				statusCode: http.StatusOK,
				request: &notion.PageUpdateRequest{
					Properties: notion.Properties{
						"SomeColumn": notion.RichTextProperty{
							Type: notion.PropertyTypeRichText,
							RichText: []notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "patch"},
								},
							},
						},
						"Important Files": notion.FilesProperty{
							Type: "files",
							Files: []notion.File{
								{
									Type: "external",
									Name: "https://google.com",
									External: &notion.FileObject{
										URL: "https://google.com",
									},
								},
								{
									Type: "external",
									Name: "https://123.com",
									External: &notion.FileObject{
										URL: "https://123.com",
									},
								},
							},
						},
					},
				},
				want: &notion.Page{
					Object:         notion.ObjectTypePage,
					ID:             "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					CreatedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					LastEditedBy: notion.User{
						Object: "user",
						ID:     "some_id",
					},
					Parent: notion.Parent{
						Type:       notion.ParentTypeDatabaseID,
						DatabaseID: "some_id",
					},
					Archived: false,
					URL:      "some_url",
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Page.Update(context.Background(), tt.id, tt.request)
				if (err != nil) != tt.wantErr {
					t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Update() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestPageCreateRequest_MarshallJSON(t *testing.T) {
	timeObj, err := time.Parse(time.RFC3339, "2020-12-08T12:00:00Z")
	if err != nil {
		t.Error(err)
		return
	}

	dateObj := notion.Date(timeObj)
	tests := []struct {
		name    string
		req     *notion.PageCreateRequest
		want    []byte
		wantErr bool
	}{
		{
			name: "create a page",
			req: &notion.PageCreateRequest{
				Parent: notion.Parent{
					DatabaseID: "some_id",
				},
				Properties: notion.Properties{
					"Type": notion.SelectProperty{
						Select: notion.Option{
							ID:    "some_id",
							Name:  "Article",
							Color: notion.ColorDefault,
						},
					},
					"Name": notion.TitleProperty{
						Title: []notion.RichText{
							{Text: &notion.Text{Content: "New Media Article"}},
						},
					},
					"Publishing/Release Date": notion.DateProperty{
						Date: &notion.DateObject{
							Start: &dateObj,
						},
					},
					"Link": notion.URLProperty{
						URL: "some_url",
					},
					"Summary": notion.TextProperty{
						Text: []notion.RichText{
							{
								Type: notion.RichTextTypeText,
								Text: &notion.Text{
									Content: "Some content",
								},
								Annotations: &notion.Annotations{
									Bold:  true,
									Color: notion.ColorBlue,
								},
								PlainText: "Some content",
							},
						},
					},
					"Read": notion.CheckboxProperty{
						Checkbox: false,
					},
				},
			},
			want: []byte(`{"parent":{"database_id":"some_id"},"properties":{"Link":{"url":"some_url"},"Name":{"title":[{"text":{"content":"New Media Article"}}]},"Publishing/Release Date":{"date":{"start":"2020-12-08T12:00:00Z","end":null}},"Read":{"checkbox":false},"Summary":{"text":[{"type":"text","text":{"content":"Some content"},"annotations":{"bold":true,"italic":false,"strikethrough":false,"underline":false,"code":false,"color":"blue"},"plain_text":"Some content"}]},"Type":{"select":{"id":"some_id","name":"Article","color":"default"}}}}`),
		},
		{
			name: "create a page with content",
			req: &notion.PageCreateRequest{
				Parent: notion.Parent{
					DatabaseID: "some_id",
				},
				Properties: notion.Properties{
					"Type": notion.SelectProperty{
						Select: notion.Option{
							ID:    "some_id",
							Name:  "Article",
							Color: notion.ColorDefault,
						},
					},
					"Name": notion.TitleProperty{
						Title: []notion.RichText{
							{Text: &notion.Text{Content: "New Media Article"}},
						},
					},
					"Publishing/Release Date": notion.DateProperty{
						Date: &notion.DateObject{
							Start: &dateObj,
						},
					},
					"Link": notion.URLProperty{
						URL: "some_url",
					},
					"Summary": notion.TextProperty{
						Text: []notion.RichText{
							{
								Type: notion.RichTextTypeText,
								Text: &notion.Text{
									Content: "Some content",
								},
								Annotations: &notion.Annotations{
									Bold:  true,
									Color: notion.ColorBlue,
								},
								PlainText: "Some content",
							},
						},
					},
					"Read": notion.CheckboxProperty{
						Checkbox: false,
					},
				},
				Children: []notion.Block{
					&notion.Heading2Block{
						BasicBlock: notion.BasicBlock{
							Object: notion.ObjectTypeBlock,
							Type:   notion.BlockTypeHeading2,
						},
						Heading2: notion.Heading{
							RichText: []notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "Lacinato"},
								},
							},
						},
					},
					&notion.ParagraphBlock{
						BasicBlock: notion.BasicBlock{
							Object: notion.ObjectTypeBlock,
							Type:   notion.BlockTypeParagraph,
						},
						Paragraph: notion.Paragraph{
							RichText: []notion.RichText{
								{
									Text: &notion.Text{
										Content: "Lacinato",
										Link: &notion.Link{
											Url: "some_url",
										},
									},
								},
							},
							Children: nil,
						},
					},
				},
			},
			want: []byte(`{"parent":{"database_id":"some_id"},"properties":{"Link":{"url":"some_url"},"Name":{"title":[{"text":{"content":"New Media Article"}}]},"Publishing/Release Date":{"date":{"start":"2020-12-08T12:00:00Z","end":null}},"Read":{"checkbox":false},"Summary":{"text":[{"type":"text","text":{"content":"Some content"},"annotations":{"bold":true,"italic":false,"strikethrough":false,"underline":false,"code":false,"color":"blue"},"plain_text":"Some content"}]},"Type":{"select":{"id":"some_id","name":"Article","color":"default"}}},"children":[{"object":"block","type":"heading_2","heading_2":{"rich_text":[{"type":"text","text":{"content":"Lacinato"}}]}},{"object":"block","type":"paragraph","paragraph":{"rich_text":[{"text":{"content":"Lacinato","link":{"url":"some_url"}}}]}}]}`),
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

func TestPageUpdateRequest_MarshallJSON(t *testing.T) {
	tests := []struct {
		name    string
		req     *notion.PageUpdateRequest
		want    []byte
		wantErr bool
	}{
		{
			name: "update checkbox",
			req: &notion.PageUpdateRequest{
				Properties: map[string]notion.Property{
					"Checked": notion.CheckboxProperty{
						Checkbox: false,
					},
				},
			},
			want: []byte(`{"properties":{"Checked":{"checkbox":false}},"archived":false}`),
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
