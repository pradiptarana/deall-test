package users

import (
	"deal-test/internal/auth"
	"deal-test/model"
	"deal-test/repository"
)

type UsersUC struct {
	repository.UserRepository
}

func NewUserUC(repo repository.UserRepository) *UsersUC {
	return &UsersUC{repo}
}

func (uc *UsersUC) SignUp(req *model.UserSocial) error {
	// Sign up a user
	err := uc.UserRepository.SignUp(req)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UsersUC) Login(req *model.LoginRequest) (string, error) {
	// Sign up a user
	success, err := uc.UserRepository.Login(req.Username, req.Password)
	if err != nil || !success {
		return "", err
	}

	// Generate JWT for the user
	token, err := auth.GenerateJWT(req.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
