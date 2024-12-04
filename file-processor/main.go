package main

import (
	"database/sql"
	"time"
)

type FileMetadata struct {
	FileId          string
	Name            string
	Type            string
	Size            string
	UploadTimestamp time.Time
}

type MetaDataStore struct {
	db *sql.DB
}

func NewMetadataStore(dataSourceName string) (*MetaDataStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &MetaDataStore{db: db},nil
}
