package notion

type Appearance interface {
	GetIcon() *Icon
	GetCover() *File
}

// AtomAppearance is a container for the shared fields of the base Notion objects.
type AtomAppearance struct {
	Icon  *Icon `json:"icon,omitempty"`
	Cover *File `json:"cover,omitempty"`
}

func (a AtomAppearance) GetIcon() *Icon  { return a.Icon }
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
