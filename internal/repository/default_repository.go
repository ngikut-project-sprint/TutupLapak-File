package repository

import "gorm.io/gorm"

type defaultFileRepository struct {
	db gorm.DB
}

func NewFileRepository(db gorm.DB) FileRepository {
	return &defaultFileRepository{db: db}
}
