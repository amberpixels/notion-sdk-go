package notion

// Reference: https://developers.notion.com/reference/rich-text
//            https://developers.notion.com/reference/rich-text#the-annotation-object

// RichTextType is a type of a RichText
type RichTextType string

// String returns the string representation of the RichTextType
func (t RichTextType) String() string { return string(t) }

// nolint:revive
const (
	RichTextTypeText     RichTextType = "text"
	RichTextTypeMention  RichTextType = "mention"
	RichTextTypeEquation RichTextType = "equation"
)

// RichText is a rich text object
type RichText struct {
	Type RichTextType `json:"type,omitempty"`

	Text     *Text     `json:"text,omitempty"`
	Mention  *Mention  `json:"mention,omitempty"`
	Equation *Equation `json:"equation,omitempty"`

	Annotations *Annotations `json:"annotations,omitempty"`

	// PlainText is the Text.Content or Mention.{*Name*} or Equation.Expression
	PlainText string `json:"plain_text,omitempty"`
	// Href is the Text.Link.Url or Mention.{*Href*}
	Href string `json:"href,omitempty"`
}

// RichTexts is a slice of RichText
type RichTexts []RichText

// Annotations is a set of annotations for RichText
type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}

// WithLink makes a new RichText with a link to the given destination
func (rt RichText) WithLink(destination string) RichText {
	if rt.Text != nil {
		rt.Text.Link = &Link{URL: destination}
		rt.Href = destination
	}

	return rt
}

// WithBold makes a new RichText annotated as bold text
func (rt RichText) WithBold() RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Bold = true
	return rt
}

// WithItalic makes a new RichText annotated as italic text
func (rt RichText) WithItalic() RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Italic = true
	return rt
}

// WithStrikethrough makes a new RichText annotated as strikethrough text
func (rt RichText) WithStrikethrough() RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Strikethrough = true
	return rt
}

// WithUnderline makes a new RichText annotated as underline text
func (rt RichText) WithUnderline() RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Underline = true
	return rt
}

// WithCode makes a new RichText annotated as code text
func (rt RichText) WithCode() RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Code = true
	return rt
}

// WithColor makes a new RichText annotated as colored text
func (rt RichText) WithColor(color Color) RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Color = color
	return rt
}
