package repository

import "database/sql"

type ContactRepository interface {
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) ContactRepository {
	return &Repository{DB: db}
}
