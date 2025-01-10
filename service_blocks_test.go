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

func TestBlocksService(t *testing.T) {
	ctx := context.Background()

	timestamp, err := time.Parse(time.RFC3339, "2021-05-24T05:06:34.827Z")
	if err != nil {
		t.Fatal(err)
	}

	user := notion.NewPersonUser("some_id", "some_user@example.com")

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
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.GetChildren(ctx, tt.id, nil)

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
				want:       decorateTestBasicBlock(notion.NewChildPageBlock("Hello"), "some_id", &timestamp, user),
				wantErr:    false,
				err:        nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.Get(ctx, tt.id)

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
				wantErr: false,
				err:     nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))
				got, err := client.Blocks.Update(ctx, tt.id, tt.req)

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
