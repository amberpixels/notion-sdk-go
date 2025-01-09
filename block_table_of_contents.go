package notion

// Ref: https://developers.notion.com/reference/block#table-of-contents

type TableOfContents struct {
	// empty
	Color string `json:"color,omitempty"`
}

type TableOfContentsBlock struct {
	BaseBlock
	TableOfContents TableOfContents `json:"table_of_contents"`
}

func NewTableOfContentsBlock(toc TableOfContents) *TableOfContentsBlock {
	return &TableOfContentsBlock{
		BaseBlock:       NewBaseBlock(BlockTypeTableOfContents),
		TableOfContents: toc,
	}
}

var (
	_ Block             = (*TableOfContentsBlock)(nil)
	_ HierarchicalBlock = (*TableOfContentsBlock)(nil)
)
