package notion

//
// Temporary: (not documented. Docs page is not working. But it exists in contents
// Reference:https://developers.notion.com/reference/block#link-to-page
// Testing is required
//

// LinkToPageType is a type for LinkToPage.Type
type LinkToPageType string

const (
	// LinkToPageTypePage is a type for LinkToPage.Type
	LinkToPageTypePage LinkToPageType = "page"
	// LinkToPageTypeDatabase is a type for LinkToPage.Type
	LinkToPageTypeDatabase LinkToPageType = "database"
)

// LinkToPage holds a link to a page or database
type LinkToPage struct {
	Type       LinkToPageType `json:"type"`
	PageID     PageID         `json:"page_id,omitempty"`
	DatabaseID DatabaseID     `json:"database_id,omitempty"`
}

// LinkToPageBlock is a Notion block for LinkToPage
type LinkToPageBlock struct {
	BasicBlock
	LinkToPage LinkToPage `json:"link_to_page"`
}

// NewLinkToPageBlock creates a new LinkToPageBlock (type==page)
func NewLinkToPageBlock(pageID PageID) *LinkToPageBlock {
	return &LinkToPageBlock{
		BasicBlock: NewBasicBlock(BlockTypeLinkToPage),
		LinkToPage: LinkToPage{
			Type:   LinkToPageTypePage,
			PageID: pageID,
		},
	}
}

// NewLinkToDatabaseBlock creates a new LinkToPageBlock (type==database)
func NewLinkToDatabaseBlock(dbID DatabaseID) *LinkToPageBlock {
	return &LinkToPageBlock{
		BasicBlock: NewBasicBlock(BlockTypeLinkToPage),
		LinkToPage: LinkToPage{
			Type:       LinkToPageTypeDatabase,
			DatabaseID: dbID,
		},
	}
}

var (
	_ Block             = (*LinkToPageBlock)(nil)
	_ HierarchicalBlock = (*LinkToPageBlock)(nil)
)

func init() {
	registerBlockDecoder(BlockTypeLinkToPage, func() Block { return &LinkToPageBlock{} })
}
