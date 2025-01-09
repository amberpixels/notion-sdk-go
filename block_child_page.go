package notion

type ChildPage struct {
	Title string `json:"title"`
}

type ChildPageBlock struct {
	BaseBlock
	ChildPage ChildPage `json:"child_database"`
}

func NewChildPageBlock(title string) *ChildPageBlock {
	cdb := &ChildPageBlock{
		BaseBlock: NewBaseBlock(BlockTypeChildPage),
	}
	cdb.ChildPage.Title = title
	return cdb
}

var (
	_ Block             = (*ChildPageBlock)(nil)
	_ HierarchicalBlock = (*ChildPageBlock)(nil)
)
