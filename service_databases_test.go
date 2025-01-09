package notion_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestDatabaseService(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}

	emoji := notion.Emoji("🎉")
	user := notion.NewPersonUser("some_id", "some@example.com")

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.DatabaseID
			want       *notion.Database
			wantErr    bool
			err        error
		}{
			{
				name:       "returns database by id",
				id:         "some_id",
				filePath:   "testdata/database_get.json",
				statusCode: http.StatusOK,
				want: &notion.Database{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeDatabase,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					AtomLastEdited: notion.AtomLastEdited{
						LastEditedTime: &timestamp,
						LastEditedBy:   user,
					},

					Title: notion.RichTexts{
						{
							Type:        notion.RichTextTypeText,
							Text:        &notion.Text{Content: "Test Database"},
							Annotations: &notion.Annotations{Color: "default"},
							PlainText:   "Test Database",
							Href:        "",
						},
					},
					//Properties: notion.PropertyConfigs{
					//	"Tags": notion.MultiSelectPropertyConfig{
					//		ID:          ";s|V",
					//		Type:        notion.PropertyConfigTypeMultiSelect,
					//		MultiSelect: notion.Select{Options: []notion.Option{{ID: "id", Name: "tag", Color: "Blue"}}},
					//	},
					//	"Some another column": notion.PeoplePropertyConfig{
					//		ID:   "rJt\\",
					//		Type: notion.PropertyConfigTypePeople,
					//	},
					//
					//	"Name": notion.TitlePropertyConfig{
					//		ID:    "title",
					//		Type:  notion.PropertyConfigTypeTitle,
					//		Title: notion.RichText{},
					//	},
					//},
				},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := notion.NewDatabasesService(client).Get(context.Background(), tt.id)

				if (err != nil) != tt.wantErr {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				// TODO: remove properties from comparing for a while. Have to compare with interface somehow
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Get() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Query", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.DatabaseID
			request    *notion.DatabaseQueryRequest
			want       *notion.DatabaseQueryResponse
			wantErr    bool
			err        error
		}{
			{
				name:       "returns query results",
				id:         "some_id",
				filePath:   "testdata/database_query.json",
				statusCode: http.StatusOK,
				request: &notion.DatabaseQueryRequest{
					Filter: &notion.PropertyFilter{
						Property: "Name",
						RichText: &notion.TextFilterCondition{
							Contains: "Hel",
						},
					},
				},
				want: &notion.DatabaseQueryResponse{
					AtomPaginatedResponse: notion.AtomPaginatedResponse{
						Object:     notion.ObjectTypeList,
						HasMore:    false,
						NextCursor: notion.EmptyCursor,
					},
					Results: notion.Pages{
						{
							AtomObject: notion.AtomObject{
								Object: notion.ObjectTypePage,
							},
							AtomID: notion.AtomID{
								ID: "some_id",
							},
							AtomCreated: notion.AtomCreated{
								CreatedTime: &timestamp,
								CreatedBy:   user,
							},
							AtomLastEdited: notion.AtomLastEdited{
								LastEditedTime: &timestamp,
								LastEditedBy:   user,
							},
							AtomParent: notion.AtomParent{
								Parent: notion.NewDatabaseParent("some_id"),
							},
							AtomArchived: notion.AtomArchived{
								Archived: false,
							},
							AtomURLs: notion.AtomURLs{
								URL: "some_url",
							},
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := notion.NewDatabasesService(client).Query(context.Background(), tt.id, tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				got.Results[0].Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Query() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.DatabaseID
			request    *notion.DatabaseUpdateRequest
			want       *notion.Database
			wantErr    bool
			err        error
		}{
			{
				name:       "returns update results",
				filePath:   "testdata/database_update.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				request: &notion.DatabaseUpdateRequest{
					Title: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "patch"},
						},
					},
					Properties: notion.PropertyConfigs{
						"patch": notion.TitlePropertyConfig{
							Type: notion.PropertyConfigTypeRichText,
						},
					},
				},
				want: &notion.Database{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeDatabase,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					AtomLastEdited: notion.AtomLastEdited{
						LastEditedTime: &timestamp,
						LastEditedBy:   user,
					},
					AtomParent: notion.AtomParent{
						Parent: notion.NewPageParent("48f8fee9-cd79-4180-bc2f-ec0398253067"),
					},
					AtomArchived: notion.AtomArchived{
						Archived: false,
					},
					AtomAppearance: notion.AtomAppearance{
						Icon: &notion.Icon{
							Type:  "emoji",
							Emoji: &emoji,
						},
						Cover: &notion.Image{
							Type: "external",
							External: &notion.FileObject{
								URL: "https://website.domain/images/image.png",
							},
						},
					},
					Title: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "patch"},
						},
					},
					Description: []notion.RichText{},
					IsInline:    false,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := notion.NewDatabasesService(client).Update(context.Background(), tt.id, tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Update() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			request    *notion.DatabaseCreateRequest
			want       *notion.Database
			wantErr    bool
			err        error
		}{
			{
				name:       "returns created db",
				filePath:   "testdata/database_create.json",
				statusCode: http.StatusOK,
				request: &notion.DatabaseCreateRequest{
					Parent: notion.Parent{
						Type:   notion.ParentTypePageID,
						PageID: "some_id",
					},
					Title: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "Grocery List"},
						},
					},
					Properties: notion.PropertyConfigs{
						"create": notion.TitlePropertyConfig{
							Type: notion.PropertyConfigTypeTitle,
						},
					},
					IsInline: false,
				},
				want: &notion.Database{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeDatabase,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					AtomLastEdited: notion.AtomLastEdited{
						LastEditedTime: &timestamp,
						LastEditedBy:   user,
					},
					AtomParent: notion.AtomParent{
						Parent: notion.NewPageParent("a7744006-9233-4cd0-bf44-3a49de2c01b5"),
					},
					AtomArchived: notion.AtomArchived{
						Archived: false,
					},
					AtomAppearance: notion.AtomAppearance{
						Icon: &notion.Icon{
							Type:  "emoji",
							Emoji: &emoji,
						},
						Cover: &notion.Image{
							Type: "external",
							External: &notion.FileObject{
								URL: "https://website.domain/images/image.png",
							},
						},
					},
					Title: []notion.RichText{
						{
							Type:        notion.RichTextTypeText,
							Text:        &notion.Text{Content: "Grocery List"},
							PlainText:   "Grocery List",
							Annotations: &notion.Annotations{Color: notion.ColorDefault},
						},
					},
					Description: []notion.RichText{},
					IsInline:    false,
				},
			},
			{
				name:       "returns created db 2",
				filePath:   "testdata/database_create_2.json",
				statusCode: http.StatusOK,
				request: &notion.DatabaseCreateRequest{
					Parent: notion.Parent{
						Type:   notion.ParentTypePageID,
						PageID: "some_id",
					},
					Title: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "Grocery List"},
						},
					},
					Properties: notion.PropertyConfigs{
						"create": notion.TitlePropertyConfig{
							Type: notion.PropertyConfigTypeTitle,
						},
					},
					IsInline: false,
				},
				want: &notion.Database{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeDatabase,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					AtomLastEdited: notion.AtomLastEdited{
						LastEditedTime: &timestamp,
						LastEditedBy:   user,
					},
					AtomParent: notion.AtomParent{
						Parent: notion.NewBlockParent("a7744006-9233-4cd0-bf44-3a49de2c01b5"),
					},
					AtomArchived: notion.AtomArchived{
						Archived: false,
					},
					AtomAppearance: notion.AtomAppearance{
						Icon: &notion.Icon{
							Type:  "emoji",
							Emoji: &emoji,
						},
						Cover: &notion.Image{
							Type: "external",
							External: &notion.FileObject{
								URL: "https://website.domain/images/image.png",
							},
						},
					},

					Title: notion.RichTexts{
						notion.NewTextRichText("Grocery List").WithColor(notion.ColorDefault),
					},
					Description: notion.RichTexts{},
					IsInline:    false,
				},
			},
			{
				name:       "returns created db 3",
				filePath:   "testdata/database_create_3.json",
				statusCode: http.StatusOK,
				request: &notion.DatabaseCreateRequest{
					Parent: notion.Parent{
						Type:   notion.ParentTypePageID,
						PageID: "some_id",
					},
					Title: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "Grocery List"},
						},
					},
					Properties: notion.PropertyConfigs{
						"create": notion.TitlePropertyConfig{
							Type: notion.PropertyConfigTypeTitle,
						},
					},
					IsInline: true,
				},
				want: &notion.Database{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeDatabase,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					AtomLastEdited: notion.AtomLastEdited{
						LastEditedTime: &timestamp,
						LastEditedBy:   user,
					},
					AtomParent: notion.AtomParent{
						Parent: notion.NewPageParent("a7744006-9233-4cd0-bf44-3a49de2c01b5"),
					},
					AtomArchived: notion.AtomArchived{
						Archived: false,
					},
					AtomAppearance: notion.AtomAppearance{
						Icon: &notion.Icon{
							Type:  "emoji",
							Emoji: &emoji,
						},
						Cover: &notion.Image{
							Type: "external",
							External: &notion.FileObject{
								URL: "https://website.domain/images/image.png",
							},
						},
					},

					Title: notion.RichTexts{
						notion.NewTextRichText("Grocery List").WithColor(notion.ColorDefault),
					},
					Description: []notion.RichText{},
					IsInline:    true,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := notion.NewDatabasesService(client).Create(context.Background(), tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				got.Properties = nil
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Get with empty database_id", func(t *testing.T) {
		client := notion.NewClient("some_token")
		_, err := notion.NewDatabasesService(client).Get(context.TODO(), notion.DatabaseID(""))
		if err.Error() != "empty database id" {
			t.Error("database id is required error is expected")
		}
	})
}

