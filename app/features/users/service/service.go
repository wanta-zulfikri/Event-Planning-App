package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"

	"gorm.io/gorm"
)

type UserService struct {
	m users.Repository
}

func New(r users.Repository) users.Service {
	return &UserService{m: r}
}

func (us *UserService) Register(newUser users.Core) error {
	_, err := us.m.Register(newUser)
	if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

func (us *UserService) Login(email string, password string) (users.Core, error) {
	user, err := us.m.Login(email, password)
	if err != nil {
		return users.Core{}, err
	}
	return user, nil
}

func (us *UserService) GetProfile(email string) (users.Core, error) {
	user, err := us.m.GetProfile(email)
	if err != nil {
		return users.Core{}, err
	}
	return user, nil
}

func (us *UserService) UpdateProfile(email, username, newEmail, password, image string) error {
	hashedPassword, err := helper.HashedPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	updatedUser := users.Core{
		Username: username,
		Email:    newEmail,
		Password: string(hashedPassword),
		Image:    image,
	}
	if err := us.m.UpdateProfile(email, updatedUser); err != nil {
		return fmt.Errorf("Error while updating %s: %v", email, err)
	}
	return nil
}

func (us *UserService) DeleteProfile(email string) error {
	if email == "" {
		return fmt.Errorf("Invalid email")
	}
	err := us.m.DeleteProfile(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("Email: %v ,not found", email)
		}
		log.Printf("Error while deleting %s: %v", email, err)
		return fmt.Errorf("Terjadi masalah pada server")
	}
	return nil
}
