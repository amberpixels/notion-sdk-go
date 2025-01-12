package notion

// Reference: http://developers.notion.com/reference/block#table

// Table is a type for table blocks
type Table struct {
	AtomChildren

	TableWidth      int  `json:"table_width"`
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
}

// TableBlock is a Notion block for table blocks
type TableBlock struct {
	BasicBlock
	Table Table `json:"table"`
}

// TableRow is a type for table row blocks
type TableRow struct {
	Cells []RichTexts `json:"cells"`
}

// TableRowBlock is a Notion block for table row blocks
type TableRowBlock struct {
	BasicBlock
	TableRow TableRow `json:"table_row"`
}

// NewTableRowBlock creates a new TableRowBlock
func NewTableRowBlock(tr TableRow) *TableRowBlock {
	return &TableRowBlock{
		BasicBlock: NewBasicBlock(BlockTypeTableRow),
		TableRow:   tr,
	}
}

// NewTableBlock creates a new TableBlock
func NewTableBlock(table Table) *TableBlock {
	return &TableBlock{
		BasicBlock: NewBasicBlock(BlockTypeTable, table.ChildCount() > 0),
		Table:      table,
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

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *TableBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *TableRowBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*TableBlock)(nil)
	_ HierarchicalBlock = (*TableBlock)(nil)
	_ BasicBlockHolder  = (*TableBlock)(nil)

	_ Block             = (*TableRowBlock)(nil)
	_ HierarchicalBlock = (*TableRowBlock)(nil)
	_ BasicBlockHolder  = (*TableRowBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeTable, func() Block { return &TableBlock{} })
	registerBlockDecoder(BlockTypeTableRow, func() Block { return &TableRowBlock{} })
}
