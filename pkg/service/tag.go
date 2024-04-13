package service

import (
	banner "github.com/zanzhit/avito_task"
	"github.com/zanzhit/avito_task/pkg/repository"
)

type TagService struct {
	repo repository.Tag
}

func NewTagSerivce(repo repository.Tag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) Create(tag banner.Tag) (int, error) {
	return s.repo.Create(tag)
}
