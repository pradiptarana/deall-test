package repository

import "github.com/pradiptarana/deall-test/model"

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks github.com/pradiptarana/deall-test/repository UserRepository
type UserRepository interface {
	SignUp(us *model.UserSocial) error
	Login(username, password string) (bool, error)
}
