package notion

// Appearance is an interface for Blocks that can have an icon and/or a cover.
type Appearance interface {
	// GetIcon returns the Icon of the Appearance.
	GetIcon() *Icon

	// GetCover returns the Cover of the Appearance.
	GetCover() *File
}

// AtomAppearance is a container for the shared fields of the base Notion objects.
type AtomAppearance struct {
	Icon  *Icon `json:"icon,omitempty"`
	Cover *File `json:"cover,omitempty"`
}

// GetIcon returns the Icon of the AtomAppearance.
func (a AtomAppearance) GetIcon() *Icon { return a.Icon }

// GetCover returns the Cover of the AtomAppearance.
func (a AtomAppearance) GetCover() *File { return a.Cover }

var _ Appearance = (*AtomAppearance)(nil)

// NewExternalIcon returns a new Icon with a given external File URL.
func NewExternalIcon(url string) Icon {
	return Icon{
		Type: FileTypeExternal,
		External: &File{
			Type: FileTypeExternal,
			External: &FileData{
				URL: url,
			},
		},
	}
}

// NewEmojiIcon returns a new Icon with a given Emoji.
func NewEmojiIcon(emoji Emoji) *Icon {
	return &Icon{
		Type:  EmojiTypeEmoji,
		Emoji: emoji,
	}
}
