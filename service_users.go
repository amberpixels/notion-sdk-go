package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	pathUsers = "users"
)

// UsersService is a service for Notion Users API.
type UsersService struct {
	api *clientAPI
}

// newUsersService creates an instance of UsersService.
func newUsersService(api *clientAPI) *UsersService {
	return &UsersService{api: api}
}

// List returns a paginated list of Users for the workspace.
// The response may containf fewer than page_size of results.
//
// See https://developers.notion.com/reference/get-users
func (s *UsersService) List(ctx context.Context, pagination *Pagination) (*UsersListResponse, error) {
	res, err := s.api.request(ctx, http.MethodGet, pathUsers, pagination.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response UsersListResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Get retrieves a User using the ID specified.
//
// See https://developers.notion.com/reference/get-user
func (s *UsersService) Get(ctx context.Context, id UserID) (*User, error) {
	res, err := s.api.request(ctx, http.MethodGet, fmt.Sprintf(pathUsers+"/%s", id.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response User
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Me retrieves the bot User associated with the API token provided in the
// authorization header. The bot will have an owner field with information about
// the person who authorized the integration.
//
// See https://developers.notion.com/reference/get-self
func (s *UsersService) Me(ctx context.Context) (*User, error) {
	res, err := s.api.request(ctx, http.MethodGet, pathUsers+"/me", nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response User
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UsersListResponse stands for a paginated list of Users.
type UsersListResponse struct {
	Object     ObjectType `json:"object"`
	Results    Users      `json:"results"`
	HasMore    bool       `json:"has_more"`
	NextCursor Cursor     `json:"next_cursor"`
}
