package repository

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/errs"
)

type UserBannerPostgres struct {
	db *sqlx.DB
}

func NewUserBannerPostgres(db *sqlx.DB) *UserBannerPostgres {
	return &UserBannerPostgres{db: db}
}

func (r *UserBannerPostgres) GetUserBanner(ban banner.UserBanner) (interface{}, error) {
	var userBanner banner.UserBanner

	existQuery := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s b INNER JOIN %s bt ON b.id = bt.banner_id WHERE bt.tags_id = $1 AND b.feature_id = $2)`, bannersTable, bannerTagsTable)
	exists := false
	err := r.db.QueryRow(existQuery, ban.Tag, ban.Feature).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &errs.ErrBannerNotFound{}
	}

	query := fmt.Sprintf(`SELECT b.content FROM %s b INNER JOIN %s bt ON b.id = bt.banner_id WHERE bt.tags_id = ($1) AND b.feature_id = $2`, bannersTable, bannerTagsTable)
	r.db.Get(&userBanner, query, ban.Tag, ban.Feature)

	if userBanner.Content == nil {
		return nil, nil
	}

	var content interface{}
	err = json.Unmarshal(userBanner.Content, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
