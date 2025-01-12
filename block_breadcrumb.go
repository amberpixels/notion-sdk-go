package notion

// Reference: https://developers.notion.com/reference/block#breadcrumb

// Breadcrumb is a block that represents a breadcrumb. Contains nothing.
type Breadcrumb struct{}

// BreadcrumbBlock is a Notion block for a breadcrumb.
type BreadcrumbBlock struct {
	BasicBlock
	Breadcrumb Breadcrumb `json:"breadcrumb"`
}

// NewBreadcrumbBlock creates a new BreadcrumbBlock.
func NewBreadcrumbBlock() *BreadcrumbBlock {
	return &BreadcrumbBlock{
		BasicBlock: NewBasicBlock(BlockTypeBreadcrumb),
		Breadcrumb: Breadcrumb{},
	}
}

// SetBasicBlock implements the SetBasicBlock method of the BasicBlockHolder interface.
func (b *BreadcrumbBlock) SetBasicBlock(block BasicBlock) Block {
	b.BasicBlock = block
	return b
}

var (
	_ Block             = (*BreadcrumbBlock)(nil)
	_ HierarchicalBlock = (*BreadcrumbBlock)(nil)
	_ BasicBlockHolder  = (*BreadcrumbBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeBreadcrumb, func() Block { return &BreadcrumbBlock{} })
}
