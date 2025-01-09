package notion_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	notion "github.com/amberpixels/notion-sdk-go"
)

func TestColor_MarshalText(t *testing.T) {
	type Foo struct {
		Test notion.Color `json:"test"`
	}

	t.Run("marshal to color if color is not empty", func(t *testing.T) {
		f := Foo{Test: notion.ColorGreen}
		r, err := json.Marshal(f)

		// Use require to fail fast if error occurs
		require.NoError(t, err, "expected no error when marshaling valid color")

		want := []byte(`{"test":"green"}`)

		// Use assert for value comparison
		assert.JSONEq(t, string(want), string(r), "unexpected marshaled result for non-empty color")
	})

	t.Run("marshal to default color if color is empty", func(t *testing.T) {
		f := Foo{}
		r, err := json.Marshal(f)

		// Use require to fail fast if error occurs
		require.NoError(t, err, "expected no error when marshaling empty color")

		want := []byte(`{"test":"default"}`)

		// Use assert for value comparison
		assert.JSONEq(t, string(want), string(r), "unexpected marshaled result for empty color")
	})
}
