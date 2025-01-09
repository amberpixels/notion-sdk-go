package notion

type ChildDatabase struct {
	Title string `json:"title"`
}

type ChildDatabaseBlock struct {
	BaseBlock
	ChildDatabase ChildDatabase `json:"child_database"`
}

func NewChildDatabaseBlock(title string) *ChildDatabaseBlock {
	cdb := &ChildDatabaseBlock{
		BaseBlock: NewBaseBlock(BlockTypeChildDatabase),
	}
	cdb.ChildDatabase.Title = title
	return cdb
}

var (
	_ Block             = (*ChildDatabaseBlock)(nil)
	_ HierarchicalBlock = (*ChildDatabaseBlock)(nil)
)
