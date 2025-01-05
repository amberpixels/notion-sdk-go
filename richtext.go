package notionapi

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type Link struct {
	Url string `json:"url,omitempty"`
}

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color,omitempty"`
}

type MentionType string

func (mType MentionType) String() string {
	return string(mType)
}

type DatabaseMention struct {
	ID ObjectID `json:"id"`
}

type PageMention struct {
	ID ObjectID `json:"id"`
}

type TemplateMentionType string

func (tMType TemplateMentionType) String() string {
	return string(tMType)
}

type TemplateMention struct {
	Type                TemplateMentionType `json:"type"`
	TemplateMentionUser string              `json:"template_mention_user,omitempty"`
	TemplateMentionDate string              `json:"template_mention_date,omitempty"`
}

type Mention struct {
	Type            MentionType      `json:"type,omitempty"`
	Database        *DatabaseMention `json:"database,omitempty"`
	Page            *PageMention     `json:"page,omitempty"`
	User            *User            `json:"user,omitempty"`
	Date            *DateObject      `json:"date,omitempty"`
	TemplateMention *TemplateMention `json:"template_mention,omitempty"`
}

type RichTextType string

func (rtType RichTextType) String() string {
	return string(rtType)
}

type RichText struct {
	Type        RichTextType `json:"type,omitempty"`
	Text        *Text        `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
}

// TODO: switch to Clone-based modifiers

// MakeLink makes the RichText a link to the given destination
func (rt *RichText) MakeLink(destination string) *RichText {
	if rt.Text != nil {
		rt.Text.Link = &Link{Url: destination}
		rt.Href = destination
	}

	return rt
}

// AnnotateBold annotates the RichText with bold text
func (rt *RichText) AnnotateBold() *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Bold = true
	return rt
}

// AnnotateItalic annotates the RichText with italic text
func (rt *RichText) AnnotateItalic() *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Italic = true
	return rt
}

// AnnotateStrikethrough annotates the RichText with strikethrough text
func (rt *RichText) AnnotateStrikethrough() *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Strikethrough = true
	return rt
}

// AnnotateUnderline annotates the RichText with underline text
func (rt *RichText) AnnotateUnderline() *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Underline = true
	return rt
}

// AnnotateCode annotates the RichText with code text
func (rt *RichText) AnnotateCode() *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Code = true
	return rt
}

// AnnotateColor annotates the RichText with a specific color
func (rt *RichText) AnnotateColor(color Color) *RichText {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Color = color
	return rt
}

// NewTextRichText creates a new RichText with the given text
// It fully builds the RichText object with all fields populated.
func NewTextRichText(text string) *RichText {
	return &RichText{
		Type: RichTextTypeText,
		Text: &Text{
			Content: text,
		},
		PlainText: text,
	}
}

// NewLinkRichText creates a new RichText with the given content and link
// It fully builds the RichText object with all fields populated.
func NewLinkRichText(content, link string) *RichText {
	return &RichText{
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

// TODO: NewMentionRichText, NewEquationRichText
/*
func NewDatabaseMentionRichText(databaseID ObjectID) *RichText {
	return &RichText{
		Type: RichTextTypeText,
		Mention: &Mention{
			Type: MentionTypeDatabase,
			Database: &DatabaseMention{
				ID: databaseID,
			},
		},
	}
}
*/
