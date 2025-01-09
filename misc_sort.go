package notion

type SortOrder string

const (
	SortOrderASC  SortOrder = "ascending"
	SortOrderDESC SortOrder = "descending"
)

type TimestampType string

const (
	TimestampCreated    TimestampType = "created_time"
	TimestampLastEdited TimestampType = "last_edited_time"
)

type SortObject struct {
	Property  string        `json:"property,omitempty"`
	Timestamp TimestampType `json:"timestamp,omitempty"`
	Direction SortOrder     `json:"direction,omitempty"`
}
