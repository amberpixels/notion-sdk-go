package notion

// Reference: https://developers.notion.com/reference/block#paragraph

// Paragraph is a type for paragraph blocks
type Paragraph struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    Color     `json:"color,omitempty"`
}

// ParagraphBlock is a Notion block for paragraph blocks
type ParagraphBlock struct {
	BasicBlock
	Paragraph Paragraph `json:"paragraph"`
}

// NewParagraphBlock creates a new ParagraphBlock
func NewParagraphBlock(p Paragraph) *ParagraphBlock {
	return &ParagraphBlock{
		BasicBlock: NewBasicBlock(BlockTypeParagraph, p.ChildCount() > 0),
		Paragraph:  p,
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

func init() {
	registerBlockDecoder(BlockTypeParagraph, func() Block { return &ParagraphBlock{} })
}
