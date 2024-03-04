package service

import (
	"congo"
	"congo/pkg/repository"
)

type AccountsList interface {
	GetAll() ([]congo.Account, error)
	Filter([]congo.Filter, int) ([]congo.Account, error)
}

type Service struct {
	AccountsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AccountsList: NewAccountService(repos.AccountsList),
	}
}
