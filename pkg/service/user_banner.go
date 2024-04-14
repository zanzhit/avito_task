package service

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

type UserBannerService struct {
	repo repository.UserBanner
}

func NewUserBannerSerivce(repo repository.UserBanner) *UserBannerService {
	return &UserBannerService{repo: repo}
}

func (s *UserBannerService) GetUserBanner(banner banner.UserBanner, lastRevision bool) (interface{}, error) {
	if lastRevision {
		return s.repo.GetUserBanner(banner)
	}

	key := fmt.Sprintf("%d,%d", banner.Tag, banner.Feature)
	if cacheBanner, found := c.Get(key); found {
		return cacheBanner, nil
	}

	userBanner, err := s.repo.GetUserBanner(banner)
	if err != nil {
		return nil, err
	}

	c.Set(key, userBanner, cache.DefaultExpiration)
	return userBanner, nil
}
