// Package notionast provides a set of functions to work with Notion AST
// It is used to convert notion.Blocks to AST and vice versa
// It also provides Walk functionality to traverse the AST.
package notionast

import (
	"cmp"
	"crypto/rand"
	"fmt"
	"math/big"

	notion "github.com/amberpixels/notion-sdk-go"
)

// NodeID is a unique identifier for a node
type NodeID string

func (n NodeID) String() string { return string(n) }

// NodeType is a type of a node (It's considered to map to a block type)
type NodeType string

func (n NodeType) String() string { return string(n) }

// Node is a generic interface for a node in the AST
type Node interface {
	GetID() NodeID
	GetType() NodeType

	GetPrevSibling() Node
	GetNextSibling() Node

	GetChildCount() int
	GetFirstChild() Node
	GetLastChild() Node
	GetParent() Node

	AppendChild(child Node)
	RemoveChild(child Node)
	RemoveChildren() // removes all children
}

// Nodes is a slice of Node
type Nodes []Node

// NodeBlock is a node implementation for notion.Block
type NodeBlock struct {
	block  notion.Block
	nodeID NodeID // Same as the block ID (or generated if was empty)

	firstChild *NodeBlock
	lastChild  *NodeBlock

	parent *NodeBlock
	prev   *NodeBlock
	next   *NodeBlock
}

var _ Node = (*NodeBlock)(nil)

// RootNodeID is a constant for the root node auto-generated ID
const RootNodeID = "tmp-0000000000"

// NewNodeBlock returns a new NodeBlock with the given block
func NewNodeBlock(block notion.Block, parent *NodeBlock) *NodeBlock {
	if block == nil && parent == nil {
		// if no block and parent specified, it needs to create a fake root node
		// notion.Block has no concept for this (a root for block, etc)
		// so we just use ParagraphBlock for that
		return &NodeBlock{
			block:  notion.NewParagraphBlock(notion.Paragraph{}),
			nodeID: RootNodeID,
		}
	}

	return &NodeBlock{
		block: block,
		nodeID: cmp.Or(
			NodeID(block.GetID()),
			NodeID(newTmpIdentifier()),
		),
		parent: parent,
	}
}

// IsRoot returns true if the node is the root node
func (n *NodeBlock) IsRoot() bool { return n.parent == nil }

// GetID returns the ID of the node.d
func (n *NodeBlock) GetID() NodeID { return n.nodeID }

// GetType returns the type of the node.
func (n *NodeBlock) GetType() NodeType { return NodeType(n.block.GetType()) }

// GetBlock returns the block of the node.
func (n *NodeBlock) GetBlock() notion.Block { return n.block }

// GetPrevSibling returns the previous sibling node, or nil if there is none.
func (n *NodeBlock) GetPrevSibling() Node {
	if n.prev == nil {
		// Because we return Interface we
		// must return excplicitly nil, so later if v == nil does work
		return nil
	}
	return n.prev
}

// GetNextSibling returns the next sibling node, or nil if there is none.
func (n *NodeBlock) GetNextSibling() Node {
	if n.next == nil {
		// Because we return Interface we
		// must return excplicitly nil, so later if v == nil does work
		return nil
	}
	return n.next
}

// GetChildCount returns the number of child nodes.
func (n *NodeBlock) GetChildCount() int {
	count := 0
	for child := n.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		count++
	}
	return count
}

// GetFirstChild returns the first child node, or nil if there are no children.
func (n *NodeBlock) GetFirstChild() Node {
	if n.firstChild == nil {
		// Because we return Interface we
		// must return excplicitly nil, so later if v == nil does work
		return nil
	}
	return n.firstChild
}

// GetLastChild returns the last child node, or nil if there are no children.
func (n *NodeBlock) GetLastChild() Node {
	if n.lastChild == nil {
		// Because we return Interface we
		// must return excplicitly nil, so later if v == nil does work
		return nil
	}
	return n.lastChild
}

// GetParent returns the parent node, or nil if the node has no parent.
func (n *NodeBlock) GetParent() Node {
	if n.parent == nil {
		// Because we return Interface we
		// must return excplicitly nil, so later if v == nil does work
		return nil
	}
	return n.parent
}

// AppendChild appends a child node to the end of the list of children.
func (n *NodeBlock) AppendChild(newNode Node) {
	if newNode == nil {
		return
	}

	newChild := newNode.(*NodeBlock)

	newChild.parent = n

	if n.lastChild != nil {
		// Update sibling links
		n.lastChild.next = newChild
		newChild.prev = n.lastChild
	} else {
		// If no children exist, set the first newChild
		n.firstChild = newChild
	}

	// Update the last child reference
	n.lastChild = newChild
}

