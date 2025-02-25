package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	pathPages = "pages"
)

// PagesService is a service for the Pages Notion API.
type PagesService struct {
	api *clientAPI
}

// newPagesService creates an instance of PagesService.
func newPagesService(api *clientAPI) *PagesService {
	return &PagesService{api: api}
}

// Create creates a new page that is a child of an existing page or database.
//
// If the new page is a child of an existing page,title is the only valid
// property in the properties body param.
//
// If the new page is a child of an existing database, the keys of the
// properties object body param must match the parent database's properties.
//
// This endpoint can be used to create a new page with or without content using
// the children option. To add content to a page after creating it, use the
// Append block children endpoint.
//
// Returns a new page object.
//
// See https://developers.notion.com/reference/post-page
func (s *PagesService) Create(ctx context.Context, requestBody *PageCreateRequest) (*Page, error) {
	res, err := s.api.request(ctx, http.MethodPost, pathPages, nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	return handlePageResponse(res)
}

// Get retrieves a Page object using the ID specified.
//
// Responses contains page properties, not page content. To fetch page content,
// use the Retrieve block children endpoint.
//
// Page properties are limited to up to 25 references per page property. To
// retrieve data related to properties that have more than 25 references, use
// the Retrieve a page property endpoint.
//
// See https://developers.notion.com/reference/get-page
func (s *PagesService) Get(ctx context.Context, id PageID) (*Page, error) {
	res, err := s.api.request(ctx, http.MethodGet, fmt.Sprintf(pathPages+"/%s", id.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	return handlePageResponse(res)
}

// Update updates the properties of a page in a database. The properties body param of
// this endpoint can only be used to update the properties of a page that is a
// child of a database. The page’s properties schema must match the parent
// database’s properties.
//
// This endpoint can be used to update any page icon or cover, and can be used
// to archive or restore any page.
//
// To add page content instead of page properties, use the append block children
// endpoint. The page_id can be passed as the block_id when adding block
// children to the page.
//
// Returns the updated page object.
//
// See https://developers.notion.com/reference/patch-page
func (s *PagesService) Update(ctx context.Context, id PageID, request *PageUpdateRequest) (*Page, error) {
	res, err := s.api.request(ctx, http.MethodPatch, fmt.Sprintf(pathPages+"/%s", id.String()), nil, request)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	return handlePageResponse(res)
}

// PageCreateRequest represents the request body for PagesClient.Create.
type PageCreateRequest struct {
	// The parent page or database where the new page is inserted, represented as
	// a JSON object with a page_id or database_id key, and the corresponding ID.
	// Required field.
	Parent Parent `json:"parent"`
	// The values of the page’s properties. If the parent is a database, then the
	// schema must match the parent database’s properties. If the parent is a page,
	// then the only valid object key is title.
	// Required field.
	Properties Properties `json:"properties"`
	// The content to be rendered on the new page, represented as an array of
	// block objects.
	Children Blocks `json:"children,omitempty"`
	// The icon of the new page. Either an emoji object or an external file object.
	Icon *Icon `json:"icon,omitempty"`
	// The cover image of the new page, represented as a file object.
	Cover *File `json:"cover,omitempty"`
}

// PageUpdateRequest represents the request body for PagesClient.Update.
type PageUpdateRequest struct {
	// The property values to update for the page. The keys are the names or IDs
	// of the property and the values are property values. If a page property ID
	// is not included, then it is not changed.
	Properties Properties `json:"properties,omitempty"`
	// Whether the page is archived (deleted). Set to true to archive a page. Set
	// to false to un-archive (restore) a page.
	Archived bool `json:"archived"`
	// A page icon for the page. Supported types are external file object or emoji
	// object.
	Icon *Icon `json:"icon,omitempty"`
	// A cover image for the page. Only external file objects are supported.
	Cover *File `json:"cover,omitempty"`
}

func handlePageResponse(res *http.Response) (*Page, error) {
	var response Page
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
