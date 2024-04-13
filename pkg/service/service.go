package service

import (
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

type Authorization interface {
	CreateUser(user banner.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Banner interface {
	Create(banner banner.Banner) (int, error)
	GetBanner(banner banner.Banner, limit, offset int) ([]banner.Banner, error)
}

type UserBanner interface {
	GetUserBanner(banner banner.UserBanner, lastRevision bool) (interface{}, error)
}

type Tag interface {
	Create(tag banner.Tag) (int, error)
}

type Feature interface {
	Create(feature banner.Feature) (int, error)
}

type Service struct {
	Authorization
	Banner
	UserBanner
	Tag
	Feature
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Banner:        NewBannerSerivce(repos.Banner),
		UserBanner:    NewUserBannerSerivce(repos.UserBanner),
		Tag:           NewTagSerivce(repos.Tag),
		Feature:       NewFeatureSerivce(repos.Feature),
	}
}
