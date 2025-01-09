package notion

// ParentType is a type of a Notion parent.
// See https://developers.notion.com/reference/parent-object
type ParentType string

const (
	ParentTypeDatabaseID ParentType = "database_id"
	ParentTypePageID     ParentType = "page_id"
	ParentTypeBlockID    ParentType = "block_id"
	ParentTypeWorkspace  ParentType = "workspace"
)

// Pages, databases, and blocks are either located inside other pages,
// databases, and blocks, or are located at the top level of a workspace. This
// location is known as the "parent". Parent information is represented by a
// consistent parent object throughout the API.
type Parent struct {
	Type       ParentType `json:"type,omitempty"`
	PageID     PageID     `json:"page_id,omitempty"`
	DatabaseID DatabaseID `json:"database_id,omitempty"`
	BlockID    BlockID    `json:"block_id,omitempty"`
	Workspace  bool       `json:"workspace,omitempty"`
}

// IsZero returns true if the Parent is empty.
// We intentionally do not use pointerish *Parent to keep it nil-safe.
func (p Parent) IsZero() bool {
	return p.Type == "" && p.PageID == "" && p.DatabaseID == "" && p.BlockID == "" && !p.Workspace
}

// NewPageParent returns a Page Parent.
func NewPageParent(pageID PageID) Parent {
	return Parent{
		Type:   ParentTypePageID,
		PageID: pageID,
	}
}

// NewDatabaseParent returns a Database parent.
func NewDatabaseParent(databaseID DatabaseID) Parent {
	return Parent{
		Type:       ParentTypeDatabaseID,
		DatabaseID: databaseID,
	}
}

// NewBlockParent returns a Block parent.
func NewBlockParent(blockID BlockID) Parent {
	return Parent{
		Type:    ParentTypeBlockID,
		BlockID: blockID,
	}
}

// NewWorkspaceParent returns a Workspace parent.
func NewWorkspaceParent() Parent {
	return Parent{
		Type:      ParentTypeWorkspace,
		Workspace: true,
	}
}
