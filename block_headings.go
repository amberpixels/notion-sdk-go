package notion

type Heading struct {
	AtomChildren

	RichText     RichTexts `json:"rich_text"`
	Color        string    `json:"color,omitempty"`
	IsToggleable bool      `json:"is_toggleable,omitempty"`
}

type Heading1Block struct {
	BaseBlock
	Heading1 Heading `json:"heading_1"`
}
type Heading2Block struct {
	BaseBlock
	Heading2 Heading `json:"heading_2"`
}
type Heading3Block struct {
	BaseBlock
	Heading3 Heading `json:"heading_3"`
}

var _ Block = (*Heading1Block)(nil)
var _ Block = (*Heading2Block)(nil)
var _ Block = (*Heading3Block)(nil)

func NewHeading1Block(h Heading) *Heading1Block {
	return &Heading1Block{
		BaseBlock: NewBaseBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading1:  h,
	}
}

func NewHeading2Block(h Heading) *Heading2Block {
	return &Heading2Block{
		BaseBlock: NewBaseBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading2:  h,
	}
}

func NewHeading3Block(h Heading) *Heading3Block {
	return &Heading3Block{
		BaseBlock: NewBaseBlock(BlockTypeColumn, h.ChildCount() > 0),
		Heading3:  h,
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

var (
	_ Block             = (*Heading1Block)(nil)
	_ HierarchicalBlock = (*Heading1Block)(nil)

	_ Block             = (*Heading2Block)(nil)
	_ HierarchicalBlock = (*Heading2Block)(nil)

	_ Block             = (*Heading3Block)(nil)
	_ HierarchicalBlock = (*Heading3Block)(nil)
)
