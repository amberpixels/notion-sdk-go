package notion

// See https://developers.notion.com/reference/block#mention

// MentionType is a type of a mention
type MentionType string

// String returns the string representation of the MentionType
func (v MentionType) String() string { return string(v) }

const (
	MentionTypeDatabase        MentionType = "database"
	MentionTypePage            MentionType = "page"
	MentionTypeUser            MentionType = "user"
	MentionTypeDate            MentionType = "date"
	MentionTypeLinkPreview     MentionType = "link_preview"
	MentionTypeTemplateMention MentionType = "template_mention"
)

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

type TemplateMentionType string

func (tMType TemplateMentionType) String() string { return string(tMType) }

const (
	TemplateMentionTypeUser TemplateMentionType = "template_mention_user"
	TemplateMentionTypeDate TemplateMentionType = "template_mention_date"
)

type TemplateMention struct {
	Type                TemplateMentionType `json:"type"`
	TemplateMentionUser string              `json:"template_mention_user,omitempty"`
	TemplateMentionDate string              `json:"template_mention_date,omitempty"`
}
