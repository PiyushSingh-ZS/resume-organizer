package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"resumeorganizer/model"
	"resumeorganizer/service"

	"gofr.dev/pkg/gofr"
)

// ResumeHandler handles HTTP requests for resume operations
type ResumeHandler struct {
	service *service.ResumeService
}

// NewResumeHandler creates a new instance of ResumeHandler
func NewResumeHandler(service *service.ResumeService) *ResumeHandler {
	return &ResumeHandler{
		service: service,
	}
}

// Create handles the creation of a new resume
func (h *ResumeHandler) Create(ctx *gofr.Context) (interface{}, error) {
	var request struct {
		Role    string                `form:"role"`
		Company string                `form:"company"`
		Version string                `form:"version"`
		Status  string                `form:"status"`
		Notes   string                `form:"notes"`
		File    *multipart.FileHeader `file:"file"`
	}

	if err := ctx.Bind(&request); err != nil {
		return nil, err
	}

	// Create resume with basic info
	resume, err := h.service.CreateResume(ctx,
		request.Role,
		request.Company,
		request.Version,
		request.Status,
		request.Notes,
	)
	if err != nil {
		return nil, err
	}

	// If file is provided, upload it
	if request.File != nil {
		file, err := request.File.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fileData, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		if err := h.service.UploadResumeFile(ctx, resume.ID, fileData, request.File.Filename); err != nil {
			return nil, err
		}
	}

	return resume, nil
}

// Get handles retrieving a resume by ID
// Route: GET /resumes/{id}
// Query Parameters:
//   - include_content: boolean (optional) - If true, includes the PDF content in the response
func (h *ResumeHandler) Get(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	// Get resume information
	resume, err := h.service.GetResume(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if content should be included
	includeContent := ctx.Param("include_content") == "true"
	if !includeContent {
		return resume, nil
	}

	// If file exists, return it with content
	if len(resume.FileContent) > 0 {
		return struct {
			*model.Resume
			Content string `json:"content,omitempty"`
		}{
			Resume:  resume,
			Content: string(resume.FileContent),
		}, nil
	}

	return resume, nil
}

// GetAll handles retrieving all resumes
// Route: GET /resumes
// Query Parameters:
//   - role: string (optional) - Filter resumes by role
//   - company: string (optional) - Filter resumes by company
//   - status: string (optional) - Filter resumes by status
//   - include_content: boolean (optional) - If true, includes PDF content in the response
func (h *ResumeHandler) GetAll(ctx *gofr.Context) (interface{}, error) {
	role := ctx.Param("role")
	company := ctx.Param("company")
	includeContent := ctx.Param("include_content") == "true"

	var resumes []*model.Resume
	var err error

	if role != "" {
		resumes, err = h.service.GetResumesByRole(ctx, role)
	} else if company != "" {
		resumes, err = h.service.GetResumesByCompany(ctx, company)
	} else {
		resumes, err = h.service.GetAllResumes(ctx)
	}

	if err != nil {
		return nil, err
	}

	// Include content if requested
	if includeContent {
		response := make([]struct {
			*model.Resume
			Content string `json:"content,omitempty"`
		}, len(resumes))

		for i, resume := range resumes {
			response[i].Resume = resume
			if len(resume.FileContent) > 0 {
				response[i].Content = string(resume.FileContent)
			}
		}

		return response, nil
	}

	return resumes, nil
}

// UpdateStatus handles updating a resume's status
func (h *ResumeHandler) UpdateStatus(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var req struct {
		Status string `json:"status"`
		Notes  string `json:"notes"`
	}

	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}

	err := h.service.UpdateResumeStatus(ctx, id, req.Status, req.Notes)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Resume status updated successfully"}, nil
}

// Delete handles deleting a resume
func (h *ResumeHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	err := h.service.DeleteResume(ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Resume deleted successfully"}, nil
}