func TestDatabaseQueryRequest_MarshalJSON(t *testing.T) {
	timeObj, err := time.Parse(time.RFC3339, "2021-05-10T02:43:42Z")
	if err != nil {
		t.Error(err)
		return
	}
	dateObj := notion.Date(timeObj)
	tests := []struct {
		name    string
		req     *notion.DatabaseQueryRequest
		want    []byte
		wantErr bool
	}{
		{
			name: "timestamp created",
			req: &notion.DatabaseQueryRequest{
				Filter: &notion.TimestampFilter{
					Timestamp: notion.TimestampCreated,
					CreatedTime: &notion.DateFilterCondition{
						NextWeek: &struct{}{},
					},
				},
			},
			want: []byte(`{"filter":{"timestamp":"created_time","created_time":{"next_week":{}}}}`),
		},
		{
			name: "timestamp last edited",
			req: &notion.DatabaseQueryRequest{
				Filter: &notion.TimestampFilter{
					Timestamp: notion.TimestampLastEdited,
					LastEditedTime: &notion.DateFilterCondition{
						Before: &dateObj,
					},
				},
			},
			want: []byte(`{"filter":{"timestamp":"last_edited_time","last_edited_time":{"before":"2021-05-10T02:43:42Z"}}}`),
		},
		{
			name: "or compound filter one level",
			req: &notion.DatabaseQueryRequest{
				Filter: notion.OrCompoundFilter{
					notion.PropertyFilter{
						Property: "Status",
						Select: &notion.SelectFilterCondition{
							Equals: "Reading",
						},
					},
					notion.PropertyFilter{
						Property: "Publisher",
						Select: &notion.SelectFilterCondition{
							Equals: "NYT",
						},
					},
				},
			},
			want: []byte(`{"filter":{"or":[{"property":"Status","select":{"equals":"Reading"}},{"property":"Publisher","select":{"equals":"NYT"}}]}}`),
		},
		{
			name: "and compound filter one level",
			req: &notion.DatabaseQueryRequest{
				Filter: notion.AndCompoundFilter{
					notion.PropertyFilter{
						Property: "Status",
						Select: &notion.SelectFilterCondition{
							Equals: "Reading",
						},
					},
					notion.PropertyFilter{
						Property: "Publisher",
						Select: &notion.SelectFilterCondition{
							Equals: "NYT",
						},
					},
				},
			},
			want: []byte(`{"filter":{"and":[{"property":"Status","select":{"equals":"Reading"}},{"property":"Publisher","select":{"equals":"NYT"}}]}}`),
		},
		{
			name: "compound filter two levels",
			req: &notion.DatabaseQueryRequest{
				Filter: notion.OrCompoundFilter{
					notion.PropertyFilter{
						Property: "Description",
						RichText: &notion.TextFilterCondition{
							Contains: "fish",
						},
					},
					notion.AndCompoundFilter{
						notion.PropertyFilter{
							Property: "Food group",
							Select: &notion.SelectFilterCondition{
								Equals: "🥦Vegetable",
							},
						},
						notion.PropertyFilter{
							Property: "Is protein rich?",
							Checkbox: &notion.CheckboxFilterCondition{
								Equals: true,
							},
						},
					},
				},
			},
			want: []byte(`{"filter":{"or":[{"property":"Description","rich_text":{"contains":"fish"}},{"and":[{"property":"Food group","select":{"equals":"🥦Vegetable"}},{"property":"Is protein rich?","checkbox":{"equals":true}}]}]}}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.req.MarshalJSON()
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
