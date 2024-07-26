package user

import (
	"gotik/internal/usecase/user"
)

type UserHandler interface {
	user.UserUsecase
}

type UserHandlerImpl struct {
	uc user.UserUsecase
}

func NewUserHandler(uc user.UserUsecase) UserHandler {
	return UserHandlerImpl{
		uc: uc,
	}
}
