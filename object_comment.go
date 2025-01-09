package notion

import (
	"time"
)

// CommentID stands for ID of Comment object.
// As Comment is an Object, then CommentID is just an alias for Object
type CommentID = ObjectID

// DiscussionID stands for ID of the Discussion.
type DiscussionID string

// String returns the string representation of the DiscussionID.
func (dID DiscussionID) String() string { return string(dID) }

// Comment is a Notion object that represents a comment.
type Comment struct {
	AtomObject
	AtomID
	AtomParent
	AtomCreated

	// LastEditedTime is declared manually, outside of AtomLastEdited,
	// because we do not have LastEditedBy field here.
	LastEditedTime *time.Time `json:"last_edited_time,omitempty"`

	DiscussionID DiscussionID `json:"discussion_id"`
	RichText     RichTexts    `json:"rich_text"`
}

// Comments is a slice of Comment objects.
type Comments []*Comment
