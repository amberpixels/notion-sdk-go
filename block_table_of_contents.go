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

var (
	_ Block             = (*TableOfContentsBlock)(nil)
	_ HierarchicalBlock = (*TableOfContentsBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeTableOfContents, func() Block { return &TableOfContentsBlock{} })
}
