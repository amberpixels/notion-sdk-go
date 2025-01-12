package notion

// Reference: https://developers.notion.com/reference/block#headings

// Heading stores the heading data.
type Heading struct {
	AtomChildren

	RichText     RichTexts `json:"rich_text"`
	Color        Color     `json:"color,omitempty"`
	IsToggleable bool      `json:"is_toggleable,omitempty"`
}

// Heading1Block is a Notion block for Heading1
type Heading1Block struct {
	BasicBlock
	Heading1 Heading `json:"heading_1"`
}

// Heading2Block is a Notion block for Heading2
type Heading2Block struct {
	BasicBlock
	Heading2 Heading `json:"heading_2"`
}

// Heading3Block is a Notion block for Heading3
type Heading3Block struct {
	BasicBlock
	Heading3 Heading `json:"heading_3"`
}

var (
	_ Block = (*Heading1Block)(nil)
	_ Block = (*Heading2Block)(nil)
	_ Block = (*Heading3Block)(nil)
)

// NewHeading1Block returns a new Heading1Block with the given heading
func NewHeading1Block(h Heading) *Heading1Block {
	return &Heading1Block{
		BasicBlock: NewBasicBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading1:   h,
	}
}

// NewHeading2Block returns a new Heading2Block with the given heading
func NewHeading2Block(h Heading) *Heading2Block {
	return &Heading2Block{
		BasicBlock: NewBasicBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading2:   h,
	}
}

// NewHeading3Block returns a new Heading3Block with the given heading
func NewHeading3Block(h Heading) *Heading3Block {
	return &Heading3Block{
		BasicBlock: NewBasicBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading3:   h,
	}
}

// NewHeadingBlock returns a new Heading[1-3]Block (hidden below Block interface)
// corresponding to the given heading level.
// It defaults to Heading 3 if the given level is not supported.
func NewHeadingBlock(heading Heading, level int) Block {
	switch level {
	case 1:
		return NewHeading1Block(heading)
	case 2:
		return NewHeading2Block(heading)
	case 3:
		return NewHeading3Block(heading)
	default:
		// fallback to level 3
		return NewHeading3Block(heading)
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *Heading1Block) SetChildren(children Blocks) {
	b.Heading1.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *Heading1Block) AppendChildren(children ...Block) {
	b.Heading1.AppendChildren(children...)
	b.HasChildren = b.Heading1.ChildCount() > 0
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *Heading2Block) SetChildren(children Blocks) {
	b.Heading2.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *Heading2Block) AppendChildren(children ...Block) {
	b.Heading2.AppendChildren(children...)
	b.HasChildren = b.Heading2.ChildCount() > 0
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *Heading3Block) SetChildren(children Blocks) {
	b.Heading3.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *Heading3Block) AppendChildren(children ...Block) {
	b.Heading3.AppendChildren(children...)
	b.HasChildren = b.Heading3.ChildCount() > 0
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *Heading1Block) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *Heading2Block) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *Heading3Block) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*Heading1Block)(nil)
	_ HierarchicalBlock = (*Heading1Block)(nil)
	_ BasicBlockHolder  = (*Heading1Block)(nil)

	_ Block             = (*Heading2Block)(nil)
	_ HierarchicalBlock = (*Heading2Block)(nil)
	_ BasicBlockHolder  = (*Heading2Block)(nil)

	_ Block             = (*Heading3Block)(nil)
	_ HierarchicalBlock = (*Heading3Block)(nil)
	_ BasicBlockHolder  = (*Heading3Block)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeHeading1, func() Block { return &Heading1Block{} })
	registerBlockDecoder(BlockTypeHeading2, func() Block { return &Heading2Block{} })
	registerBlockDecoder(BlockTypeHeading3, func() Block { return &Heading3Block{} })
}
