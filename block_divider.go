package notion

type Divider struct {
	// empty
}

type DividerBlock struct {
	BaseBlock
	Divider Divider `json:"divider"`
}

func NewDividerBlock() *DividerBlock {
	return &DividerBlock{
		BaseBlock: NewBaseBlock(BlockTypeDivider),
		Divider:   Divider{},
	}
}

var (
	_ Block             = (*DividerBlock)(nil)
	_ HierarchicalBlock = (*DividerBlock)(nil)
)
