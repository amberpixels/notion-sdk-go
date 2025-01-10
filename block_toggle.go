package notion

// Reference: https://developers.notion.com/reference/block#toggle-blocks

// Toggle is a type for toggle blocks
type Toggle struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

// ToggleBlock is a Notion block for toggle blocks
type ToggleBlock struct {
	BasicBlock
	Toggle Toggle `json:"toggle"`
}

// NewToggleBlock creates a new ToggleBlock
func NewToggleBlock(t Toggle) *ToggleBlock {
	return &ToggleBlock{
		BasicBlock: NewBasicBlock(BlockTypeToggle, t.ChildCount() > 0),
		Toggle:     t,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *ToggleBlock) SetChildren(children Blocks) {
	b.Toggle.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *ToggleBlock) AppendChildren(children ...Block) {
	b.Toggle.AppendChildren(children...)
	b.HasChildren = b.Toggle.ChildCount() > 0
}

var (
	_ Block             = (*ToggleBlock)(nil)
	_ HierarchicalBlock = (*ToggleBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeToggle, func() Block { return &ToggleBlock{} })
}
