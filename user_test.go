package notion_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestUserClient(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			id         notion.UserID
			want       *notion.User
			wantErr    bool
			err        error
		}{
			{
				name:       "returns user by id",
				id:         "some_id",
				filePath:   "testdata/user_get.json",
				statusCode: http.StatusOK,
				want: &notion.User{
					Object:    notion.ObjectTypeUser,
					ID:        "some_id",
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

				got, err := client.User.Get(context.Background(), tt.id)
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

	t.Run("List", func(t *testing.T) {
		tests := []struct {
			name       string
			filePath   string
			statusCode int
			want       *notion.UsersListResponse
			wantErr    bool
			err        error
		}{
			{
				name:       "returns list of users",
				filePath:   "testdata/user_list.json",
				statusCode: http.StatusOK,
				want: &notion.UsersListResponse{
					Object: notion.ObjectTypeList,
					Results: []notion.User{
						{
							Object:    notion.ObjectTypeUser,
							ID:        "some_id",
							Type:      notion.UserTypePerson,
							Name:      "John Doe",
							AvatarURL: "some.url",
							Person:    &notion.Person{Email: "some@email.com"},
						},
						{
							Object: notion.ObjectTypeUser,
							ID:     "some_id",
							Type:   notion.UserTypeBot,
							Name:   "Test",
							Bot:    &notion.Bot{},
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
				got, err := client.User.List(context.Background(), nil)
				if (err != nil) != tt.wantErr {
					t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("List() got = %v, want %v", got, tt.want)
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
			err        error
		}{
			{
				name:       "returns me-user",
				filePath:   "testdata/user_me.json",
				statusCode: http.StatusOK,
				want: &notion.User{
					Object:    notion.ObjectTypeUser,
					ID:        "some_id",
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

				got, err := client.User.Me(context.Background())
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
