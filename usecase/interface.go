package usecase

import "github.com/pradiptarana/deall-test/model"

type UsersUsecase interface {
	SignUp(req *model.UserSocial) error
	Login(req *model.LoginRequest) (string, error)
}
