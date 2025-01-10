package notion

// Reference: https://developers.notion.com/reference/block#bulleted-list-item
// Reference: https://developers.notion.com/reference/block#numbered-list-item

// ListItem is a type for bulleted and numbered list items
type ListItem struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

// BulletedListItemBlock is a Notion block for bulleted list items
type BulletedListItemBlock struct {
	BasicBlock
	BulletedListItem ListItem `json:"bulleted_list_item"`
}

// NewBulletedListItemBlock creates a new BulletedListItemBlock
func NewBulletedListItemBlock(li ListItem) *BulletedListItemBlock {
	return &BulletedListItemBlock{
		BasicBlock:       NewBasicBlock(BlockTypeBulletedListItem, li.ChildCount() > 0),
		BulletedListItem: li,
	}
}

// NumberedListItemBlock is a Notion block for numbered list items
type NumberedListItemBlock struct {
	BasicBlock
	NumberedListItem ListItem `json:"numbered_list_item"`
}

// NewNumberedListItemBlock creates a new NumberedListItemBlock
func NewNumberedListItemBlock(li ListItem) *NumberedListItemBlock {
	return &NumberedListItemBlock{
		BasicBlock:       NewBasicBlock(BlockTypeNumberedListItem, li.ChildCount() > 0),
		NumberedListItem: li,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *BulletedListItemBlock) SetChildren(children Blocks) {
	b.BulletedListItem.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *BulletedListItemBlock) AppendChildren(children ...Block) {
	b.BulletedListItem.AppendChildren(children...)
	b.HasChildren = b.BulletedListItem.ChildCount() > 0
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *NumberedListItemBlock) SetChildren(children Blocks) {
	b.NumberedListItem.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *NumberedListItemBlock) AppendChildren(children ...Block) {
	b.NumberedListItem.AppendChildren(children...)
	b.HasChildren = b.NumberedListItem.ChildCount() > 0
}

var (
	_ Block             = (*BulletedListItemBlock)(nil)
	_ HierarchicalBlock = (*BulletedListItemBlock)(nil)

	_ Block             = (*NumberedListItemBlock)(nil)
	_ HierarchicalBlock = (*NumberedListItemBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeBulletedListItem, func() Block { return &BulletedListItemBlock{} })
	registerBlockDecoder(BlockTypeNumberedListItem, func() Block { return &NumberedListItemBlock{} })
}
