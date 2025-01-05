package notionast_test

import (
	"testing"

	"github.com/amberpixels/notion-sdk-go"
	"github.com/amberpixels/notion-sdk-go/x/notionast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromBlocks(t *testing.T) {
	// Create blocks
	childBlock1 := notion.NewParagraphBlock(notion.Paragraph{})
	childBlock1.ID = notion.BlockID("child1-id")

	childBlock2 := notion.NewParagraphBlock(notion.Paragraph{})
	childBlock2.ID = notion.BlockID("child2-id")

	paragraphBlock := notion.NewParagraphBlock(notion.Paragraph{
		Children: notion.Blocks{childBlock1, childBlock2},
	})
	paragraphBlock.ID = notion.BlockID("paragraph-id")

	blocks := notion.Blocks{paragraphBlock}

	// Build the tree
	node := notionast.BlocksToAST(blocks, nil)

	notionast.PrintAST(node)

	/* Expected output:
	- [tmp-0000000000] paragraph
	  - [paragraph-id] paragraph
	    - [child1-id] paragraph
	    - [child2-id] paragraph
	*/

	// Verify the tree structure
	require.Equal(t, 1, node.GetChildCount(), "Root node should have 1 child")

	paragraphNode := node.GetFirstChild()
	require.NotNil(t, paragraphNode, "First child of root node should not be nil")
	assert.Equal(t, "paragraph-id", paragraphNode.GetID().String(), "Paragraph node ID mismatch")
	assert.Equal(t, 2, paragraphNode.GetChildCount(), "Paragraph node should have 2 children")

	firstChild := paragraphNode.GetFirstChild()
	lastChild := paragraphNode.GetLastChild()

	require.NotNil(t, firstChild, "First child of paragraph node should not be nil")
	require.NotNil(t, lastChild, "Last child of paragraph node should not be nil")

	assert.Equal(t, "child1-id", firstChild.GetID().String(), "First child ID mismatch")
	assert.Equal(t, "child2-id", lastChild.GetID().String(), "Last child ID mismatch")

	assert.Nil(t, firstChild.GetPrevSibling(), "First child's previous sibling should be nil")
	assert.Nil(t, lastChild.GetNextSibling(), "Last child's next sibling should be nil")

	assert.NotNil(t, firstChild.GetParent(), "First child's parent should not be nil")
	assert.Equal(t, paragraphNode, firstChild.GetParent(), "First child's parent mismatch")

	assert.NotNil(t, lastChild.GetParent(), "Last child's parent should not be nil")
	assert.Equal(t, paragraphNode, lastChild.GetParent(), "Last child's parent mismatch")
}
