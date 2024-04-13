package service

import (
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

type BannerService struct {
	repo repository.Banner
}

func NewBannerSerivce(repo repository.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) Create(banner banner.Banner) (int, error) {
	return s.repo.Create(banner)
}

func (s *BannerService) GetBanner(banner banner.Banner, limit, offset int) ([]banner.Banner, error) {
	return s.repo.GetBanner(banner, limit, offset)
}
