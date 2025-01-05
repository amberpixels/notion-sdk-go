package notionast_test

import (
	"testing"

	"github.com/amberpixels/notion-sdk-go"
	"github.com/amberpixels/notion-sdk-go/x/notionast"
)

func TestFromBlocks(t *testing.T) {
	// Create blocks
	childBlock1 := notion.NewParagraphBlock(notion.Paragraph{})
	childBlock1.ID = "child1-id"

	childBlock2 := notion.NewParagraphBlock(notion.Paragraph{})
	childBlock2.ID = "child2-id"

	paragraphBlock := notion.NewParagraphBlock(notion.Paragraph{
		Children: notion.Blocks{childBlock1, childBlock2},
	})
	paragraphBlock.ID = "paragraph-id"

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

	if paragraphNode.GetFirstChild() == nil {
		t.Errorf("expected first child to be not nil")
	}
	if paragraphNode.GetLastChild() == nil {
		t.Errorf("expected last child to be not nil")
	}
	firstChild := paragraphNode.GetFirstChild()
	lastChild := paragraphNode.GetLastChild()
	if firstChild.GetID() != "child1-id" {
		t.Errorf("expected first child ID to be 'child1-id', got %s", firstChild.GetID())
	}
	if lastChild.GetID() != "child2-id" {
		t.Errorf("expected last child ID to be 'child2-id', got %s", lastChild.GetID())
	}
	if firstChild.GetPrevSibling() != nil {
		t.Errorf("expected first child prev sibling to be nil")
	}
	if lastChild.GetNextSibling() != nil {
		t.Errorf("expected last child next sibling to be nil")
	}

	if firstChild.GetParent() == nil {
		t.Errorf("expected first child parent to be nil")
	} else {
		if firstChild.GetParent() != paragraphNode {
			t.Errorf("expected first child parent to be %s, got %s", paragraphNode.GetID(), firstChild.GetParent().GetID())
		}
	}

	if lastChild.GetParent() == nil {
		t.Errorf("expected last child parent to be nil")
	} else {
		if lastChild.GetParent() != paragraphNode {
			t.Errorf("expected last child parent to be %s, got %s", paragraphNode.GetID(), lastChild.GetParent().GetID())
		}
	}
}
