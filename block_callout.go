package notion

// Reference: https://developers.notion.com/reference/block#callout

// Callout is a block that represents a callout (icon + color + rich text).
type Callout struct {
	AtomChildren

	RichText RichTexts `json:"rich_text"`
	Icon     *Icon     `json:"icon,omitempty"`
	Color    Color     `json:"color,omitempty"`
}

// CalloutBlock is a Notion block for a callout.
type CalloutBlock struct {
	BasicBlock
	Callout Callout `json:"callout"`
}

// NewCalloutBlock creates a new CalloutBlock.
func NewCalloutBlock(c Callout) *CalloutBlock {
	return &CalloutBlock{
		BasicBlock: NewBasicBlock(BlockTypeCallout, c.ChildCount() > 0),
		Callout:    c,
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

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *CalloutBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*CalloutBlock)(nil)
	_ HierarchicalBlock = (*CalloutBlock)(nil)
	_ BasicBlockHolder  = (*CalloutBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeCallout, func() Block { return &CalloutBlock{} })
}
