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

type RichText struct {
	Type        ObjectType   `json:"type,omitempty"`
	Text        *Text        `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
}

// MakeLink makes the RichText a link to the given destination
func (rt *RichText) MakeLink(destination string) {
	if rt.Text != nil {
		rt.Text.Link = &Link{Url: destination}
		rt.Href = destination
	}
}

// AnnotateBold annotates the RichText with bold text
func (rt *RichText) AnnotateBold() {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Bold = true
}

// AnnotateItalic annotates the RichText with italic text
func (rt *RichText) AnnotateItalic() {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Italic = true
}

// AnnotateStrikethrough annotates the RichText with strikethrough text
func (rt *RichText) AnnotateStrikethrough() {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Strikethrough = true
}

// AnnotateUnderline annotates the RichText with underline text
func (rt *RichText) AnnotateUnderline() {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Underline = true
}

// AnnotateCode annotates the RichText with code text
func (rt *RichText) AnnotateCode() {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Code = true
}

// AnnotateColor annotates the RichText with a specific color
func (rt *RichText) AnnotateColor(color Color) {
	if rt.Annotations == nil {
		rt.Annotations = &Annotations{}
	}

	rt.Annotations.Color = color
}

// NewTextRichText creates a new RichText with the given text
// It fully builds the RichText object with all fields populated.
func NewTextRichText(text string) *RichText {
	return &RichText{
		Type: ObjectTypeText,
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
		Type: ObjectTypeText,
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
		Type: ObjectTypeText,
		Mention: &Mention{
			Type: MentionTypeDatabase,
			Database: &DatabaseMention{
				ID: databaseID,
			},
		},
	}
}
*/
