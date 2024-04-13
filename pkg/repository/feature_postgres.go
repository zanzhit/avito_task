package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
)

type FeaturePostgres struct {
	db *sqlx.DB
}

func NewFeaturePostgres(db *sqlx.DB) *FeaturePostgres {
	return &FeaturePostgres{db: db}
}

func (r *FeaturePostgres) Create(feature banner.Feature) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (feature_id) VALUES ($1) RETURNING feature_id", featureTable)

	row := r.db.QueryRow(query, feature.ID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
