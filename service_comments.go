package notion

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

const (
	pathComments = "comments"
)

// CommentsService is a service for the Comments Notion API.
type CommentsService struct {
	api *clientAPI
}

// newCommentsService creates an instance CommentsService.
func newCommentsService(api *clientAPI) *CommentsService {
	return &CommentsService{api: api}
}

// Create creates a comment in a page or existing discussion thread.
//
// There are two locations you can add a new comment to:
// 1. A page
// 2. An existing discussion thread
//
// If the intention is to add a new comment to a page, a parent object must be
// provided in the body params. Alternatively, if a new comment is being added
// to an existing discussion thread, the discussion_id string must be provided
// in the body params. Exactly one of these parameters must be provided.
//
// See https://developers.notion.com/reference/create-a-comment
func (s *CommentsService) Create(ctx context.Context, requestBody *CommentCreateRequest) (*Comment, error) {
	res, err := s.api.request(ctx, http.MethodPost, pathComments, nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response Comment
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Get retrieves a list of un-resolved Comment objects from a page or block.
//
// See https://developers.notion.com/reference/retrieve-a-comment
func (s *CommentsService) Get(ctx context.Context, id ObjectID, pagination *Pagination) (*CommentQueryResponse, error) {
	queryParams := map[string]string{}
	if pagination != nil {
		queryParams = pagination.ToQuery()
	}

	queryParams["block_id"] = id.String()

	res, err := s.api.request(ctx, http.MethodGet, pathComments, queryParams, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response CommentQueryResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CommentCreateRequest represents the request body for CommentClient.Create.
type CommentCreateRequest struct {
	Parent       Parent       `json:"parent,omitempty"`
	DiscussionID DiscussionID `json:"discussion_id,omitempty"`
	RichText     RichTexts    `json:"rich_text"`
}

// CommentQueryResponse is a type for comment query response.
type CommentQueryResponse struct {
	AtomPaginatedResponse
	Results Comments `json:"results"`
}
