package service

import "congo/pkg/repository"

type AccountsList interface {
}

type Service struct {
	AccountsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
