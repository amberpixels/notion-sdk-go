package ntast

import (
	"testing"

	nt "github.com/jomei/notionapi"
)

func TestFromBlocks(t *testing.T) {
	// Create blocks
	childBlock1 := nt.NewParagraphBlock(nt.Paragraph{})
	childBlock1.ID = "child1-id"

	childBlock2 := nt.NewParagraphBlock(nt.Paragraph{})
	childBlock2.ID = "child2-id"

	paragraphBlock := nt.NewParagraphBlock(nt.Paragraph{
		Children: nt.Blocks{childBlock1, childBlock2},
	})
	paragraphBlock.ID = "paragraph-id"

	blocks := nt.Blocks{paragraphBlock}

	// Build the tree
	node := FromBlocks(blocks, nil)

	PrintAST(node)

	/* Expected output:
	- [tmp-0000000000] paragraph
	  - [paragraph-id] paragraph
	    - [child1-id] paragraph
	    - [child2-id] paragraph
	*/

	// Verify the tree structure
	if node.GetChildCount() != 1 { // Root node has 1 child
		t.Fatalf("expected 1 child, got %d", node.GetChildCount())
	}

	paragraphNode := node.GetFirstChild()
	if paragraphNode.GetID() != "paragraph-id" {
		t.Errorf("expected paragraph ID to be 'paragraph-id', got %s", paragraphNode.GetID())
	}
	if paragraphNode.GetChildCount() != 2 { // Paragraph node has 2 children
		t.Fatalf("expected 2 children, got %d", paragraphNode.GetChildCount())
	}
}
