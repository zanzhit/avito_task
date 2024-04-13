package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
)

type BannerPostgres struct {
	db *sqlx.DB
}

func NewBannerPostgres(db *sqlx.DB) *BannerPostgres {
	return &BannerPostgres{db: db}
}

func (r *BannerPostgres) Create(banner banner.Banner) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	bannerContent, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, err
	}

	createUdpateTime := time.Now()

	var id int
	createBannerQuery := fmt.Sprintf("INSERT INTO %s (content, is_active, feature_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", bannersTable)
	row := tx.QueryRow(createBannerQuery, bannerContent, banner.IsActive, banner.Feature, createUdpateTime, createUdpateTime)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createBannerTagsQuery := fmt.Sprintf("INSERT INTO %s (banner_id, tags_id) VALUES ($1, $2)", bannerTagsTable)
	for _, tag := range banner.Tag {
		_, err = tx.Exec(createBannerTagsQuery, id, tag)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *BannerPostgres) GetBanner(ban banner.Banner, limit, offset int) ([]banner.Banner, error) {
	var banners []banner.Banner
	if len(ban.Tag) == 0 && ban.Feature == 0 {
		return nil, errors.New("tags and feature missing")
	}

	limitOffset := ""
	if limit > 0 && offset >= 0 {
		limitOffset = fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	} else if limit > 0 {
		limitOffset = fmt.Sprintf("LIMIT %d", limit)
	} else if offset >= 0 {
		limitOffset = fmt.Sprintf("OFFSET %d", offset)
	}

	var query string

	if len(ban.Tag) != 0 && ban.Feature == 0 {
		tags := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ban.Tag)), ","), "[]")
		query = fmt.Sprintf(`SELECT b.id, b.content, b.is_active, b.feature_id, b.created_at, b.updated_at FROM %s b 
							INNER JOIN %s bt ON b.id = bt.banner_id WHERE bt.tags_id IN ($1) %s`, bannersTable, bannerTagsTable, limitOffset)
		if err := r.db.Select(&banners, query, tags); err != nil {
			return nil, err
		}
	}

	if len(ban.Tag) == 0 && ban.Feature != 0 {
		query = fmt.Sprintf(`SELECT * FROM %s WHERE feature_id = $1 %s`, bannersTable, limitOffset)
		if err := r.db.Select(&banners, query, ban.Feature); err != nil {
			return nil, err
		}
	}

	if len(ban.Tag) != 0 && ban.Feature != 0 {
		tags := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ban.Tag)), ","), "[]")
		query = fmt.Sprintf(`SELECT * FROM %s b INNER JOIN %s bt ON b.id = bt.banner_id WHERE bt.tags_id IN ($1) AND b.feature_id = $2 %s`, bannersTable, bannerTagsTable, limitOffset)
		if err := r.db.Select(&banners, query, tags, ban.Feature); err != nil {
			return nil, err
		}
	}

	for i := range banners {
		query = fmt.Sprintf(`SELECT tags_id FROM %s WHERE banner_id = $1`, bannerTagsTable)
		if err := r.db.Select(&banners[i].Tag, query, banners[i].Id); err != nil {
			return nil, err
		}
	}

	return banners, nil
}
