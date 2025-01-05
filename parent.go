package notionapi

// Pages, databases, and blocks are either located inside other pages,
// databases, and blocks, or are located at the top level of a workspace. This
// location is known as the "parent". Parent information is represented by a
// consistent parent object throughout the API.
//
// See https://developers.notion.com/reference/parent-object
type Parent struct {
	Type       ParentType `json:"type,omitempty"`
	PageID     PageID     `json:"page_id,omitempty"`
	DatabaseID DatabaseID `json:"database_id,omitempty"`
	BlockID    BlockID    `json:"block_id,omitempty"`
	Workspace  bool       `json:"workspace,omitempty"`
}

type ParentType string

func (p Parent) IsZero() bool {
	return p.Type == "" && p.PageID == "" && p.DatabaseID == "" && p.BlockID == ""
}

func NewPageParent(pageID PageID) Parent {
	return Parent{
		Type:   ParentTypePageID,
		PageID: pageID,
	}
}
func NewDatabaseParent(databaseID DatabaseID) Parent {
	return Parent{
		Type:       ParentTypeDatabaseID,
		DatabaseID: databaseID,
	}
}
func NewBlockParent(blockID BlockID) Parent {
	return Parent{
		Type:    ParentTypeBlockID,
		BlockID: blockID,
	}
}
