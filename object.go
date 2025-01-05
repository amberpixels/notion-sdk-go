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
