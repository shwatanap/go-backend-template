package repository

import (
	"database/sql"

	"github.com/shwatanap/go-backend-template/srcs/domain/repository"
)

type templateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) repository.ITemplateRepository {
	return &templateRepository{db: db}
}
