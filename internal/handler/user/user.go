package user

import (
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
	uc user.UserUsecase
}

func NewUserHandler(uc user.UserUsecase) UserHandler {
	return UserHandlerImpl{
		uc: uc,
	}
}
