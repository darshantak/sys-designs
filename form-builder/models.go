package main

import (
	"sync"
	"time"
)

// Structs (same as earlier)
type Form struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FkId string 
	// Questions   []Question `json:"questions"`
	CreatedAt   time.Time  `json:"created_at"`
}

// fk(FormID -> QuestionID)

type Question struct {
	ID           uint     `json:"id"`
	Label        string   `json:"label"`
	QuestionType string   `json:"question_type"`
	Order        int      `json:"order"`
	Options      []Option `json:"options"`
	Required     bool     `json:"required"`
	Properties interface{} 
	FkId 		string `js`
}

type Option struct {
	ID          uint   `json:"id"`
	Value       string `json:"value"`
	DisplayText string `json:"display_text"`
	
}

type Submission struct {
	ID        uint       `json:"id"`
	FormID    uint       `json:"form_id"`
	Responses []Response `json:"responses"`
}

type Response struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	AnswerText string `json:"answer_text"`
}

// In-memory storage
var (
	Forms                   = make(map[uint]Form)
	Submissions             = make(map[uint]Submission)
	FormID       uint       = 1
	SubmissionID uint       = 1
	Mutex        sync.Mutex // For thread-safe access
)
