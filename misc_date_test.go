package notion_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestDate(t *testing.T) {
	t.Run(".UnmarshalText", func(t *testing.T) {
		var d notion.Date

		t.Run("OK datetime with timezone", func(t *testing.T) {
			data := []byte("1987-02-13T00:00:00.000+01:00")
			err := d.UnmarshalText(data)
			require.NoError(t, err, "expected no error for valid datetime with timezone")
		})

		t.Run("OK date", func(t *testing.T) {
			data := []byte("1985-01-02")
			err := d.UnmarshalText(data)
			require.NoError(t, err, "expected no error for valid date")
		})

		t.Run("NOK", func(t *testing.T) {
			data := []byte("1985")
			err := d.UnmarshalText(data)
			assert.Error(t, err, "expected an error for invalid date format")
		})
	})
}
