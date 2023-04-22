package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users/mock"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users/service"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userService := service.New(mockRepo)

	// Test successful registration
	mockUser := users.Core{Username: "Test User"}
	mockRepo.EXPECT().Register(mockUser).Return(nil, errors.New("Failed to register user"))

	err := userService.Register(mockUser)
	assert.NoError(t, err)

	// Test failed registration
	mockRepo.EXPECT().Register(mockUser).Return(nil, errors.New("Failed to register user"))
	err = userService.Register(mockUser)
	assert.Error(t, err)
	assert.Equal(t, "Failed to register user", err.Error())
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userService := service.New(mockRepo)

	// Test successful login
	email := "test@test.com"
	password := "password"
	mockUser := users.Core{Email: email, Password: password}
	mockRepo.EXPECT().Login(email, password).Return(mockUser, nil)
	user, err := userService.Login(email, password)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// Test failed login
	mockRepo.EXPECT().Login(email, password).Return(users.Core{}, errors.New("Invalid email or password"))
	user, err = userService.Login(email, password)
	assert.Error(t, err)
	assert.Equal(t, users.Core{}, user)
	assert.Equal(t, "Invalid email or password", err.Error())
}

func TestUserService_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userService := service.New(mockRepo)

	// Test successful get profile
	email := "test@test.com"
	mockUser := users.Core{Email: email}
	mockRepo.EXPECT().GetProfile(email).Return(mockUser, nil)
	user, err := userService.GetProfile(email)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	// Test failed get profile
	mockRepo.EXPECT().GetProfile(email).Return(users.Core{}, errors.New("Failed to get user profile"))
	user, err = userService.GetProfile(email)
	assert.Error(t, err)
	assert.Equal(t, users.Core{}, user)
	assert.Equal(t, "Failed to get user profile", err.Error())
}

func HashedPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func VerifyPassword(passhash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func TestUserService_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)
	userService := service.New(mockRepo)

	// Test successful update profile
	email := "test@test.com"
	username := "newusername"
	newEmail := "newemail@test.com"
	password := "newpassword"
	image := "newimage"
	hashedPassword := []byte("hashedpassword")
	mockRepo.EXPECT().UpdateProfile(email, users.Core{
		Username: username,
		Email:    newEmail,
		Password: string(hashedPassword),
		Image:    image,
	}).Return(nil)
	err := userService.UpdateProfile(email, username, newEmail, password, image)
	assert.NoError(t, err)

	// Test failed update profile
	mockRepo.EXPECT().UpdateProfile(email, users.Core{
		Username: username,
		Email:    newEmail,
		Password: password,
		Image:    image,
	}).Return(errors.New("Failed to update user"))
	err = userService.UpdateProfile(email, username, newEmail, password, image)
	assert.Error(t, err)
	assert.Equal(t, "Error while updating test@test.com: Failed to update user", err.Error())
}

func TestUserService_DeleteProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	userService := service.New(mockRepo)

	// Test successful delete profile
	email := "test@test.com"
	mockRepo.EXPECT().DeleteProfile(email).Return(nil)
	err := userService.DeleteProfile(email)
	assert.NoError(t, err)

	// Test failed delete profile
	mockRepo.EXPECT().DeleteProfile(email).Return(errors.New("Failed to delete user"))
	err = userService.DeleteProfile(email)
	assert.Error(t, err)
	assert.Equal(t, "Terjadi masalah pada server", err.Error())
}
