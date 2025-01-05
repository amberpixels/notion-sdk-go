package notion_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestCommentClient(t *testing.T) {
	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}

	var user = notion.User{
		Object: "user",
		ID:     "some_id",
	}

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			want       *notion.CommentQueryResponse
			wantErr    bool
			err        error
		}{
			{
				name:       "returns comments for given block",
				filePath:   "testdata/comment_get.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				want: &notion.CommentQueryResponse{
					Object: notion.ObjectTypeList,
					Results: []notion.Comment{
						{
							Object:         notion.ObjectTypeComment,
							ID:             "some_id",
							DiscussionID:   "some_id",
							CreatedTime:    timestamp,
							LastEditedTime: timestamp,
							CreatedBy:      user,
							Parent: notion.Parent{
								Type:   notion.ParentTypePageID,
								PageID: "some_id",
							},
							RichText: []notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "content"},
								},
							},
						},
						{
							Object:         notion.ObjectTypeComment,
							ID:             "some_id",
							DiscussionID:   "some_id",
							CreatedTime:    timestamp,
							LastEditedTime: timestamp,
							CreatedBy:      user,
							Parent: notion.Parent{
								Type:   notion.ParentTypePageID,
								PageID: "some_id",
							},
							RichText: []notion.RichText{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "content"},
								},
							},
						},
					},
					HasMore:    false,
					NextCursor: "",
				},
				wantErr: false,
				err:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, err := client.Comment.Get(context.Background(), tt.id, nil)

				if (err != nil) != tt.wantErr {
					t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Query() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Create", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			request    *notion.CommentCreateRequest
			want       *notion.Comment
			wantErr    bool
			err        error
		}{
			{
				name:       "returns created comment",
				filePath:   "testdata/comment_create.json",
				statusCode: http.StatusOK,
				request: &notion.CommentCreateRequest{
					Parent: notion.Parent{
						Type:   notion.ParentTypePageID,
						PageID: "some_id",
					},
					RichText: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "Hello world"},
						},
					},
				},
				want: &notion.Comment{
					Object:         notion.ObjectTypeComment,
					ID:             "some_id",
					DiscussionID:   "some_id",
					CreatedTime:    timestamp,
					LastEditedTime: timestamp,
					CreatedBy:      user,
					Parent: notion.Parent{
						Type:   notion.ParentTypePageID,
						PageID: "some_id",
					},
					RichText: []notion.RichText{
						{
							Type: notion.RichTextTypeText,
							Text: &notion.Text{Content: "Hello world"},
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := client.Comment.Create(context.Background(), tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Create() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
