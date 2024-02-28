package service

import (
	"congo"
	"congo/pkg/repository"
)

type AccountsList interface {
	GetAll() ([]congo.Account, error)
	FilterSex(string) ([]congo.Account, error)
}

type Service struct {
	AccountsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AccountsList: NewAccountService(repos.AccountsList),
	}
}
