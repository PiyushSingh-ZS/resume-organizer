package store

import (
	"resumeorganizer/model"

	"gofr.dev/pkg/gofr"
)

// ResumeStore defines the interface for resume storage operations
type ResumeStore struct {
}

// NewResumeStore creates a new instance of ResumeStore
func NewResumeStore() *ResumeStore {
	return &ResumeStore{}
}

// Create adds a new resume to the store
func (s *ResumeStore) Create(ctx *gofr.Context, resume *model.Resume) error {
	_, err := ctx.SQL.ExecContext(ctx, `
		INSERT INTO resumes (id, role, company, version, status, notes, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())
	`, resume.ID, resume.Role, resume.Company, resume.Version, resume.Status, resume.Notes)
	return err
}

// GetByID retrieves a resume by its ID
func (s *ResumeStore) GetByID(ctx *gofr.Context, id string) (*model.Resume, error) {
	var resume model.Resume
	err := ctx.SQL.QueryRowContext(ctx, `
		SELECT id, role, company, version, status, notes, file_name, file_content, created_at, updated_at
		FROM resumes
		WHERE id = ?
	`, id).Scan(
		&resume.ID,
		&resume.Role,
		&resume.Company,
		&resume.Version,
		&resume.Status,
		&resume.Notes,
		&resume.FileName,
		&resume.FileContent,
		&resume.CreatedAt,
		&resume.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}

// GetAll retrieves all resumes
func (s *ResumeStore) GetAll(ctx *gofr.Context) ([]*model.Resume, error) {
	rows, err := ctx.SQL.QueryContext(ctx, `
		SELECT id, role, company, version, status, notes, file_name, file_content, created_at, updated_at
		FROM resumes
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []*model.Resume
	for rows.Next() {
		var resume model.Resume
		err := rows.Scan(
			&resume.ID,
			&resume.Role,
			&resume.Company,
			&resume.Version,
			&resume.Status,
			&resume.Notes,
			&resume.FileName,
			&resume.FileContent,
			&resume.CreatedAt,
			&resume.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}
	return resumes, nil
}

// GetByRole retrieves all resumes for a specific role
func (s *ResumeStore) GetByRole(ctx *gofr.Context, role string) ([]*model.Resume, error) {
	rows, err := ctx.SQL.QueryContext(ctx, `
		SELECT id, role, company, version, status, notes, file_name, file_content, created_at, updated_at
		FROM resumes
		WHERE role = ?
	`, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []*model.Resume
	for rows.Next() {
		var resume model.Resume
		err := rows.Scan(
			&resume.ID,
			&resume.Role,
			&resume.Company,
			&resume.Version,
			&resume.Status,
			&resume.Notes,
			&resume.FileName,
			&resume.FileContent,
			&resume.CreatedAt,
			&resume.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}
	return resumes, nil
}

// GetByCompany retrieves all resumes for a specific company
func (s *ResumeStore) GetByCompany(ctx *gofr.Context, company string) ([]*model.Resume, error) {
	rows, err := ctx.SQL.QueryContext(ctx, `
		SELECT id, role, company, version, status, notes, file_name, file_content, created_at, updated_at
		FROM resumes
		WHERE company = ?
	`, company)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []*model.Resume
	for rows.Next() {
		var resume model.Resume
		err := rows.Scan(
			&resume.ID,
			&resume.Role,
			&resume.Company,
			&resume.Version,
			&resume.Status,
			&resume.Notes,
			&resume.FileName,
			&resume.FileContent,
			&resume.CreatedAt,
			&resume.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}
	return resumes, nil
}

// UpdateStatus updates the status and notes of a resume
func (s *ResumeStore) UpdateStatus(ctx *gofr.Context, id, status, notes string) error {
	_, err := ctx.SQL.ExecContext(ctx, `
		UPDATE resumes 
		SET status = ?, notes = ?, updated_at = NOW()
		WHERE id = ?
	`, status, notes, id)
	return err
}

// UpdateFileInfo updates the file information of a resume
func (s *ResumeStore) UpdateFileInfo(ctx *gofr.Context, id string, fileData []byte, fileName string) error {
	_, err := ctx.SQL.ExecContext(ctx, `
		UPDATE resumes 
		SET file_name = ?, file_content = ?, updated_at = NOW()
		WHERE id = ?
	`, fileName, fileData, id)
	return err
}

// Delete removes a resume from the store
func (s *ResumeStore) Delete(ctx *gofr.Context, id string) error {
	_, err := ctx.SQL.ExecContext(ctx, "DELETE FROM resumes WHERE id = ?", id)
	return err
}
