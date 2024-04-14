package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/errs"
)

type BannerPostgres struct {
	db *sqlx.DB
}

func NewBannerPostgres(db *sqlx.DB) *BannerPostgres {
	return &BannerPostgres{db: db}
}

func (r *BannerPostgres) Create(banner banner.Banner) (int, error) {
	alreadyExist, err := r.isBannerUnique(banner)
	if err != nil {
		return 0, err
	}

	if alreadyExist > 0 {
		return -1, nil
	}

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

func (r *BannerPostgres) Patch(ban banner.UpdateBanner, bannerId int) error {
	var updateBanner banner.Banner
	findQuery := fmt.Sprintf("SELECT b.feature_id, bt.tags_id FROM %s b JOIN %s bt ON b.id = bt.banner_id WHERE b.id = $1", bannersTable, bannerTagsTable)

	rows, err := r.db.Query(findQuery, bannerId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tagId int
		err = rows.Scan(&updateBanner.Feature, &tagId)
		if err != nil {
			return err
		}
		updateBanner.Tag = append(updateBanner.Tag, tagId)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	ardId := 1

	if ban.Feature != nil {
		setValues = append(setValues, fmt.Sprintf("feature_id=$%d", ardId))
		args = append(args, *ban.Feature)
		ardId++
		updateBanner.Feature = *ban.Feature
	}

	if ban.Content != nil {
		bannerContent, err := json.Marshal(*ban.Content)
		if err != nil {
			return err
		}

		setValues = append(setValues, fmt.Sprintf("content=$%d", ardId))
		args = append(args, bannerContent)
		ardId++
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if len(setValues) > 0 {
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", bannersTable, setQuery, bannerId)
		_, err = r.db.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if ban.Tag != nil {
		copy(updateBanner.Tag, *ban.Tag)
		deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE banner_id = $1", bannerTagsTable)
		_, err := r.db.Exec(deleteQuery, bannerId)
		if err != nil {
			tx.Rollback()
			return err
		}

		for _, tagId := range *ban.Tag {
			insertQuery := fmt.Sprintf("INSERT INTO %s (banner_id, tags_id) VALUES ($1, $2)", bannerTagsTable)
			_, err := r.db.Query(insertQuery, bannerId, tagId)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	isUnique, err := r.isBannerUnique(updateBanner)
	if err != nil {
		tx.Rollback()
		return err
	}

	if isUnique > 1 {
		tx.Rollback()
		return &errs.ErrBannerNotUnique{}
	}

	return tx.Commit()
}

func (r *BannerPostgres) Delete(bannerId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", bannersTable)
	deletedBanner, err := r.db.Exec(query, bannerId)
	if err != nil {
		return err
	}

	count, err := deletedBanner.RowsAffected()
	if err != nil {
		return err
	}

	if count != 0 {
		return &errs.ErrBannerNotFound{}
	}

	return nil
}

func (r *BannerPostgres) isBannerUnique(ban banner.Banner) (int, error) {
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s b
						  INNER JOIN %s bt ON b.id = bt.banner_id
						  WHERE b.feature_id = $1 AND bt.tags_id = ANY($2)`, bannersTable, bannerTagsTable)

	var count int
	err := r.db.QueryRow(query, ban.Feature, pq.Array(ban.Tag)).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
