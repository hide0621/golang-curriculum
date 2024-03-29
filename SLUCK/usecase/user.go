package usecase

import (
	"context"
	"sluck/model"
	"sluck/repository"
	"sluck/transaction"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id int) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
}

type userUsecase struct {
	r           repository.UserRepository
	mr          repository.MessageRepository
	transaction transaction.Transaction
}

func NewUserUsecase(r repository.UserRepository, mr repository.MessageRepository, transaction transaction.Transaction) UserUsecase {
	return &userUsecase{
		r:           r,
		mr:          mr,
		transaction: transaction,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {

	user, err := u.r.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {

	id, err := u.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil

}

func (u *userUsecase) Update(ctx context.Context, user *model.User) error {

	err := u.r.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil

}

func (u *userUsecase) Delete(ctx context.Context, id int) error {

	// トランザクションを開始して、そこでDB操作を行う
	u.transaction.DoInTx(ctx, func(ctx context.Context) (any, error) {

		// 以下の2つのメソッドがDoInTxメソッドの第二引数に相当する(2つのメソッドの結果をDoInTxメソッドの第二引数に渡している)

		err := u.r.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		err = u.mr.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		return nil, nil

	})

	return nil

}
