package service

import (
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

type UserBannerService struct {
	repo repository.UserBanner
}

func NewUserBannerSerivce(repo repository.UserBanner) *UserBannerService {
	return &UserBannerService{repo: repo}
}

func (s *UserBannerService) GetUserBanner(banner banner.UserBanner, lastRevision bool) (interface{}, error) {
	return s.repo.GetUserBanner(banner, lastRevision)
}
