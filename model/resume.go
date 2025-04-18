package model

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gen2brain/go-fitz"
)

// Resume represents a resume document with its metadata
type Resume struct {
	ID          string    `json:"id"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Version     string    `json:"version"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
	FileName    string    `json:"file_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	FileContent []byte    `json:"content,omitempty"`
}

// extractTextFromPDF extracts text content from PDF bytes
func extractTextFromPDF(pdfData []byte) (string, error) {
	doc, err := fitz.NewFromMemory(pdfData)
	if err != nil {
		return "", err
	}
	defer doc.Close()

	var text bytes.Buffer
	for i := 0; i < doc.NumPage(); i++ {
		pageText, err := doc.Text(i)
		if err != nil {
			continue
		}
		text.WriteString(pageText)
		text.WriteString("\n")
	}

	return text.String(), nil
}

// MarshalJSON customizes the JSON marshaling of Resume
func (r Resume) MarshalJSON() ([]byte, error) {
	type Alias Resume
	var content string
	if len(r.FileContent) > 0 {
		text, err := extractTextFromPDF(r.FileContent)
		if err != nil {
			content = "Error extracting text from PDF"
		} else {
			content = text
		}
	}
	return json.Marshal(&struct {
		Content string `json:"content,omitempty"`
		Alias
	}{
		Content: content,
		Alias:   (Alias)(r),
	})
}

// UnmarshalJSON customizes the JSON unmarshaling of Resume
func (r *Resume) UnmarshalJSON(data []byte) error {
	type Alias Resume
	aux := &struct {
		Content string `json:"content,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Content != "" {
		r.FileContent = []byte(aux.Content)
	}
	return nil
}

// NewResume creates a new Resume instance with default values
func NewResume(role, company, version, status, notes string) *Resume {
	now := time.Now()
	return &Resume{
		ID:        generateID(),
		Role:      role,
		Company:   company,
		Version:   version,
		Status:    status,
		Notes:     notes,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateStatus updates the resume status and notes
func (r *Resume) UpdateStatus(status, notes string) {
	r.Status = status
	r.Notes = notes
	r.UpdatedAt = time.Now()
}

// SetFileInfo updates the file information
func (r *Resume) SetFileInfo(fileName string, fileContent []byte) {
	r.FileName = fileName
	r.FileContent = fileContent
	r.UpdatedAt = time.Now()
}

// generateID generates a unique ID for the resume
func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

// randomString generates a random string of specified length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
