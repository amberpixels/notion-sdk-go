package notion

// SortOrder is a type for sort order.
type SortOrder string

// nolint:revive
const (
	SortOrderASC  SortOrder = "ascending"
	SortOrderDESC SortOrder = "descending"
)

// TimestampType is a type for timestamp type.
type TimestampType string

// nolint:revive
const (
	TimestampCreated    TimestampType = "created_time"
	TimestampLastEdited TimestampType = "last_edited_time"
)

// SortObject is a type for sort object.
type SortObject struct {
	Property  string        `json:"property,omitempty"`
	Timestamp TimestampType `json:"timestamp,omitempty"`
	Direction SortOrder     `json:"direction,omitempty"`
}
