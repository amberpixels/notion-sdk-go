package notion

type UnsupportedBlock struct {
	BaseBlock
}

func NewUnsupportedBlock() *UnsupportedBlock {
	return &UnsupportedBlock{NewBaseBlock(BlockTypeUnsupported)}
}

var (
	_ Block             = (*UnsupportedBlock)(nil)
	_ HierarchicalBlock = (*UnsupportedBlock)(nil)
)
