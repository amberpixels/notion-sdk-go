package notion_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/amberpixels/notion-sdk-go"
	"github.com/stretchr/testify/assert"
)

func decorateTestBasicBlock(block notion.Block, id notion.BlockID, timestamp *time.Time, user *notion.User) notion.Block {
	return block.(notion.BasicBlockHolder).SetBasicBlock(
		newTestBasicBlock(
			id, block.GetType(), timestamp, user,
		),
	)
}

// Helper function to marshal to JSON and handle errors
func toJSON(t *testing.T, v any) []byte {
	data, err := json.Marshal(v)
	assert.NoError(t, err, "Failed to marshal to JSON")
	return data
}

func newTestBasicBlock(id notion.BlockID, blockType notion.BlockType, timestamp *time.Time, u *notion.User) notion.BasicBlock {
	return notion.BasicBlock{
		AtomObject: notion.AtomObject{
			Object: notion.ObjectTypeBlock,
		},
		AtomID: notion.AtomID{
			ID: id,
		},
		AtomCreated: notion.AtomCreated{
			CreatedTime: timestamp,
			CreatedBy:   u,
		},
		AtomLastEdited: notion.AtomLastEdited{
			LastEditedTime: timestamp,
			LastEditedBy:   u,
		},
		Type: blockType,
	}
}
