package dto

import (
	"mime/multipart"
	"time"
)

type SnapshotRequest struct {
	AssetID     string                `form:"asset_id" validate:"required,uuid"`
	Description string                `form:"description" validate:"max=500"`
	Timestamp   *time.Time            `form:"timestamp"`
	File        *multipart.FileHeader `form:"file" validate:"required"`
}

type SnapshotResponse struct {
	ID          string    `json:"id"`
	AssetID     string    `json:"asset_id"`
	Description string    `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
	FileName    string    `json:"file_name"`
	FileSize    int64     `json:"file_size"`
	FileType    string    `json:"file_type"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
}
