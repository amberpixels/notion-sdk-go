package notion

// Reference: https://developers.notion.com/reference/block#divider

// Divider stands for a divider
type Divider struct{}

// DividerBlock is a Notion block for Divider
type DividerBlock struct {
	BasicBlock
	Divider Divider `json:"divider"`
}

// NewDividerBlock returns a new DividerBlock
func NewDividerBlock() *DividerBlock {
	return &DividerBlock{
		BasicBlock: NewBasicBlock(BlockTypeDivider),
		Divider:    Divider{},
	}
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *DividerBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*DividerBlock)(nil)
	_ HierarchicalBlock = (*DividerBlock)(nil)
	_ BasicBlockHolder  = (*DividerBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeDivider, func() Block { return &DividerBlock{} })
}
