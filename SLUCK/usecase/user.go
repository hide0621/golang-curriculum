package usecase

import (
	"context"
	"sluck/repository"
)

type UserUsecase interface {
	Create(ctx context.Context) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r: r}
}

func (u *userUsecase) Create(ctx context.Context) error {
	return nil
}
