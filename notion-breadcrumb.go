package notion

type Breadcrumb struct {
	// empty
}

type BreadcrumbBlock struct {
	BaseBlock
	Breadcrumb Breadcrumb `json:"breadcrumb"`
}

func NewBreadcrumbBlock() *BreadcrumbBlock {
	return &BreadcrumbBlock{
		BaseBlock:  NewBaseBlock(BlockTypeBreadcrumb),
		Breadcrumb: Breadcrumb{},
	}
}

var _ Block = (*BreadcrumbBlock)(nil)
