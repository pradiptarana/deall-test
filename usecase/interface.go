package usecase

import "deal-test/model"

type UsersUsecase interface {
	SignUp(req *model.UserSocial) error
	Login(req *model.LoginRequest) (string, error)
}
