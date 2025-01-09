package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	pathDatabases = "databases"
)

// DatabasesService is a service for Notion Databases API.
type DatabasesService struct {
	api *clientAPI
}

// NewDatabasesService creates an instance of DatabasesService.
func NewDatabasesService(api *clientAPI) *DatabasesService {
	return &DatabasesService{api: api}
}

// Create creates a database as a subpage in the specified parent page, with the
// specified properties schema. Currently, the parent of a new database must be
// a Notion page.
//
// See https://developers.notion.com/reference/create-a-database
func (s *DatabasesService) Create(ctx context.Context, requestBody *DatabaseCreateRequest) (*Database, error) {
	res, err := s.api.request(ctx, http.MethodPost, pathDatabases, nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response Database
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Query gets a list of Pages contained in the database, filtered and ordered
// according to the filter conditions and sort criteria provided in the request.
// The response may contain fewer than page_size of results. If the response
// includes a next_cursor value, refer to the pagination reference for details
// about how to use a cursor to iterate through the list.
//
// Filters are similar to the filters provided in the Notion UI where the set of
// filters and filter groups chained by "And" in the UI is equivalent to having
// each filter in the array of the compound "and" filter. Similar a set of
// filters chained by "Or" in the UI would be represented as filters in the
// array of the "or" compound filter.
//
// Filters operate on database properties and can be combined. If no filter is
// provided, all the pages in the database will be returned with pagination.
//
// See https://developers.notion.com/reference/post-database-query
func (s *DatabasesService) Query(ctx context.Context, id DatabaseID, requestBody *DatabaseQueryRequest) (*DatabaseQueryResponse, error) {
	res, err := s.api.request(ctx, http.MethodPost, fmt.Sprintf(pathDatabases+"/%s/query", id.String()), nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response DatabaseQueryResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Get gets a database by ID.
//
// See https://developers.notion.com/reference/get-database
func (s *DatabasesService) Get(ctx context.Context, id DatabaseID) (*Database, error) {
	if id == "" {
		return nil, errors.New("empty database id")
	}

	res, err := s.api.request(ctx, http.MethodGet, fmt.Sprintf(pathDatabases+"/%s", id.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response Database

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Update updates a Database by id
//
// https://developers.notion.com/reference/update-a-database
func (s *DatabasesService) Update(ctx context.Context, id DatabaseID, requestBody *DatabaseUpdateRequest) (*Database, error) {
	res, err := s.api.request(ctx, http.MethodPatch, fmt.Sprintf(pathDatabases+"/%s", id.String()), nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response Database
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DatabaseCreateRequest represents the request body for DatabaseClient.Create.
type DatabaseCreateRequest struct {
	// A page parent.
	Parent Parent `json:"parent"`
	// Title of database as it appears in Notion. An array of rich text objects.
	Title RichTexts `json:"title"`
	// Property schema of database. The keys are the names of properties as they
	// appear in Notion and the values are property schema objects.
	Properties PropertyConfigs `json:"properties"`
	IsInline   bool            `json:"is_inline"`
}

// DatabaseQueryRequest represents the request body for DatabaseClient.Query.
type DatabaseQueryRequest struct {
	// When supplied, limits which pages are returned based on the filter
	// conditions.
	Filter Filter
	// When supplied, orders the results based on the provided sort criteria.
	Sorts []SortObject `json:"sorts,omitempty"`
	// When supplied, returns a page of results starting after the cursor provided.
	// If not supplied, this endpoint will return the first page of results.
	StartCursor Cursor `json:"start_cursor,omitempty"`
	// The number of items from the full list desired in the response. Maximum: 100
	PageSize int `json:"page_size,omitempty"`
}

// DatabaseUpdateRequest represents the request body for DatabaseClient.Update.
type DatabaseUpdateRequest struct {
	// An array of rich text objects that represents the title of the database
	// that is displayed in the Notion UI. If omitted, then the database title
	// remains unchanged.
	Title RichTexts `json:"title,omitempty"`
	// The properties of a database to be changed in the request, in the form of
	// a JSON object. If updating an existing property, then the keys are the
	// names or IDs of the properties as they appear in Notion, and the values are
	// property schema objects. If adding a new property, then the key is the name
	// of the new database property and the value is a property schema object.
	Properties PropertyConfigs `json:"properties,omitempty"`
}

// DatabaseQueryResponse is a response for DatabaseQueryRequest.
type DatabaseQueryResponse struct {
	AtomPaginatedResponse
	Results Pages `json:"results"`
}

// MarshalJSON implements custom Query marshalling
func (qr *DatabaseQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sorts       []SortObject `json:"sorts,omitempty"`
		StartCursor Cursor       `json:"start_cursor,omitempty"`
		PageSize    int          `json:"page_size,omitempty"`
		Filter      any          `json:"filter,omitempty"`
	}{
		Sorts:       qr.Sorts,
		StartCursor: qr.StartCursor,
		PageSize:    qr.PageSize,
		Filter:      qr.Filter,
	})
}
