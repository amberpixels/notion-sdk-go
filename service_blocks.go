package notion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	pathBlocks = "blocks"
)

// BlocksService is a service for the Blocks Notion API.
type BlocksService struct {
	api *clientAPI
}

// newBlocksService creates an instance BlocksService.
func newBlocksService(api *clientAPI) *BlocksService {
	return &BlocksService{api: api}
}

// Get retrieves a Block object using the ID specified.
//
// Get https://developers.notion.com/reference/retrieve-a-block
func (s *BlocksService) Get(ctx context.Context, id BlockID) (Block, error) {
	res, err := s.api.request(ctx, http.MethodGet, fmt.Sprintf(pathBlocks+"/%s", id.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return decodeBlock(response)
}

// GetChildren returns a paginated array of child block objects contained in the block using
// the ID specified. In order to receive a complete representation of a block,
// you may need to recursively retrieve the block children of child blocks.
//
// See https://developers.notion.com/reference/get-block-children
func (s *BlocksService) GetChildren(ctx context.Context, id BlockID, pagination *Pagination) (*GetChildrenResponse, error) {
	res, err := s.api.request(ctx, http.MethodGet, fmt.Sprintf(pathBlocks+"/%s/children", id.String()), pagination.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	response := &GetChildrenResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Appendchildren creates and appends new children blocks to the parent block_id specified.
// Blocks can be parented by other blocks, pages, or databases.
//
// Returns a paginated list of newly created first level children block objects.
//
// Existing blocks cannot be moved using this endpoint. Blocks are appended to
// the bottom of the parent block. Once a block is appended as a child, it can't
// be moved elsewhere via the API.
//
// For blocks that allow children, we allow up to two levels of nesting in a
// single request.
//
// See https://developers.notion.com/reference/patch-block-children
func (s *BlocksService) AppendChildren(ctx context.Context, id BlockID, requestBody *AppendBlockChildrenRequest) (*AppendBlockChildrenResponse, error) {
	res, err := s.api.request(ctx, http.MethodPatch, fmt.Sprintf(pathBlocks+"/%s/children", id.String()), nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response AppendBlockChildrenResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Update updates the content for the specified block_id based on the block type.
// Supported fields based on the block object type (see Block object for
// available fields and the expected input for each field).
//
// Note: The update replaces the entire value for a given field. If a field is
// omitted (ex: omitting checked when updating a to_do block), the value will not be changed.
//
// See https://developers.notion.com/reference/update-a-block
func (s *BlocksService) Update(ctx context.Context, id BlockID, requestBody *BlockUpdateRequest) (Block, error) {
	res, err := s.api.request(ctx, http.MethodPatch, fmt.Sprintf(pathBlocks+"/%s", id.String()), nil, requestBody)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return decodeBlock(response)
}

// Delete sets a Block object, including page blocks, to archived: true using the ID
// specified. Note: in the Notion UI application, this moves the block to the
// "Trash" where it can still be accessed and restored.
//
// To restore the block with the API, use the Update a block or Update page respectively.
//
// See https://developers.notion.com/reference/delete-a-block
func (s *BlocksService) Delete(ctx context.Context, id BlockID) (Block, error) {
	res, err := s.api.request(ctx, http.MethodDelete, fmt.Sprintf(pathBlocks+"/%s", id.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if errClose := res.Body.Close(); errClose != nil {
			log.Println("failed to close body, should never happen")
		}
	}()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return decodeBlock(response)
}

// AppendBlockChildrenRequest is a type for append block children request.
type AppendBlockChildrenRequest struct {
	// Append new children after a specific block. If empty, new children with be appended to the bottom of the parent block.
	After BlockID `json:"after,omitempty"`
	// Child content to append to a container block as an array of block objects.
	Children Blocks `json:"children"`
}

// GetChildrenResponse is a type for get children response.
type GetChildrenResponse struct {
	AtomPaginatedResponse
	Results Blocks `json:"results"`
}

// BlockUpdateRequest is a type for block update request.
type BlockUpdateRequest struct {
	Paragraph        *Paragraph `json:"paragraph,omitempty"`
	Heading1         *Heading   `json:"heading_1,omitempty"`
	Heading2         *Heading   `json:"heading_2,omitempty"`
	Heading3         *Heading   `json:"heading_3,omitempty"`
	BulletedListItem *ListItem  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItem  `json:"numbered_list_item,omitempty"`
	Code             *Code      `json:"code,omitempty"`
	ToDo             *ToDo      `json:"to_do,omitempty"`
	Toggle           *Toggle    `json:"toggle,omitempty"`
	Embed            *Embed     `json:"embed,omitempty"`
	Image            *File      `json:"image,omitempty"`
	Video            *File      `json:"video,omitempty"`
	File             *File      `json:"file,omitempty"`
	Pdf              *File      `json:"pdf,omitempty"`
	Bookmark         *Bookmark  `json:"bookmark,omitempty"`
	Template         *Template  `json:"template,omitempty"`
	Callout          *Callout   `json:"callout,omitempty"`
	Equation         *Equation  `json:"equation,omitempty"`
	Quote            *Quote     `json:"quote,omitempty"`
	TableRow         *TableRow  `json:"table_row,omitempty"`
}

type AppendBlockChildrenResponse struct {
	Object  ObjectType `json:"object"`
	Results []Block    `json:"results"`
}

// UnmarshalJSON does custom unmarshalling for AppendBlockChildrenResponse
func (r *AppendBlockChildrenResponse) UnmarshalJSON(data []byte) error {
	type appendBlockResponse struct {
		Object  ObjectType       `json:"object"`
		Results []map[string]any `json:"results"`
	}

	var raw appendBlockResponse
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	blocks := make([]Block, 0)
	for _, b := range raw.Results {
		block, err := decodeBlock(b)
		if err != nil {
			return err
		}
		blocks = append(blocks, block)
	}

	*r = AppendBlockChildrenResponse{
		Object:  raw.Object,
		Results: blocks,
	}
	return nil
}

// TODO: write a test that uses reflect and ensure that such structs
//       above always have the all types being included
