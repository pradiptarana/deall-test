package users_test

import (
	"errors"
	"testing"

	"github.com/pradiptarana/deall-test/mocks"
	"github.com/pradiptarana/deall-test/model"

	"github.com/golang/mock/gomock"

	usersUC "github.com/pradiptarana/deall-test/usecase/users"
)

func TestLoginFailed(t *testing.T) {
	req := &model.LoginRequest{
		Username: "user",
		Password: "pass",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().Login(req.Username, req.Password).Return(false, errors.New("user not found")).Times(1)

	_, err := userUC.Login(req)
	if err == nil {
		t.Fail()
	}
}

func TestLoginSuccess(t *testing.T) {
	req := &model.LoginRequest{
		Username: "user",
		Password: "pass",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().Login(req.Username, req.Password).Return(true, nil).Times(1)

	_, err := userUC.Login(req)
	if err != nil {
		t.Fail()
	}
}

func TestSignUpSuccess(t *testing.T) {
	req := &model.UserSocial{
		Username: "user",
		Password: "pass",
		Fullname: "rana pradipta",
		Email:    "ranapradipta4@gmail.com",
		Age:      25,
		Location: "jakarta",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().SignUp(req).Return(nil).Times(1)

	err := userUC.SignUp(req)
	if err != nil {
		t.Fail()
	}
}

func TestSignUpFailed(t *testing.T) {
	req := &model.UserSocial{
		Username: "user",
		Password: "pass",
		Fullname: "rana pradipta",
		Email:    "ranapradipta4@gmail.com",
		Age:      25,
		Location: "jakarta",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().SignUp(req).Return(errors.New("user already exist")).Times(1)

	err := userUC.SignUp(req)
	if err == nil {
		t.Fail()
	}
}
