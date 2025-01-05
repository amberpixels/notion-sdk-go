package notionapi

import (
	"fmt"
	"time"
)

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type ObjectID string

func (oID ObjectID) String() string {
	return string(oID)
}

type Object interface {
	GetObject() ObjectType
}

// DataObject wraps the logic of Notion Block, Page and Database
type DataObject interface {
	Object

	GetID() ObjectID
	GetParent() Parent

	GetCreatedTime() *time.Time
	GetCreatedBy() *User

	GetLastEditedTime() *time.Time
	GetLastEditedBy() *User

	GetArchived() bool
	GetInTrash() bool
}

// DataObjectAtom is the base struct for Notion Block, Page and Database
type DataObjectAtom struct {
	Object ObjectType `json:"object"`

	ID     ObjectID `json:"id,omitempty"`
	Parent Parent   `json:"parent,omitempty"`

	CreatedTime *time.Time `json:"created_time"`
	CreatedBy   *User      `json:"created_by"`

	LastEditedTime *time.Time `json:"last_edited_time"`
	LastEditedBy   *User      `json:"last_edited_by"`

	Archived bool `json:"archived"`
	InTrash  bool `json:"in_trash"`
}

func (d DataObjectAtom) GetID() ObjectID { return d.ID }

func (d DataObjectAtom) GetCreatedTime() *time.Time { return d.CreatedTime }

func (d DataObjectAtom) GetLastEditedTime() *time.Time { return d.LastEditedTime }

func (d DataObjectAtom) GetCreatedBy() *User { return d.CreatedBy }

func (d DataObjectAtom) GetLastEditedBy() *User { return d.LastEditedBy }

func (d DataObjectAtom) GetArchived() bool { return d.Archived }

func (d DataObjectAtom) GetInTrash() bool { return d.InTrash }

func (d DataObjectAtom) GetParent() Parent { return d.Parent }

type ContentMedia interface {
	GetIcon() *Icon
	GetCover() *Image
}

type ContentMediaAtom struct {
	Icon  *Icon  `json:"icon,omitempty"`
	Cover *Image `json:"cover,omitempty"`
}

func (c ContentMediaAtom) GetIcon() *Icon   { return c.Icon }
func (c ContentMediaAtom) GetCover() *Image { return c.Cover }

type Color string

func (c Color) String() string {
	return string(c)
}

func (c Color) MarshalText() ([]byte, error) {
	if c == "" {
		return []byte(ColorDefault), nil
	}

	return []byte(c), nil
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

type RelationObject struct {
	Database           DatabaseID `json:"database"`
	SyncedPropertyName string     `json:"synced_property_name"`
}

type FunctionType string

func (ft FunctionType) String() string {
	return string(ft)
}

type Cursor string

func (c Cursor) String() string {
	return string(c)
}

type Date time.Time

func (d *Date) String() string {
	return time.Time(*d).Format(time.RFC3339)
}

func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.RFC3339, string(data))

	// Because the API does not distinguish between datetime with a
	// timezone and dates, we eventually have to try both.
	if err != nil {
		if _, ok := err.(*time.ParseError); !ok {
			return err
		} else {
			t, err = time.Parse("2006-01-02", string(data)) // Date
			if err != nil {
				// Still cannot parse it, nothing else to try.
				return err
			}
		}
	}

	*d = Date(t)
	return nil
}

type FileType string

type File struct {
	Name     string      `json:"name"`
	Type     FileType    `json:"type"`
	File     *FileObject `json:"file,omitempty"`
	External *FileObject `json:"external,omitempty"`
}

type FileObject struct {
	URL        string     `json:"url,omitempty"`
	ExpiryTime *time.Time `json:"expiry_time,omitempty"`
}

type Icon struct {
	Type     FileType    `json:"type"`
	Emoji    *Emoji      `json:"emoji,omitempty"`
	File     *FileObject `json:"file,omitempty"`
	External *FileObject `json:"external,omitempty"`
}

// GetURL returns the external or internal URL depending on the image type.
func (i Icon) GetURL() string {
	if i.File != nil {
		return i.File.URL
	}
	if i.External != nil {
		return i.External.URL
	}
	return ""
}

type Emoji string

type PropertyID string

func (pID PropertyID) String() string {
	return string(pID)
}

type Status = Option

type UniqueID struct {
	Prefix *string `json:"prefix,omitempty"`
	Number int     `json:"number"`
}

func (uID UniqueID) String() string {
	if uID.Prefix != nil {
		return fmt.Sprintf("%s-%d", *uID.Prefix, uID.Number)
	}
	return fmt.Sprintf("%d", uID.Number)
}

type VerificationState string

func (vs VerificationState) String() string {
	return string(vs)
}

// Verification documented here: https://developers.notion.com/reference/page-property-values#verification
type Verification struct {
	State      VerificationState `json:"state"`
	VerifiedBy *User             `json:"verified_by,omitempty"`
	Date       *DateObject       `json:"date,omitempty"`
}

type Button struct {
}
