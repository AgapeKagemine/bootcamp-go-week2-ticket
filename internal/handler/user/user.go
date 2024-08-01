package user

import (
	"context"

	"gotik/internal/domain"
	"gotik/internal/handler/contract"
	"gotik/internal/usecase/user"
)

type UserHandler interface {
	contract.FindAll
	contract.FindById
	contract.Save
	contract.Update
	contract.DeleteById
}

type UserHandlerImpl struct {
	uc user.UserUsecase[context.Context, domain.User]
}

func NewUserHandler(uc user.UserUsecase[context.Context, domain.User]) UserHandler {
	return &UserHandlerImpl{
		uc: uc,
	}
}
