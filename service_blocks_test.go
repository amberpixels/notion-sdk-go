package notion_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	notion "github.com/amberpixels/notion-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestBlocksService(t *testing.T) {
	ctx := context.Background()

	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	assert.NoError(t, err, "Failed to parse timestamp")

	user := notion.NewPersonUser("some_id", "some_user@example.com")

	t.Run("GetChildren", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.BlockID
			len        int
			wantErr    bool
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
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.GetChildren(ctx, tt.id, nil)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				assert.Len(t, got.Results, tt.len, "Unexpected number of results")
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
		}{
			{
				name:       "return list object",
				id:         "some_id",
				filePath:   "testdata/block_append_children.json",
				statusCode: http.StatusOK,
				request: &notion.AppendBlockChildrenRequest{
					Children: notion.Blocks{
						decorateTestBasicBlock(notion.NewHeading2Block(
							notion.Heading{
								RichText: notion.RichTexts{
									notion.NewTextRichText("Hello"),
								},
							},
						), "block1", &timestamp, user),
					},
				},
				want: &notion.AppendBlockChildrenResponse{
					Object: notion.ObjectTypeList,
					Results: notion.Blocks{
						decorateTestBasicBlock(notion.NewHeading2Block(
							notion.Heading{
								RichText: notion.RichTexts{
									notion.NewTextRichText("Hello"),
								},
							},
						), "block1", &timestamp, user),
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.AppendChildren(ctx, tt.id, tt.request)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				assert.JSONEq(t, string(toJSON(t, tt.want)), string(toJSON(t, got)), "Unexpected response")
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
		}{
			{
				name:       "returns block object",
				filePath:   "testdata/block_get.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				want:       decorateTestBasicBlock(notion.NewChildPageBlock("Hello"), "some_id", &timestamp, user),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.Get(ctx, tt.id)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				assert.Equal(t, tt.want, got, "Unexpected block")
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
		}{
			{
				name:       "updates block and returns it",
				filePath:   "testdata/block_update.json",
				statusCode: http.StatusOK,
				id:         "some_id",
				req: &notion.BlockUpdateRequest{
					Paragraph: &notion.Paragraph{
						RichText: notion.RichTexts{
							notion.NewTextRichText("Hello"),
						},
						Color: notion.ColorYellow,
					},
				},
				want: decorateTestBasicBlock(
					notion.NewParagraphBlock(notion.Paragraph{
						RichText: notion.RichTexts{
							notion.NewTextRichText("Hello"),
						},
						Color: notion.ColorYellow,
					}), "some_id", &timestamp, user),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.Update(ctx, tt.id, tt.req)

				assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
				assert.Equal(t, tt.want, got, "Unexpected block update result")
			})
		}
	})
}

func TestBlockUpdateRequest_MarshalJSON(t *testing.T) {
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
			assert.Equal(t, tt.wantErr, err != nil, "Unexpected error state")
			assert.JSONEq(t, string(tt.want), string(got), "Unexpected JSON")
		})
	}
}
