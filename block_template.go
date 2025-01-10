package notion

// Reference: https://developers.notion.com/reference/block#template
// Note:
// 	As of March 27, 2023 creation of template blocks will no longer be supported.

// Template is a type for template blocks
type Template struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
}

// Deprecated
// TemplateBlock is a Notion block for template blocks
type TemplateBlock struct {
	BasicBlock
	Template Template `json:"template"`
}

// Deprecated
// NewTemplateBlock creates a new TemplateBlock
func NewTemplateBlock(t Template) *TemplateBlock {
	return &TemplateBlock{
		BasicBlock: NewBasicBlock(BlockTypeTemplate, t.ChildCount() > 0),
		Template:   t,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *TemplateBlock) SetChildren(children Blocks) {
	b.Template.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *TemplateBlock) AppendChildren(children ...Block) {
	b.Template.AppendChildren(children...)
	b.HasChildren = b.Template.ChildCount() > 0
}

var (
	_ Block             = (*TemplateBlock)(nil)
	_ HierarchicalBlock = (*TemplateBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeTemplate, func() Block { return &TemplateBlock{} })
}
