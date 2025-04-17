# Resume Organizer

A modern resume management system built with GoFr framework that helps you organize and manage different versions of your resume for various job roles.

## Features

- **Role-based Resume Management**
  - Store multiple resumes for different job roles
  - Associate metadata with each resume (e.g., target company, role, version)
  - Track resume versions and updates

- **Resume Information Management**
  - Store detailed information about each resume
  - Track application status for each resume
  - Add notes and feedback for each version

- **File Management**
  - Upload and store resume files (PDF, DOCX)
  - Download resumes for specific roles
  - Version control for resume files

- **Search and Filter**
  - Search resumes by role, company, or keywords
  - Filter resumes by status or date
  - Quick access to recent resumes

## API Endpoints

### Resume Management
```
POST /resumes
Content-Type: application/json

{
    "role": "Software Engineer",
    "company": "Tech Corp",
    "version": "1.0",
    "status": "draft",
    "notes": "Initial version for backend positions"
}
```

### File Upload
```
POST /resumes/{id}/upload
Content-Type: multipart/form-data

file: [resume file]
```

### Resume Information
```
GET /resumes/{id}
GET /resumes?role=backend&company=tech
```

### Update Status
```
PATCH /resumes/{id}/status
Content-Type: application/json

{
    "status": "submitted",
    "notes": "Submitted for review"
}
```

## Data Model

### Resume
```go
type Resume struct {
    ID          string    `json:"id"`
    Role        string    `json:"role"`
    Company     string    `json:"company"`
    Version     string    `json:"version"`
    Status      string    `json:"status"`
    Notes       string    `json:"notes"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    FilePath    string    `json:"file_path"`
    FileName    string    `json:"file_name"`
}
```

## Architecture

The application follows a 3-layer architecture:

1. **Handler Layer**
   - Handles HTTP requests
   - Validates input
   - Manages responses
   - Coordinates with service layer

2. **Service Layer**
   - Contains business logic
   - Manages resume operations
   - Handles file operations
   - Coordinates with store layer

3. **Store Layer**
   - Manages data persistence
   - Handles file storage
   - Provides CRUD operations

## Getting Started

1. Install dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

The server will start on port 8000 by default.

## Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 8000)
- `FILE_STORAGE_PATH`: Path to store resume files
- `MAX_FILE_SIZE`: Maximum file size for uploads (default: 5MB)

## Future Enhancements

- Resume comparison tool
- Resume template suggestions
- Integration with job boards
- Automated resume formatting
- Resume analytics and tracking
- Collaboration features for feedback

## Note

This implementation uses in-memory storage for demonstration purposes. For production use, consider using:
- A database for resume metadata
- A file storage service for resume files
- Authentication and authorization
- Rate limiting and security measures 