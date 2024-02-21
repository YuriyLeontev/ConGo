package service

import (
	"congo"
	"congo/pkg/repository"
)

type AccountService struct {
	repo repository.AccountsList
}

func NewAccountService(repo repository.AccountsList) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetAll() ([]congo.Account, error) {
	return s.repo.GetAll()
}
