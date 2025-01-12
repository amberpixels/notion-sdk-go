package notion_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestAuthenticationService(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateToken", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			request    *notion.TokenCreateRequest
			want       *notion.TokenCreateResponse
			wantErr    error
		}{
			{
				name:       "Creates token",
				filePath:   "testdata/create_token.json",
				statusCode: http.StatusOK,
				request: &notion.TokenCreateRequest{
					Code:        "code1",
					GrantType:   "authorization_code",
					RedirectURI: "www.example.com",
				},
				want: &notion.TokenCreateResponse{
					AccessToken:          "token1",
					BotID:                "bot1",
					DuplicatedTemplateID: "template_id1",
					WorkspaceIcon:        "🎉",
					WorkspaceID:          "workspaceid_1",
					WorkspaceName:        "workspace_1",
				},
				wantErr: nil,
			},
			{
				name:       "Fails to create token",
				filePath:   "testdata/create_token_error.json",
				statusCode: http.StatusBadRequest,
				request: &notion.TokenCreateRequest{
					Code:        "code1",
					GrantType:   "authorization_code",
					RedirectURI: "www.example.com",
				},
				wantErr: &notion.TokenCreateError{
					Code:    "invalid_grant",
					Message: "Invalid code.",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Mock client
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.New("some_token", notion.WithTransport(c))

				// Call method under test
				got, gotErr := client.Auth.CreateToken(ctx, tt.request)

				// Assertions
				if tt.wantErr != nil {
					assert.Error(t, gotErr)
					assert.Equal(t, tt.wantErr, gotErr, "error mismatch")
				} else {
					require.NoError(t, gotErr)
				}
				assert.Equal(t, tt.want, got, "response mismatch")
			})
		}
	})
}
