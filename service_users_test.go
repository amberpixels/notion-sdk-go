package notion_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestUsersService(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.UserID
			want       *notion.User
			wantErr    bool
		}{
			{
				name:       "returns user by id",
				id:         "some_id",
				filePath:   "testdata/user_get.json",
				statusCode: http.StatusOK,
				want: &notion.User{
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeUser,
					},
					Type:      notion.UserTypePerson,
					Name:      "John Doe",
					AvatarURL: "some.url",
					Person:    &notion.Person{Email: "some@email.com"},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := notion.NewUsersService(client).Get(context.Background(), tt.id)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	t.Run("List", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			want       *notion.UsersListResponse
			wantErr    bool
		}{
			{
				name:       "returns list of users",
				filePath:   "testdata/user_list.json",
				statusCode: http.StatusOK,
				want: &notion.UsersListResponse{
					Object: notion.ObjectTypeList,
					Results: notion.Users{
						{
							AtomObject: notion.AtomObject{
								Object: notion.ObjectTypeUser,
							},
							AtomID: notion.AtomID{
								ID: "some_id",
							},
							Type:      notion.UserTypePerson,
							Name:      "John Doe",
							AvatarURL: "some.url",
							Person:    &notion.Person{Email: "some@email.com"},
						},
						{
							AtomObject: notion.AtomObject{
								Object: notion.ObjectTypeUser,
							},
							AtomID: notion.AtomID{
								ID: "some_id",
							},
							Type: notion.UserTypeBot,
							Name: "Test",
							Bot:  &notion.Bot{},
						},
					},
					HasMore: false,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := notion.NewUsersService(client).List(context.Background(), nil)
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	t.Run("Me", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			want       *notion.User
			wantErr    bool
		}{
			{
				name:       "returns me-user",
				filePath:   "testdata/user_me.json",
				statusCode: http.StatusOK,
				want: &notion.User{
					AtomObject: notion.AtomObject{
						Object: notion.ObjectTypeUser,
					},
					AtomID: notion.AtomID{
						ID: "some_id",
					},
					Type:      notion.UserTypePerson,
					Name:      "John Doe",
					AvatarURL: "some.url",
					Bot: &notion.Bot{Owner: notion.Owner{
						Type:      "workspace",
						Workspace: true,
					}, WorkspaceName: "John Doe's Workspace"},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := newMockedClient(t, tt.filePath, tt.statusCode)
				client := notion.NewClient("some_token", notion.WithHTTPClient(c))

				got, err := notion.NewUsersService(client).Me(context.Background())
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
