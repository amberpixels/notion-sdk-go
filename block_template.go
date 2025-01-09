package notion

type Template struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
}

// Deprecated
type TemplateBlock struct {
	BaseBlock
	Template Template `json:"template"`
}

// Deprecated
func NewTemplateBlock(t Template) *TemplateBlock {
	return &TemplateBlock{
		BaseBlock: NewBaseBlock(BlockTypeTemplate, t.ChildCount() > 0),
		Template:  t,
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
