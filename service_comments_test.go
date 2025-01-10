package notion_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestCommentsService(t *testing.T) {
	ctx := context.Background()

	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	require.NoError(t, err)

	user := notion.NewPersonUser("some_id", "tests@example.com")

	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.ObjectID
			want       *notion.CommentQueryResponse
			wantErr    bool
		}{
			{
				name:       "returns comments for given block",
				filePath:   "testdata/comment_get.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				want: &notion.CommentQueryResponse{
					AtomPaginatedResponse: notion.AtomPaginatedResponse{
						Object:     notion.ObjectTypeList,
						HasMore:    false,
						NextCursor: notion.EmptyCursor,
					},
					Results: notion.Comments{
						{
							AtomObject: notion.AtomObject{
								Object: notion.ObjectTypeComment,
							},
							AtomID: notion.AtomID{
								ID: "some_id",
							},
							AtomParent: notion.AtomParent{
								Parent: notion.NewPageParent("some_id"),
							},
							AtomCreated: notion.AtomCreated{
								CreatedTime: &timestamp,
								CreatedBy:   user,
							},
							DiscussionID:   "some_id",
							LastEditedTime: &timestamp,
							RichText: notion.RichTexts{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "content"},
								},
							},
						},
						{
							AtomObject: notion.AtomObject{
								Object: notion.ObjectTypeComment,
							},
							AtomID: notion.AtomID{
								ID: "some_id",
							},
							AtomParent: notion.AtomParent{
								Parent: notion.NewPageParent("some_id"),
							},
							AtomCreated: notion.AtomCreated{
								CreatedTime: &timestamp,
								CreatedBy:   user,
							},
							DiscussionID:   "some_id",
							LastEditedTime: &timestamp,
							RichText: notion.RichTexts{
								{
									Type: notion.RichTextTypeText,
									Text: &notion.Text{Content: "content2"},
								},
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

				got, err := client.Comments.Get(ctx, tt.id, nil)

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, got)
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
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeComment,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomParent: notion.AtomParent{
						Parent: notion.NewPageParent("some_id"),
					},
					AtomCreated: notion.AtomCreated{
						CreatedTime: &timestamp,
						CreatedBy:   user,
					},
					DiscussionID:   "some_id",
					LastEditedTime: &timestamp,
					RichText: notion.RichTexts{
						notion.NewTextRichText("Hello world"),
					},
				},
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))

				got, err := client.Comments.Create(ctx, tt.request)

				if tt.wantErr {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})
}
