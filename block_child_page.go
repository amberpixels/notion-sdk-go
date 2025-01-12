package notion

// Reference: https://developers.notion.com/reference/block#child-page

// ChildPage stores the title of the child page
type ChildPage struct {
	Title string `json:"title"`
}

// ChildPageBlock is a Notion block for ChildPage
type ChildPageBlock struct {
	BasicBlock
	ChildPage ChildPage `json:"child_database"`
}

// NewChildPageBlock returns a new ChildPageBlock with the given title
func NewChildPageBlock(title string) *ChildPageBlock {
	cdb := &ChildPageBlock{
		BasicBlock: NewBasicBlock(BlockTypeChildPage),
	}
	cdb.ChildPage.Title = title
	return cdb
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *ChildPageBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*ChildPageBlock)(nil)
	_ HierarchicalBlock = (*ChildPageBlock)(nil)
	_ BasicBlockHolder  = (*ChildPageBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeChildPage, func() Block { return &ChildPageBlock{} })
}
