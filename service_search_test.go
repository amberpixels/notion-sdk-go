package notion_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestSearchClient(t *testing.T) {
	ctx := context.Background()

	t.Run("Do", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			request    *notion.SearchRequest
			wantErr    bool
		}{
			{
				name:       "returns search result",
				filePath:   "testdata/search.json",
				statusCode: http.StatusOK,
				request: &notion.SearchRequest{
					Query: "Hel",
				},
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Setup mocked client
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))

				// Perform the search
				got, err := client.Search.Do(ctx, tt.request)

				// Assert the error
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.NotNil(t, got, "Search result should not be nil")
				}
			})
		}
	})
}
