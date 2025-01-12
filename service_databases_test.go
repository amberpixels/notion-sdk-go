package notion_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseService(t *testing.T) {
	ctx := context.Background()

	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	assert.NoError(t, err, "Failed to parse timestamp")

	user := notion.NewPersonUser("some_id", "some@example.com")

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.DatabaseID
			want       *notion.Database
			wantErr    bool
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
						notion.NewTextRichText("Test Database"),
					},
				},
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Databases.Get(ctx, tt.id)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				got.Properties = nil // Skip comparing properties
				assert.Equal(t, tt.want, got, "Unexpected database result")
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
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Databases.Query(ctx, tt.id, tt.request)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				got.Results[0].Properties = nil // Skip comparing properties
				assert.Equal(t, tt.want, got, "Unexpected query result")
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
		}{
			{
				name:       "returns update results",
				filePath:   "testdata/database_update.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				request: &notion.DatabaseUpdateRequest{
					Title: notion.RichTexts{
						notion.NewTextRichText("patch"),
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
					Title: notion.RichTexts{
						notion.NewTextRichText("patch"),
					},
				},
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Databases.Update(ctx, tt.id, tt.request)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				got.Properties = nil // Skip comparing properties
				assert.Equal(t, tt.want, got, "Unexpected update result")
			})
		}
	})
}

func TestDatabaseQueryRequest_MarshalJSON(t *testing.T) {
	timeObj, err := time.Parse(time.RFC3339, "2021-05-10T02:43:42Z")
	assert.NoError(t, err, "Failed to parse time")

	dateObj := notion.Date(timeObj)

	tests := []struct {
		name    string
		req     *notion.DatabaseQueryRequest
		want    string
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
			want: `{"filter":{"timestamp":"created_time","created_time":{"next_week":{}}}}`,
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
			want: `{"filter":{"timestamp":"last_edited_time","last_edited_time":{"before":"2021-05-10T02:43:42Z"}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.req.MarshalJSON()
			assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
			assert.JSONEq(t, tt.want, string(got), "Unexpected JSON result")
		})
	}
}
