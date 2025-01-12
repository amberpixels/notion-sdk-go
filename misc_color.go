package notion

// Color is a type for Notion colors.
type Color string

// String returns the string representation of the Color.
func (c Color) String() string { return string(c) }

// nolint:revive
const (
	ColorDefault           Color = "default"
	ColorGray              Color = "gray"
	ColorBrown             Color = "brown"
	ColorOrange            Color = "orange"
	ColorYellow            Color = "yellow"
	ColorGreen             Color = "green"
	ColorBlue              Color = "blue"
	ColorPurple            Color = "purple"
	ColorPink              Color = "pink"
	ColorRed               Color = "red"
	ColorDefaultBackground Color = "default_background"
	ColorGrayBackground    Color = "gray_background"
	ColorBrownBackground   Color = "brown_background"
	ColorOrangeBackground  Color = "orange_background"
	ColorYellowBackground  Color = "yellow_background"
	ColorGreenBackground   Color = "green_background"
	ColorBlueBackground    Color = "blue_background"
	ColorPurpleBackground  Color = "purple_background"
	ColorPinkBackground    Color = "pink_background"
	ColorRedBackground     Color = "red_background"
)

// MarshalText implements the encoding.TextMarshaler interface.
func (c Color) MarshalText() ([]byte, error) {
	if c == "" {
		return []byte(ColorDefault), nil
	}

	return []byte(c), nil
}
