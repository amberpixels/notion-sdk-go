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

var (
	_ Block             = (*DividerBlock)(nil)
	_ HierarchicalBlock = (*DividerBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeDivider, func() Block { return &DividerBlock{} })
}
