package notion

// NOTE: will only be returned by the API. Cannot be created by the API.
// https://developers.notion.com/reference/block#link-preview-blocks
type LinkPreviewBlock struct {
	BaseBlock
	LinkPreview LinkPreview `json:"link_preview"`
}

type LinkPreview struct {
	URL string `json:"url"`
}

// Deprecated: as now publishing new LinkPreview blocks to Notion is allowed
func NewLinkPreviewBlock(lp LinkPreview) *LinkPreviewBlock {
	return &LinkPreviewBlock{
		BaseBlock:   NewBaseBlock(BlockTypeLinkPreview),
		LinkPreview: lp,
	}
}

var (
	_ Block             = (*LinkPreviewBlock)(nil)
	_ HierarchicalBlock = (*LinkPreviewBlock)(nil)
)

//
// Temporary: (not documented)
// Testing is required
//

type LinkToPage struct {
	Type       BlockType  `json:"type"`
	PageID     PageID     `json:"page_id,omitempty"`
	DatabaseID DatabaseID `json:"database_id,omitempty"`
}

type LinkToPageBlock struct {
	BaseBlock
	LinkToPage LinkToPage `json:"link_to_page"`
}

func NewLinkToPageBlock(ltp LinkToPage) *LinkToPageBlock {
	return &LinkToPageBlock{
		BaseBlock:  NewBaseBlock(BlockTypeLinkToPage),
		LinkToPage: ltp,
	}
}

var (
	_ Block             = (*LinkToPageBlock)(nil)
	_ HierarchicalBlock = (*LinkToPageBlock)(nil)
)
