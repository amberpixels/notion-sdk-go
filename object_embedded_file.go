package notion

import "time"

// FileType is a type of a Notion file.
// See https://developers.notion.com/reference/file-object
type FileType string

// nolint:revive
const (
	FileTypeFile     FileType = "file"
	FileTypeExternal FileType = "external"
)

// File is a file object.
type File struct {
	Caption RichTexts `json:"caption,omitempty"`

	Type FileType `json:"type"`

	File     *FileData `json:"file,omitempty"`
	External *FileData `json:"external,omitempty"`
}

// Files is a slice of File objects.
type Files []File

// FileData is a file Data object
type FileData struct {
	URL        string     `json:"url,omitempty"`
	ExpiryTime *time.Time `json:"expiry_time,omitempty"`
}

// GetURL returns the URL of the File.
func (f File) GetURL() string {
	if f.File != nil {
		return f.File.URL
	}
	if f.External != nil {
		return f.External.URL
	}
	return ""
}

// GetExpiryTime returns the ExpiryTime of the File.
func (f File) GetExpiryTime() *time.Time {
	if f.File != nil {
		return f.File.ExpiryTime
	}
	return nil
}

// Icon is an type union of FileObject(type==external) and Emoji
// Icon must be filled in 2 possible ways:
// 1. Type==external, External is not nil (Emoji is nil)
// 2. Type==emoji, Emoji is not nil (External is nil)
type Icon struct {
	Type     FileType `json:"type"` // external or emoji
	External *File    `json:"external,omitempty"`

	Emoji Emoji `json:"emoji,omitempty"`
}
