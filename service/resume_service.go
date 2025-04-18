package service

import (
	"resumeorganizer/model"
	"resumeorganizer/store"

	"gofr.dev/pkg/gofr"
)

// ResumeService handles business logic for resume operations
type ResumeService struct {
	store *store.ResumeStore
}

// NewResumeService creates a new instance of ResumeService
func NewResumeService(store *store.ResumeStore) *ResumeService {
	return &ResumeService{
		store: store,
	}
}

// CreateResume creates a new resume
func (s *ResumeService) CreateResume(ctx *gofr.Context, role, company, version, status, notes string) (*model.Resume, error) {
	resume := model.NewResume(role, company, version, status, notes)
	err := s.store.Create(ctx, resume)
	if err != nil {
		return nil, err
	}
	return resume, nil
}

// GetResume retrieves a resume by ID
func (s *ResumeService) GetResume(ctx *gofr.Context, id string) (*model.Resume, error) {
	return s.store.GetByID(ctx, id)
}

// GetAllResumes retrieves all resumes
func (s *ResumeService) GetAllResumes(ctx *gofr.Context) ([]*model.Resume, error) {
	return s.store.GetAll(ctx)
}

// GetResumesByRole retrieves all resumes for a specific role
func (s *ResumeService) GetResumesByRole(ctx *gofr.Context, role string) ([]*model.Resume, error) {
	return s.store.GetByRole(ctx, role)
}

// GetResumesByCompany retrieves all resumes for a specific company
func (s *ResumeService) GetResumesByCompany(ctx *gofr.Context, company string) ([]*model.Resume, error) {
	return s.store.GetByCompany(ctx, company)
}

// UpdateResumeStatus updates the status and notes of a resume
func (s *ResumeService) UpdateResumeStatus(ctx *gofr.Context, id, status, notes string) error {
	return s.store.UpdateStatus(ctx, id, status, notes)
}

// UploadResumeFile uploads a resume file
func (s *ResumeService) UploadResumeFile(ctx *gofr.Context, id string, fileData []byte, fileName string) error {
	return s.store.UpdateFileInfo(ctx, id, fileData, fileName)
}

// DeleteResume removes a resume
func (s *ResumeService) DeleteResume(ctx *gofr.Context, id string) error {
	return s.store.Delete(ctx, id)
}
