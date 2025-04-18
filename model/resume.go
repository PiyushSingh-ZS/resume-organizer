package model

import (
	"time"
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
	FileContent []byte    `json:"file_content,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
