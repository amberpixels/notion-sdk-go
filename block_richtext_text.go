package notion

// See https://developers.notion.com/reference/rich-text#text

// Text is a text object for RichText.{Type==RichTextTypeText}
type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

// Link is a link object to be used in RichText.Text.Link
type Link struct {
	Url string `json:"url,omitempty"`
}
