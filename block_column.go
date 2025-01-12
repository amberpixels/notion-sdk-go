package notion

// Reference: https://developers.notion.com/reference/block#column-list-and-column

// Column stores the children of the column block
type Column struct {
	AtomChildren
}

// ColumnBlock is a Notion block for Column
type ColumnBlock struct {
	BasicBlock
	Column Column `json:"column"`
}

// NewColumnBlock returns a new ColumnBlock with the given column
func NewColumnBlock(col Column) *ColumnBlock {
	return &ColumnBlock{
		BasicBlock: NewBasicBlock(BlockTypeColumn, col.ChildCount() > 0),
		Column:     col,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *ColumnBlock) SetChildren(children Blocks) {
	b.Column.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *ColumnBlock) AppendChildren(children ...Block) {
	b.Column.AppendChildren(children...)
	b.HasChildren = b.Column.ChildCount() > 0
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *ColumnBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*ColumnBlock)(nil)
	_ HierarchicalBlock = (*ColumnBlock)(nil)
	_ BasicBlockHolder  = (*ColumnBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeColumn, func() Block { return &ColumnBlock{} })
}
