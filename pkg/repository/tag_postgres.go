package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
)

type TagPostgres struct {
	db *sqlx.DB
}

func NewTagPostgres(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) Create(tag banner.Tag) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (tag_id) VALUES ($1) RETURNING tag_id", tagsTable)

	row := r.db.QueryRow(query, tag.ID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
