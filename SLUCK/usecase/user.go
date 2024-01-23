package usecase

import (
	"context"
	"fmt"
	"sluck/model"
	"sluck/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, user *model.User) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r: r}
}

func (u *userUsecase) Create(ctx context.Context, user *model.User) error {

	fmt.Println("usecase creating ... ")

	return nil
}
