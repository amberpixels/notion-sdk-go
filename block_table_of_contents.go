package notion

// Ref: https://developers.notion.com/reference/block#table-of-contents

// TableOfContents is a type for table of contents blocks
type TableOfContents struct {
	// empty
	Color string `json:"color,omitempty"`
}

// TableOfContentsBlock is a Notion block for table of contents blocks
type TableOfContentsBlock struct {
	BasicBlock
	TableOfContents TableOfContents `json:"table_of_contents"`
}

// NewTableOfContentsBlock creates a new TableOfContentsBlock
func NewTableOfContentsBlock(toc TableOfContents) *TableOfContentsBlock {
	return &TableOfContentsBlock{
		BasicBlock:      NewBasicBlock(BlockTypeTableOfContents),
		TableOfContents: toc,
	}
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *TableOfContentsBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*TableOfContentsBlock)(nil)
	_ HierarchicalBlock = (*TableOfContentsBlock)(nil)
	_ BasicBlockHolder  = (*TableOfContentsBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeTableOfContents, func() Block { return &TableOfContentsBlock{} })
}
