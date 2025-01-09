package notion

type Toggle struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

type ToggleBlock struct {
	BaseBlock
	Toggle Toggle `json:"toggle"`
}

func NewToggleBlock(t Toggle) *ToggleBlock {
	return &ToggleBlock{
		BaseBlock: NewBaseBlock(BlockTypeToggle, t.ChildCount() > 0),
		Toggle:    t,
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
