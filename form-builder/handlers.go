package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateForm(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON payload
	var formData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Questions   []struct {
			Label        string `json:"label"`
			QuestionType string `json:"question_type"`
			Order        int    `json:"order"`
			Options      []struct {
				Value       string `json:"value"`
				DisplayText string `json:"display_text"`
			} `json:"options"`
			Required bool `json:"is_required,omitempty"`
		} `json:"questions"`
	}

	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Lock for thread-safe access
	Mutex.Lock()
	defer Mutex.Unlock()

	// Create a new Form
	form := Form{
		ID:          FormID,
		Title:       formData.Title,
		Description: formData.Description,
		CreatedAt:   time.Now(),
	}

	// Add Questions and Options
	var questionID uint = 1
	for _, questionData := range formData.Questions {
		question := Question{
			ID:           questionID,
			Label:        questionData.Label,
			QuestionType: questionData.QuestionType,
			Order:        questionData.Order,
			Required:     questionData.Required,
		}

		var optionID uint = 1
		for _, optionData := range questionData.Options {
			option := Option{
				ID:          optionID,
				Value:       optionData.Value,
				DisplayText: optionData.DisplayText,
			}
			question.Options = append(question.Options, option)
			optionID++
		}

		form.Questions = append(form.Questions, question)
		questionID++
	}

	// Store the form in memory
	Forms[FormID] = form
	FormID++

	// Return the created form as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(form)
}

func ListForms(w http.ResponseWriter, r *http.Request) {
	Mutex.Lock()
	defer Mutex.Unlock()

	// Convert the Forms map to a slice
	var formList []Form
	for _, form := range Forms {
		formList = append(formList, form)
	}

	// Return the Forms as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formList)
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON payload
	var submissionData struct {
		FormID    uint `json:"form_id"`
		Responses []struct {
			QuestionID uint   `json:"question_id"`
			AnswerText string `json:"answer_text"`
		} `json:"responses"`
	}

	if err := json.NewDecoder(r.Body).Decode(&submissionData); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Validate if the form exists
	Mutex.Lock()
	form, exists := Forms[submissionData.FormID]
	Mutex.Unlock()

	if !exists {
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}

	// Validate required questions
	for _, question := range form.Questions {
		if question.Required { // Check if the question is required
			found := false
			for _, response := range submissionData.Responses {
				if response.QuestionID == question.ID && response.AnswerText != "" {
					found = true
					break
				}
			}
			if !found {
				http.Error(w, "Required question not answered", http.StatusBadRequest)
				return
			}
		}
	}

	// Create a new Submission
	submission := Submission{
		ID:     SubmissionID,
		FormID: submissionData.FormID,
	}

	// Add Responses
	for _, responseData := range submissionData.Responses {
		response := Response{
			ID:         uint(len(submission.Responses) + 1),
			QuestionID: responseData.QuestionID,
			AnswerText: responseData.AnswerText,
		}
		submission.Responses = append(submission.Responses, response)
	}

	// Store the submission
	Mutex.Lock()
	Submissions[SubmissionID] = submission
	SubmissionID++
	Mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submission)
}

func GetFormResponses(w http.ResponseWriter, r *http.Request) {
	// Extract form_id from URL
	vars := mux.Vars(r)
	formIDStr, ok := vars["form_id"]
	if !ok {
		http.Error(w, "Form ID not provided", http.StatusBadRequest)
		return
	}

	// Convert form_id to integer
	formID, err := strconv.Atoi(formIDStr)
	if err != nil {
		http.Error(w, "Invalid Form ID", http.StatusBadRequest)
		return
	}

	// Lock for thread-safe access
	Mutex.Lock()
	form, formExists := Forms[uint(formID)]
	Mutex.Unlock()

	if !formExists {
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}

	// Fetch all submissions for the form
	var formSubmissions []Submission
	Mutex.Lock()
	for _, submission := range Submissions {
		if submission.FormID == uint(formID) {
			formSubmissions = append(formSubmissions, submission)
		}
	}
	Mutex.Unlock()

	// Build the response data structure
	type ResponseOutput struct {
		QuestionID uint   `json:"question_id"`
		AnswerText string `json:"answer_text"`
	}

	type SubmissionOutput struct {
		SubmissionID uint             `json:"submission_id"`
		Responses    []ResponseOutput `json:"responses"`
	}

	type FormResponses struct {
		FormID      uint               `json:"form_id"`
		Title       string             `json:"title"`
		Description string             `json:"description"`
		Submissions []SubmissionOutput `json:"submissions"`
	}

	formResponses := FormResponses{
		FormID:      form.ID,
		Title:       form.Title,
		Description: form.Description,
	}

	// Populate submissions and responses
	for _, submission := range formSubmissions {
		submissionOutput := SubmissionOutput{
			SubmissionID: submission.ID,
		}
		for _, response := range submission.Responses {
			submissionOutput.Responses = append(submissionOutput.Responses, ResponseOutput{
				QuestionID: response.QuestionID,
				AnswerText: response.AnswerText,
			})
		}
		formResponses.Submissions = append(formResponses.Submissions, submissionOutput)
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formResponses)
}
