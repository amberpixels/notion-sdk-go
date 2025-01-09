package notion

type Paragraph struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

type ParagraphBlock struct {
	BaseBlock
	Paragraph Paragraph `json:"paragraph"`
}

func NewParagraphBlock(p Paragraph) *ParagraphBlock {
	return &ParagraphBlock{
		BaseBlock: NewBaseBlock(BlockTypeParagraph, p.ChildCount() > 0),
		Paragraph: p,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *ParagraphBlock) SetChildren(children Blocks) {
	b.Paragraph.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *ParagraphBlock) AppendChildren(children ...Block) {
	b.Paragraph.AppendChildren(children...)
	b.HasChildren = b.Paragraph.ChildCount() > 0
}

var (
	_ Block             = (*ParagraphBlock)(nil)
	_ HierarchicalBlock = (*ParagraphBlock)(nil)
)
