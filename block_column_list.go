package notion

// Reference: https://developers.notion.com/reference/block#column-list-and-column

// ColumnList stores the children of the column list block
type ColumnList struct {
	// Children can only contain column blocks
	// Children should have at least 2 blocks when appending.
	// TODO^ validate this
	AtomChildren
}

// ColumnListBlock is a Notion block for ColumnList
type ColumnListBlock struct {
	BasicBlock
	ColumnList ColumnList `json:"column_list"`
}

// NewColumnListBlock returns a new ColumnListBlock with the given column list
func NewColumnListBlock(col ColumnList) *ColumnListBlock {
	return &ColumnListBlock{
		BasicBlock: NewBasicBlock(BlockTypeColumnList, col.ChildCount() > 0),
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

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *ColumnListBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*ColumnListBlock)(nil)
	_ HierarchicalBlock = (*ColumnListBlock)(nil)
	_ BasicBlockHolder  = (*ColumnListBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeColumnList, func() Block { return &ColumnListBlock{} })
}
