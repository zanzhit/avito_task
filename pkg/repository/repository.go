package repository

import (
	"github.com/jmoiron/sqlx"
	banner "github.com/zanzhit/avito_task"
)

type Authorization interface {
	CreateUser(user banner.User) (int, error)
	GetUser(username, password string) (banner.User, error)
}

type Banner interface {
	Create(banner banner.Banner) (int, error)
	GetBanner(banner banner.Banner, limit, offset int) ([]banner.Banner, error)
	Patch(banner banner.UpdateBanner, bannerId int) error
	Delete(bannerId int) error
}

type UserBanner interface {
	GetUserBanner(banner banner.UserBanner) (interface{}, error)
}

type Tag interface {
	Create(tag banner.Tag) (int, error)
}

type Feature interface {
	Create(feature banner.Feature) (int, error)
}

type Repository struct {
	Authorization
	Banner
	UserBanner
	Tag
	Feature
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Banner:        NewBannerPostgres(db),
		UserBanner:    NewUserBannerPostgres(db),
		Tag:           NewTagPostgres(db),
		Feature:       NewFeaturePostgres(db),
	}
}
