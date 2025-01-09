package notion

// Ref: https://developers.notion.com/reference/block#table

type Table struct {
	AtomChildren

	TableWidth      int  `json:"table_width"`
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
}

type TableBlock struct {
	BaseBlock
	Table Table `json:"table"`
}

type TableRow struct {
	Cells []RichTexts `json:"cells"`
}

type TableRowBlock struct {
	BaseBlock
	TableRow TableRow `json:"table_row"`
}

func NewTableRowBlock(tr TableRow) *TableRowBlock {
	return &TableRowBlock{
		BaseBlock: NewBaseBlock(BlockTypeTableRow),
		TableRow:  tr,
	}
}

func NewTableBlock(table Table) *TableBlock {
	return &TableBlock{
		BaseBlock: NewBaseBlock(BlockTypeTable, table.ChildCount() > 0),
		Table:     table,
	}
}

// SetChildren calls inner .SetChildren + updates the HasChildren field
func (b *TableBlock) SetChildren(children Blocks) {
	b.Table.SetChildren(children)
	b.HasChildren = len(children) > 0
}

// AppendChildren calls inner .AppendChildren + updates the HasChildren field
func (b *TableBlock) AppendChildren(children ...Block) {
	b.Table.AppendChildren(children...)
	b.HasChildren = b.Table.ChildCount() > 0
}

var (
	_ Block             = (*TableBlock)(nil)
	_ HierarchicalBlock = (*TableBlock)(nil)

	_ Block             = (*TableRowBlock)(nil)
	_ HierarchicalBlock = (*TableRowBlock)(nil)
)
