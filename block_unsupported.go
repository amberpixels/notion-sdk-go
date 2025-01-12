package notion

// UnsupportedBlock is a Notion block for unsupported blocks
type UnsupportedBlock struct {
	BasicBlock
}

// NewUnsupportedBlock creates a new UnsupportedBlock
func NewUnsupportedBlock() *UnsupportedBlock {
	return &UnsupportedBlock{NewBasicBlock(BlockTypeUnsupported)}
}

var (
	_ Block             = (*UnsupportedBlock)(nil)
	_ HierarchicalBlock = (*UnsupportedBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeUnsupported, func() Block { return &UnsupportedBlock{} })
}
