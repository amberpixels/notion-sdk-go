package notion

// Reference: https://developers.notion.com/reference/rich-text#text

// Text is a text object for RichText.{Type==RichTextTypeText}
type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

// Link is a link object to be used in RichText.Text.Link
type Link struct {
	Url string `json:"url,omitempty"`
}

// NewTextRichText creates a new RichText with the given text
// It fully builds the RichText object with all fields populated.
func NewTextRichText(text string) RichText {
	return RichText{
		Type: RichTextTypeText,
		Text: &Text{
			Content: text,
		},
		PlainText: text,
	}.WithColor(ColorDefault)
}

// NewLinkRichText creates a new RichText with the given content and link
// It fully builds the RichText object with all fields populated.
func NewLinkRichText(content, link string) RichText {
	return RichText{
		Type: RichTextTypeText,
		Text: &Text{
			Content: content,
			Link: &Link{
				Url: link,
			},
		},
		PlainText: content,
		Href:      link,
	}
}
