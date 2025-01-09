package notion

type Column struct {
	AtomChildren
}

type ColumnBlock struct {
	BaseBlock
	Column Column `json:"column"`
}

func NewColumnBlock(col Column) *ColumnBlock {
	return &ColumnBlock{
		BaseBlock: NewBaseBlock(BlockTypeColumn, col.ChildCount() > 0),
		Column:    col,
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

var (
	_ Block             = (*ColumnBlock)(nil)
	_ HierarchicalBlock = (*ColumnBlock)(nil)
)
