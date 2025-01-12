package notion

// Reference: https://developers.notion.com/reference/block#link-preview-blocks

// LinkPreviewBlock is a Notion block for link preview blocks
// NOTE: will only be returned by the API. Cannot be created by the API.
type LinkPreviewBlock struct {
	BasicBlock
	LinkPreview LinkPreview `json:"link_preview"`
}

// LinkPreview is a type for link preview.
type LinkPreview struct {
	URL string `json:"url"`
}

// NewLinkPreviewBlock creates a new LinkPreviewBlock.
// Deprecated: as now publishing new LinkPreview blocks to Notion is allowed
func NewLinkPreviewBlock(lp LinkPreview) *LinkPreviewBlock {
	return &LinkPreviewBlock{
		BasicBlock:  NewBasicBlock(BlockTypeLinkPreview),
		LinkPreview: lp,
	}
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *LinkPreviewBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*LinkPreviewBlock)(nil)
	_ HierarchicalBlock = (*LinkPreviewBlock)(nil)
	_ BasicBlockHolder  = (*LinkPreviewBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeLinkPreview, func() Block { return &LinkPreviewBlock{} })
}
