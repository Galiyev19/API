package service

import "API/pkg/repository"

type Authorization interface {
}

type Users interface {
}

type Service struct {
	Authorization
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
