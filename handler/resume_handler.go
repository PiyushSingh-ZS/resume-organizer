package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"resumeorganizer/model"
	"resumeorganizer/service"

	"github.com/gen2brain/go-fitz"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
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

// Create handles the request to create a new resume with both info and file
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
	resume, err := h.service.CreateResume(
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

		if err := h.service.UploadResumeFile(resume.ID, fileData, request.File.Filename); err != nil {
			return nil, err
		}
	}

	return resume, nil
}

// Get handles the request to retrieve a resume by ID
// Route: GET /resumes/{id}
// Query Parameters:
//   - include_content: boolean (optional) - If true, includes the PDF content in the response
func (h *ResumeHandler) Get(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	// Get resume information
	resume, err := h.service.GetResume(id)
	if err != nil {
		return nil, err
	}

	// Check if content should be included
	includeContent := ctx.Param("include_content") == "true"
	if !includeContent {
		return resume, nil
	}

	// If file exists, read its content
	if resume.FilePath != "" {
		// Open the PDF file
		doc, err := fitz.New(resume.FilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open PDF: %v", err)
		}
		defer doc.Close()

		// Extract text from all pages
		var text string
		for i := 0; i < doc.NumPage(); i++ {
			pageText, err := doc.Text(i)
			if err != nil {
				continue
			}
			text += pageText + "\n"
		}

		// Return resume with content
		return struct {
			*model.Resume
			Content string `json:"content,omitempty"`
		}{
			Resume:  resume,
			Content: text,
		}, nil
	}

	return resume, nil
}

// GetAll handles the request to retrieve all resumes
// Route: GET /resumes
// Query Parameters:
//   - role: string (optional) - Filter resumes by role
//   - company: string (optional) - Filter resumes by company
//   - status: string (optional) - Filter resumes by status
//   - include_content: boolean (optional) - If true, includes PDF content in the response
func (h *ResumeHandler) GetAll(ctx *gofr.Context) (interface{}, error) {
	role := ctx.Param("role")
	company := ctx.Param("company")
	status := ctx.Param("status")
	includeContent := ctx.Param("include_content") == "true"

	var resumes []*model.Resume
	var err error

	if role != "" {
		resumes, err = h.service.GetResumesByRole(role)
	} else if company != "" {
		resumes, err = h.service.GetResumesByCompany(company)
	} else {
		resumes, err = h.service.GetAllResumes()
	}

	if err != nil {
		return nil, err
	}

	// Filter by status if provided
	if status != "" {
		filteredResumes := make([]*model.Resume, 0)
		for _, resume := range resumes {
			if resume.Status == status {
				filteredResumes = append(filteredResumes, resume)
			}
		}
		resumes = filteredResumes
	}

	// Include content if requested
	if includeContent {
		response := make([]struct {
			*model.Resume
			Content string `json:"content,omitempty"`
		}, len(resumes))

		for i, resume := range resumes {
			response[i].Resume = resume
			if resume.FilePath != "" {
				doc, err := fitz.New(resume.FilePath)
				if err != nil {
					continue
				}
				defer doc.Close()

				var text string
				for i := 0; i < doc.NumPage(); i++ {
					pageText, err := doc.Text(i)
					if err != nil {
						continue
					}
					text += pageText + "\n"
				}

				response[i].Content = text
			}
		}

		return response, nil
	}

	return resumes, nil
}

// UpdateStatus handles the request to update a resume's status
func (h *ResumeHandler) UpdateStatus(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	var request struct {
		Status string `json:"status"`
		Notes  string `json:"notes"`
	}

	if err := ctx.Bind(&request); err != nil {
		return nil, err
	}

	if err := h.service.UpdateResumeStatus(id, request.Status, request.Notes); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Status updated successfully"}, nil
}

// UploadFile handles the request to upload a resume file
func (h *ResumeHandler) UploadFile(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	var request struct {
		FileHeader *multipart.FileHeader `file:"file"`
	}

	if err := ctx.Bind(&request); err != nil {
		return nil, err
	}

	file, err := request.FileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := h.service.UploadResumeFile(id, fileData, request.FileHeader.Filename); err != nil {
		return nil, err
	}

	return map[string]string{"message": "File uploaded successfully"}, nil
}

// Delete handles the request to delete a resume
func (h *ResumeHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	if err := h.service.DeleteResume(id); err != nil {
		return nil, err
	}

	return map[string]string{"message": "Resume deleted successfully"}, nil
}

// Download handles the request to download a resume file
// Route: GET /resumes/{id}/download
func (h *ResumeHandler) Download(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	resume, err := h.service.GetResume(id)
	if err != nil {
		return nil, err
	}

	if resume.FilePath == "" {
		return nil, fmt.Errorf("file not found for resume %s", id)
	}

	// Open the file
	file, err := os.Open(resume.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Return the file path and name for GoFr to handle the download
	return struct {
		FilePath string `json:"file_path"`
		FileName string `json:"file_name"`
	}{
		FilePath: resume.FilePath,
		FileName: resume.FileName,
	}, nil
}
