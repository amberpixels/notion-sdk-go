package notion_test

import (
	"context"
	"net/http"
	"testing"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestSearchClient(t *testing.T) {
	t.Run("Do", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			request    *notion.SearchRequest
			wantErr    bool
			err        error
		}{
			{
				name:       "returns search result",
				filePath:   "testdata/search.json",
				statusCode: http.StatusOK,
				request: &notion.SearchRequest{
					Query: "Hel",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := client.Search.Do(context.Background(), tt.request)

				if (err != nil) != tt.wantErr {
					t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if got == nil {
					t.Errorf("Search result is nil")
				}
			})
		}
	})
}
