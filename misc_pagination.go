package notion

import "strconv"

// Cursor is the Notion's cursor value from the pagination.
type Cursor string

// String returns the string representation of the Cursor.
func (c Cursor) String() string { return string(c) }

// EmptyCursor is a constant for empty cursor.
const EmptyCursor = ""

// AtomPaginatedResponse is embedded in all paginated responses.
type AtomPaginatedResponse struct {
	Object     ObjectType `json:"object"`
	NextCursor Cursor     `json:"next_cursor"`
	HasMore    bool       `json:"has_more"`
}

// Pagination is a type for Notion pagination.
type Pagination struct {
	StartCursor Cursor
	PageSize    int
}

// ToQuery returns a map of query parameters for the pagination.
func (p *Pagination) ToQuery() map[string]string {
	if p == nil {
		return nil
	}
	r := map[string]string{}
	if p.StartCursor != "" {
		r["start_cursor"] = p.StartCursor.String()
	}

	if p.PageSize != 0 {
		r["page_size"] = strconv.Itoa(p.PageSize)
	}

	return r
}
