package notion

// UnsupportedBlock is a Notion block for unsupported blocks
type UnsupportedBlock struct {
	BasicBlock
}

// NewUnsupportedBlock creates a new UnsupportedBlock
func NewUnsupportedBlock() *UnsupportedBlock {
	return &UnsupportedBlock{NewBasicBlock(BlockTypeUnsupported)}
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *UnsupportedBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*UnsupportedBlock)(nil)
	_ HierarchicalBlock = (*UnsupportedBlock)(nil)
	_ BasicBlockHolder  = (*UnsupportedBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeUnsupported, func() Block { return &UnsupportedBlock{} })
}
