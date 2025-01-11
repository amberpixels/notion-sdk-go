package notion

import "time"

// Date is a type for Notion dates.
type Date time.Time

// String returns the string representation of the Date.
func (d *Date) String() string {
	return time.Time(*d).Format(time.RFC3339)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.RFC3339, string(data))

	// Because the API does not distinguish between datetime with a
	// timezone and dates, we eventually have to try both.
	if err != nil {
		if _, ok := err.(*time.ParseError); !ok {
			return err
		}

		t, err = time.Parse("2006-01-02", string(data)) // Date
		if err != nil {
			// Still cannot parse it, nothing else to try.
			return err
		}
	}

	*d = Date(t)
	return nil
}
