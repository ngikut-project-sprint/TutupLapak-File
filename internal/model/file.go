package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	FileID           uint           `gorm:"primaryKey;autoIncrement"`
	FileURI          string         `gorm:"type:text;not null"`
	FileThumbnailURI string         `gorm:"type:text;not null"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type AddFileResponse struct {
	FileID           string `json:"fileId"`
	FileURI          string `json:"fileUri"`
	FileThumbnailURI string `json:"fileThumbnailUri"`
}
