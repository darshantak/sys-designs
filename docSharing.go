// Design and implement a system that allows users to:

// Upload and share documents with other users.
// Manage permissions (read, write, delete) for shared documents.
// Track access logs to see who accessed the document and when.
// Handle concurrent access and ensure consistency when documents are modified.

package main

import "time"

type AccessType string

const (
	Admin AccessType = "admin"
	User  AccessType = "user"
)

type Doc struct {
	Id         string
	Name       string
	CreatedOn  time.Time
	AccessLogs []AccessLog
}
type AccessLog struct {
	UserId     string
	AccessType AccessType
	Time       string
}

type DocumentSystem struct {
	Documents map[string]Document
}

func NewDocumentSystem() *DocumentSystem {
	return &DocumentSystem{
		Documents: make(map[string]Document),
	}
}

func (ds *DocumentSystem) CreateDocument(id, name string) {
	// ds.Documents[]
}
