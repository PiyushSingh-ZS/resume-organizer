package service

import (
	"errors"
	"os"
	"path/filepath"
	"resumeorganizer/model"
	"resumeorganizer/store"
)

// ResumeService handles business logic for resume operations
type ResumeService struct {
	store store.ResumeStore
}

// NewResumeService creates a new instance of ResumeService
func NewResumeService(store store.ResumeStore) *ResumeService {
	return &ResumeService{
		store: store,
	}
}

// CreateResume creates a new resume with the provided information
func (s *ResumeService) CreateResume(role, company, version, status, notes string) (*model.Resume, error) {
	if role == "" || company == "" {
		return nil, errors.New("role and company are required")
	}

	resume := model.NewResume(role, company, version, status, notes)
	if err := s.store.Create(resume); err != nil {
		return nil, err
	}

	return resume, nil
}

// GetResume retrieves a resume by its ID
func (s *ResumeService) GetResume(id string) (*model.Resume, error) {
	return s.store.GetByID(id)
}

// GetAllResumes retrieves all resumes
func (s *ResumeService) GetAllResumes() ([]*model.Resume, error) {
	return s.store.GetAll()
}

// GetResumesByRole retrieves all resumes for a specific role
func (s *ResumeService) GetResumesByRole(role string) ([]*model.Resume, error) {
	return s.store.GetByRole(role)
}

// GetResumesByCompany retrieves all resumes for a specific company
func (s *ResumeService) GetResumesByCompany(company string) ([]*model.Resume, error) {
	return s.store.GetByCompany(company)
}

// UpdateResumeStatus updates the status and notes of a resume
func (s *ResumeService) UpdateResumeStatus(id, status, notes string) error {
	return s.store.UpdateStatus(id, status, notes)
}

// UploadResumeFile handles the file upload and updates the resume's file information
func (s *ResumeService) UploadResumeFile(id string, fileData []byte, fileName string) error {
	// Create uploads directory if it doesn't exist
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}

	// Generate unique file path
	filePath := filepath.Join(uploadDir, id+"_"+fileName)

	// Write file to disk
	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return err
	}

	// Update resume's file information
	return s.store.UpdateFileInfo(id, filePath, fileName)
}

// DeleteResume removes a resume and its associated file
func (s *ResumeService) DeleteResume(id string) error {
	// Get resume to check if it has an associated file
	resume, err := s.store.GetByID(id)
	if err != nil {
		return err
	}

	// Delete associated file if it exists
	if resume.FilePath != "" {
		if err := os.Remove(resume.FilePath); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	// Delete resume from store
	return s.store.Delete(id)
}
