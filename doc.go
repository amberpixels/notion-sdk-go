// Package notion provides a simple and feature-rich Go client for the Notion API.
package notion

// Glossary:
// We refer `Object` to a Notion Object (that does have ID+Object fields)
// 	Block
// 	Database
// 	Page
//  User
//  Comment
// Such objects have their endpoints as well.
//
// Other objects we call EmbeddedObjects as they are still Objects (in Notion docs)
// but they won't have ID, probably won't have Object field, and are always embedded in other objects.
// They don't have their own endpoint groups, but they can have endpoints (e.g. Properties) in other's object (Page) endpoint group.
//
// Parent
// Unfurl Attribute (Link Preview)
// File
// Emoji
