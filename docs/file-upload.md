# File Upload and Extraction Guide

This document explains how to handle file uploads in the smart-city application, specifically for camera snapshots.

## Overview

The file upload system provides:
- **Multipart form handling** for file uploads
- **File validation** (type, size, format)
- **Secure file storage** with unique naming
- **Database integration** for file metadata
- **Error handling** for upload failures

## API Endpoint

### POST `/api/v1/cameras/snapshot`

Upload a camera snapshot with metadata.

#### Request Format

**Content-Type:** `multipart/form-data`

**Form Fields:**
- `camera_id` (required): UUID of the camera
- `description` (optional): Description of the snapshot
- `timestamp` (optional): When the snapshot was taken
- `file` (required): Image file (JPEG, PNG, GIF, WebP, BMP)

#### Example using cURL

```bash
curl -X POST http://localhost:8080/api/v1/cameras/snapshot \
  -F "camera_id=123e4567-e89b-12d3-a456-426614174000" \
  -F "description=Security checkpoint at main entrance" \
  -F "file=@/path/to/snapshot.jpg"
```

#### Example using JavaScript/Fetch

```javascript
const formData = new FormData();
formData.append('camera_id', '123e4567-e89b-12d3-a456-426614174000');
formData.append('description', 'Security checkpoint at main entrance');
formData.append('file', fileInput.files[0]);

fetch('/api/v1/cameras/snapshot', {
  method: 'POST',
  body: formData
})
.then(response => response.json())
.then(data => console.log(data));
```

#### Response Format

```json
{
  "id": "456e7890-e89b-12d3-a456-426614174001",
  "camera_id": "123e4567-e89b-12d3-a456-426614174000",
  "description": "Security checkpoint at main entrance",
  "timestamp": "2024-01-01T12:00:00Z",
  "file_name": "20240101_120000_abc12345.jpg",
  "file_size": 2048576,
  "file_type": "image/jpeg",
  "file_path": "./uploads/snapshots/123e4567-e89b-12d3-a456-426614174000/20240101_120000_abc12345.jpg",
  "created_at": "2024-01-01T12:00:00Z"
}
```

## File Validation

### Supported Image Types
- JPEG/JPG (`image/jpeg`)
- PNG (`image/png`)
- GIF (`image/gif`)
- WebP (`image/webp`)
- BMP (`image/bmp`)

### File Size Limits
- Maximum file size: **10MB**
- Form size limit: **10MB**

### Validation Errors

```json
{
  "error": {
    "type": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "file",
        "message": "file size 15728640 bytes exceeds maximum allowed size 10485760 bytes"
      }
    ]
  },
  "request_id": "req-123456",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## File Storage Structure

Files are stored in the following directory structure:

```
uploads/
└── snapshots/
    └── {camera_id}/
        ├── 20240101_120000_abc12345.jpg
        ├── 20240101_130000_def67890.png
        └── ...
```

### File Naming Convention

Files are renamed using the pattern: `{timestamp}_{unique_id}.{extension}`

- `timestamp`: Format `YYYYMMDD_HHMMSS`
- `unique_id`: 8-character UUID segment
- `extension`: Original file extension

## Implementation Details

### Controller Layer

The controller handles:
1. **Multipart form parsing**
2. **Form field extraction**
3. **File extraction**
4. **Request validation**
5. **Service delegation**

```go
func (h *Handler) CreateSnapshot() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Parse multipart form
        if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
            return errors.NewBadRequestError("Failed to parse multipart form")
        }

        // Extract form data
        snapShotRequest := &dto.SnapshotRequest{}
        snapShotRequest.CameraID = c.FormValue("camera_id")
        snapShotRequest.Description = c.FormValue("description")

        // Extract file
        file, err := c.FormFile("file")
        if err != nil {
            return errors.NewBadRequestError("File is required")
        }
        snapShotRequest.File = file

        // Validate and process
        if err := validation.ValidateStruct(snapShotRequest); err != nil {
            return err
        }

        snapshot, err := h.svc.CreateSnapshot(c.Request().Context(), snapShotRequest, uploadDir)
        if err != nil {
            return err
        }

        return c.JSON(201, snapshot)
    }
}
```

### Service Layer

The service handles:
1. **Business logic validation**
2. **File validation**
3. **File storage**
4. **Database operations**
5. **Cleanup on failure**

### File Utilities

The `pkg/utils/file.go` package provides:
- `SaveUploadedFile()`: Save multipart files securely
- `ValidateImageFile()`: Validate image files
- `ValidateVideoFile()`: Validate video files
- `DeleteFile()`: Clean up files safely

## Error Handling

### Common Error Scenarios

1. **Missing file**: `File is required`
2. **Invalid file type**: `file type image/svg+xml is not allowed`
3. **File too large**: `file size exceeds maximum allowed size`
4. **Invalid camera ID**: `Invalid camera ID format`
5. **Camera not found**: `camera not found`
6. **Storage failure**: `Failed to save uploaded file`

### Error Response Format

All errors follow the standard error response format:

```json
{
  "error": {
    "type": "ERROR_TYPE",
    "message": "Human readable message",
    "details": "Additional error details"
  },
  "request_id": "req-123456",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Security Considerations

1. **File type validation**: Only allowed image types are accepted
2. **File size limits**: Prevents large file uploads
3. **Unique file names**: Prevents file name conflicts
4. **Directory isolation**: Files are stored in camera-specific directories
5. **Error handling**: Failed uploads are cleaned up automatically

## Configuration

### Upload Directory

The upload directory can be configured in the controller:

```go
uploadDir := "./uploads" // Make this configurable via environment variables
```

### File Size Limits

File size limits are defined in the service:

```go
const maxFileSize = 10 * 1024 * 1024 // 10MB
```

## Testing

### Test File Upload

```bash
# Create a test image
echo "test image content" > test.jpg

# Upload the file
curl -X POST http://localhost:8080/api/v1/cameras/snapshot \
  -F "camera_id=123e4567-e89b-12d3-a456-426614174000" \
  -F "description=Test snapshot" \
  -F "file=@test.jpg"
```

### Validate Response

Check that the response includes:
- Unique snapshot ID
- Correct camera ID
- File metadata (name, size, type)
- File path
- Creation timestamp
