package service

import (
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

type FeatureService struct {
	repo repository.Feature
}

func NewFeatureSerivce(repo repository.Feature) *FeatureService {
	return &FeatureService{repo: repo}
}

func (s *FeatureService) Create(feature banner.Feature) (int, error) {
	return s.repo.Create(feature)
}
