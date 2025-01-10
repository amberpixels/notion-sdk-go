package notion

// Reference: https://developers.notion.com/reference/block#child-database

// ChildDatabase stores the title of the child database
type ChildDatabase struct {
	Title string `json:"title"`
}

// ChildDataBasicBlock is a Notion block for ChildDatabase
type ChildDataBasicBlock struct {
	BasicBlock
	ChildDatabase ChildDatabase `json:"child_database"`
}

// NewChildDataBasicBlock returns a new ChildDataBasicBlock with the given title
func NewChildDataBasicBlock(title string) *ChildDataBasicBlock {
	cdb := &ChildDataBasicBlock{
		BasicBlock: NewBasicBlock(BlockTypeChildDatabase),
	}
	cdb.ChildDatabase.Title = title
	return cdb
}

var (
	_ Block             = (*ChildDataBasicBlock)(nil)
	_ HierarchicalBlock = (*ChildDataBasicBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeChildDatabase, func() Block { return &ChildDataBasicBlock{} })
}
