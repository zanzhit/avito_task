package repository

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
)

type UserBannerPostgres struct {
	db *sqlx.DB
}

func NewUserBannerPostgres(db *sqlx.DB) *UserBannerPostgres {
	return &UserBannerPostgres{db: db}
}

func (r *UserBannerPostgres) GetUserBanner(ban banner.UserBanner, lastRevision bool) (interface{}, error) {
	var userBanner banner.UserBanner

	query := fmt.Sprintf(`SELECT b.content FROM %s b INNER JOIN %s bt ON b.id = bt.banner_id WHERE bt.tags_id = ($1) AND b.feature_id = $2`, bannersTable, bannerTagsTable)
	r.db.Get(&userBanner, query, ban.Tag, ban.Feature)

	var content interface{}
	err := json.Unmarshal(userBanner.Content, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
