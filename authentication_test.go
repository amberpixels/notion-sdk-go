package notion_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestAuthenticationClient(t *testing.T) {
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
					RedirectUri: "www.example.com",
				},
				want: &notion.TokenCreateResponse{
					AccessToken:          "token1",
					BotId:                "bot1",
					DuplicatedTemplateId: "template_id1",
					WorkspaceIcon:        "🎉",
					WorkspaceId:          "workspaceid_1",
					WorkspaceName:        "workspace_1",
				},
				wantErr: nil,
			},
			{
				name:       "Creates token",
				filePath:   "testdata/create_token_error.json",
				statusCode: http.StatusBadRequest,
				request: &notion.TokenCreateRequest{
					Code:        "code1",
					GrantType:   "authorization_code",
					RedirectUri: "www.example.com",
				},
				wantErr: &notion.TokenCreateError{
					Code:    "invalid_grant",
					Message: "Invalid code.",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))
				got, gotErr := client.Authentication.CreateToken(context.Background(), tt.request)

				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					t.Errorf("Query() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Query() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
