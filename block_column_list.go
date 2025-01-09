package notion

type ColumnList struct {
	// Children can only contain column blocks
	// Children should have at least 2 blocks when appending.
	// TODO^ validate this
	AtomChildren
}

type ColumnListBlock struct {
	BaseBlock
	ColumnList ColumnList `json:"column_list"`
}

func NewColumnListBlock(col ColumnList) *ColumnListBlock {
	return &ColumnListBlock{
		BaseBlock:  NewBaseBlock(BlockTypeColumnList, col.ChildCount() > 0),
		ColumnList: col,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *ColumnListBlock) SetChildren(children Blocks) {
	b.ColumnList.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *ColumnListBlock) AppendChildren(children ...Block) {
	b.ColumnList.AppendChildren(children...)
	b.HasChildren = b.ColumnList.ChildCount() > 0
}

var (
	_ Block             = (*ColumnListBlock)(nil)
	_ HierarchicalBlock = (*ColumnListBlock)(nil)
)
