package repository

import "deal-test/model"

//go:generate mockgen -destination=../mocks/mock_user.go -package=mocks deal-test/repository UserRepository
type UserRepository interface {
	SignUp(us *model.UserSocial) error
	Login(username, password string) (bool, error)
}
