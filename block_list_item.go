package notion

type ListItem struct {
	AtomChildren
	RichText RichTexts `json:"rich_text"`
	Color    string    `json:"color,omitempty"`
}

type BulletedListItemBlock struct {
	BaseBlock
	BulletedListItem ListItem `json:"bulleted_list_item"`
}

func NewBulletedListItemBlock(li ListItem) *BulletedListItemBlock {
	return &BulletedListItemBlock{
		BaseBlock:        NewBaseBlock(BlockTypeBulletedListItem, li.ChildCount() > 0),
		BulletedListItem: li,
	}
}

type NumberedListItemBlock struct {
	BaseBlock
	NumberedListItem ListItem `json:"numbered_list_item"`
}

func NewNumberedListItemBlock(li ListItem) *NumberedListItemBlock {
	return &NumberedListItemBlock{
		BaseBlock:        NewBaseBlock(BlockTypeNumberedListItem, li.ChildCount() > 0),
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
