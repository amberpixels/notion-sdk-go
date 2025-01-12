package notion

// Reference: https://developers.notion.com/reference/rich-text#mention
// 	          https://developers.notion.com/reference/block#mention

// MentionType is a type of a mention
type MentionType string

// String returns the string representation of the MentionType
func (v MentionType) String() string { return string(v) }

// nolint:revive
const (
	MentionTypeDatabase        MentionType = "database"
	MentionTypePage            MentionType = "page"
	MentionTypeUser            MentionType = "user"
	MentionTypeDate            MentionType = "date"
	MentionTypeLinkPreview     MentionType = "link_preview"
	MentionTypeTemplateMention MentionType = "template_mention"
)

// Mention is an Object that holds mention to something (database, page, etc)
type Mention struct {
	Type            MentionType      `json:"type,omitempty"`
	Database        *DatabaseMention `json:"database,omitempty"`
	Page            *PageMention     `json:"page,omitempty"`
	User            *UserMention     `json:"user,omitempty"`
	Date            *DateObject      `json:"date,omitempty"`
	TemplateMention *TemplateMention `json:"template_mention,omitempty"`
}

// DatabaseMention is a database mention object
type DatabaseMention struct {
	ID ObjectID `json:"id"`
}

// PageMention is a page mention object
type PageMention struct {
	ID ObjectID `json:"id"`
}

// UserMention is a user mention object
type UserMention struct {
	Object ObjectType `json:"object"` // always "user"
	ID     ObjectID   `json:"id"`
}

// TemplateMentionType is a type of a template mention
type TemplateMentionType string

// String returns the string representation of the TemplateMentionType
func (t TemplateMentionType) String() string { return string(t) }

// nolint:revive
const (
	TemplateMentionTypeUser TemplateMentionType = "template_mention_user"
	TemplateMentionTypeDate TemplateMentionType = "template_mention_date"
)

// TemplateMention is a template Mention object.
type TemplateMention struct {
	Type                TemplateMentionType `json:"type"`
	TemplateMentionUser string              `json:"template_mention_user,omitempty"`
	TemplateMentionDate string              `json:"template_mention_date,omitempty"`
}

// NewDatabaseMentionRichText creates a new RichText with mention to the given database ID
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

// NewPageMentionRichText creates a new RichText with mention to the given page ID
func NewPageMentionRichText(pageID ObjectID) *RichText {
	return &RichText{
		Type: RichTextTypeText,
		Mention: &Mention{
			Type: MentionTypePage,
			Page: &PageMention{ID: pageID},
		},
	}
}

// NewUserMentionRichText creates a new RichText with mention to the given user ID
func NewUserMentionRichText(userID ObjectID) *RichText {
	return &RichText{
		Type: RichTextTypeText,
		Mention: &Mention{
			Type: MentionTypeUser,
			User: &UserMention{ID: userID},
		},
	}
}
