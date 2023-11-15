package service

import (
	"github.com/afikrim/go-hexa-template/core/module"
	"github.com/afikrim/go-hexa-template/core/repository"
)

type (
	base struct {
		repo repository.BaseRepository
	}
)

func New(repo repository.BaseRepository) module.BaseModule {
	return &base{
		repo: repo,
	}
}
