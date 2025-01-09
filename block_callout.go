package notion

type Callout struct {
	AtomChildren

	RichText RichTexts `json:"rich_text"`
	Icon     *Icon     `json:"icon,omitempty"`
	Color    string    `json:"color,omitempty"`
}

type CalloutBlock struct {
	BaseBlock
	Callout Callout `json:"callout"`
}

func NewCalloutBlock(callout Callout) *CalloutBlock {
	return &CalloutBlock{
		BaseBlock: NewBaseBlock(BlockTypeCallout, callout.ChildCount() > 0),
		Callout:   callout,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *CalloutBlock) SetChildren(children Blocks) {
	b.Callout.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *CalloutBlock) AppendChildren(children ...Block) {
	b.Callout.AppendChildren(children...)
	b.HasChildren = b.Callout.ChildCount() > 0
}

var (
	_ Block             = (*CalloutBlock)(nil)
	_ HierarchicalBlock = (*CalloutBlock)(nil)
)
