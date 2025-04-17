package store

import (
	"errors"
	"resumeorganizer/model"
	"sync"
)

// ResumeStore defines the interface for resume storage operations
type ResumeStore interface {
	Create(resume *model.Resume) error
	GetByID(id string) (*model.Resume, error)
	GetAll() ([]*model.Resume, error)
	GetByRole(role string) ([]*model.Resume, error)
	GetByCompany(company string) ([]*model.Resume, error)
	UpdateStatus(id, status, notes string) error
	UpdateFileInfo(id, filePath, fileName string) error
	Delete(id string) error
}

// InMemoryResumeStore implements ResumeStore using in-memory storage
type InMemoryResumeStore struct {
	resumes map[string]*model.Resume
	mu      sync.RWMutex
}

// NewInMemoryResumeStore creates a new instance of InMemoryResumeStore
func NewInMemoryResumeStore() *InMemoryResumeStore {
	return &InMemoryResumeStore{
		resumes: make(map[string]*model.Resume),
	}
}

// Create adds a new resume to the store
func (s *InMemoryResumeStore) Create(resume *model.Resume) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.resumes[resume.ID]; exists {
		return errors.New("resume already exists")
	}

	s.resumes[resume.ID] = resume
	return nil
}

// GetByID retrieves a resume by its ID
func (s *InMemoryResumeStore) GetByID(id string) (*model.Resume, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resume, exists := s.resumes[id]
	if !exists {
		return nil, errors.New("resume not found")
	}

	return resume, nil
}

// GetAll retrieves all resumes
func (s *InMemoryResumeStore) GetAll() ([]*model.Resume, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resumes := make([]*model.Resume, 0, len(s.resumes))
	for _, resume := range s.resumes {
		resumes = append(resumes, resume)
	}

	return resumes, nil
}

// GetByRole retrieves all resumes for a specific role
func (s *InMemoryResumeStore) GetByRole(role string) ([]*model.Resume, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resumes := make([]*model.Resume, 0)
	for _, resume := range s.resumes {
		if resume.Role == role {
			resumes = append(resumes, resume)
		}
	}

	return resumes, nil
}

// GetByCompany retrieves all resumes for a specific company
func (s *InMemoryResumeStore) GetByCompany(company string) ([]*model.Resume, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resumes := make([]*model.Resume, 0)
	for _, resume := range s.resumes {
		if resume.Company == company {
			resumes = append(resumes, resume)
		}
	}

	return resumes, nil
}

// UpdateStatus updates the status and notes of a resume
func (s *InMemoryResumeStore) UpdateStatus(id, status, notes string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	resume, exists := s.resumes[id]
	if !exists {
		return errors.New("resume not found")
	}

	resume.UpdateStatus(status, notes)
	return nil
}

// UpdateFileInfo updates the file information of a resume
func (s *InMemoryResumeStore) UpdateFileInfo(id, filePath, fileName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	resume, exists := s.resumes[id]
	if !exists {
		return errors.New("resume not found")
	}

	resume.SetFileInfo(filePath, fileName)
	return nil
}

// Delete removes a resume from the store
func (s *InMemoryResumeStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.resumes[id]; !exists {
		return errors.New("resume not found")
	}

	delete(s.resumes, id)
	return nil
}