// RemoveChildren removes all child nodes from the node.
func (n *NodeBlock) RemoveChildren() {
	if n.GetFirstChild() == nil {
		return
	}

	// Iterate through all children and clear their links
	var cursor *NodeBlock
	for cursor = n.firstChild; cursor != nil; {
		nextCursor := cursor.next

		cursor.parent = nil
		cursor.next = nil
		cursor.prev = nil

		cursor = nextCursor
	}

	n.firstChild = nil
	n.lastChild = nil
}

// RemoveChild removes a child node (matched by ID) from the node.
func (n *NodeBlock) RemoveChild(childNode Node) {
	if childNode == nil || childNode.GetParent().GetID() != n.GetID() {
		return
	}

	child := childNode.(*NodeBlock)

	// Update sibling links
	if childPrev := child.GetPrevSibling(); childPrev != nil {
		childPrev.(*NodeBlock).next = child.GetNextSibling().(*NodeBlock)
	} else {
		// If the node is the first child, update the firstChild reference
		n.firstChild = child.GetNextSibling().(*NodeBlock)
	}

	if child.GetNextSibling() != nil {
		child.GetNextSibling().(*NodeBlock).prev = child.GetPrevSibling().(*NodeBlock)
	} else {
		// If the node is the last child, update the lastChild reference
		n.lastChild = child.GetPrevSibling().(*NodeBlock)
	}

	// Break links
	child.parent = nil
	child.prev = nil
	child.next = nil
}

// BlocksToAST creates a new AST from the given notion.Blocks
func BlocksToAST(blocks notion.Blocks, parentArg ...*NodeBlock) *NodeBlock {
	var parent *NodeBlock
	if len(parentArg) > 0 && parentArg[0] != nil {
		parent = parentArg[0]
	} else {
		parent = NewNodeBlock(nil, nil)
	}

	var children Nodes
	var prev *NodeBlock
	for _, block := range blocks {

		node := NewNodeBlock(block, parent)

		// Recursively process children if the block has them
		if block.GetHasChildren() {
			if deeper := BlocksToAST(block.(notion.HierarchicalBlock).GetChildren(), node); deeper.GetChildCount() > 0 {
				node.firstChild = deeper.GetFirstChild().(*NodeBlock)
				node.lastChild = deeper.GetLastChild().(*NodeBlock)
			}
		}

		// Set sibling relationships
		if prev != nil {
			node.prev = prev
			prev.next = node
		}
		prev = node

		children = append(children, node)
	}

	if len(children) > 0 {
		parent.firstChild = children[0].(*NodeBlock)
		parent.lastChild = children[len(children)-1].(*NodeBlock)
	}

	return parent
}

// ASTToBlocks converts the AST to notion.Blocks
func ASTToBlocks(n *NodeBlock) notion.Blocks {
	var blocks notion.Blocks

	// Helper function to recursively traverse the AST
	var traverse func(node *NodeBlock)
	traverse = func(node *NodeBlock) {
		if node == nil {
			return
		}

		if node.GetParent() == nil { // no parent means root
			traverse(node.GetFirstChild().(*NodeBlock))
			return
		}

		block := node.GetBlock()
		blocks = append(blocks, block)

		// Add children blocks recursively
		if child := node.GetFirstChild(); child != nil {
			children := ASTToBlocks(child.(*NodeBlock))
			block.(notion.HierarchicalBlock).SetChildren(children)
		}

		// Process the next sibling
		traverse(node.GetNextSibling().(*NodeBlock))
	}

	traverse(n)
	return blocks
}

// Walk traverses the AST and calls the given function for each node
func Walk(node Node, fn func(node Node)) {
	fn(node)

	for child := node.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		Walk(child, fn)
	}
}

// PrintAST prints the AST in a tree-like format
func PrintAST(node Node, levelArg ...int) {
	level := 0
	if len(levelArg) > 0 {
		level = levelArg[0]
	}

	fmt.Printf("%s- [%s] %s\n", indent(level), node.GetID(), node.GetType())

	for child := node.GetFirstChild(); child != nil; child = child.GetNextSibling() {
		PrintAST(child, level+1)
	}
}

//
// Helpers
//

// indent returns a string of spaces for the given level
func indent(level int) string {
	ident := " "
	for i := 0; i < level; i++ {
		ident += "  "
	}
	return ident
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var nCharset = big.NewInt(int64(len(charset)))

// newTmpIdentifier returns a temp random identifier
// All nodes do have generated IDs, so we can compare them (when finding a node to be deleted, etc)
func newTmpIdentifier() string {
	b := make([]byte, 10)
	for i := range b {
		r, _ := rand.Int(rand.Reader, nCharset)
		b[i] = charset[r.Uint64()]
	}

	return "tmp-" + string(b)
}
