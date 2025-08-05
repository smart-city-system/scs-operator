package models

import (
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	Base
	AssetID     uuid.UUID `json:"asset_id" gorm:"not null"`
	Asset       Asset     `gorm:"foreignKey:AssetID" json:"asset"`
	Description string    `json:"description" gorm:"type:text"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null"`
	FileName    string    `json:"file_name" gorm:"not null"`
	FileSize    int64     `json:"file_size" gorm:"not null"`
	FileType    string    `json:"file_type" gorm:"not null"`
	FilePath    string    `json:"file_path" gorm:"not null"`
}
