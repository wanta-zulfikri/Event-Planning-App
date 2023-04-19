package service

import (
	"Event-Planning-App/app/features/users"
	"Event-Planning-App/helper"
	"errors"
	"fmt"
	"log"
	"strings"

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
	result, err := us.m.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "Email not found") {
			return users.Core{}, errors.New("Email not found")
		} else if strings.Contains(err.Error(), "Invalid password") {
			return users.Core{}, errors.New("Invalid password")
		}
		return users.Core{}, errors.New("Failed to login user")
	}
	return result, nil
}

func (us *UserService) GetProfile(userID int) (users.Core, error) {
	tmp, err := us.m.GetProfile(userID)
	if err != nil {
		return users.Core{}, err
	}
	return tmp, nil
}

func (us *UserService) UpdateProfile(id uint, username, email, password string) error {
	hashedPassword, err := helper.HashedPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	updatedUser := users.Core{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	if err := us.m.UpdateProfile(id, updatedUser); err != nil {
		return fmt.Errorf("failed to update profile: %v", err)
	}
	return nil
}

func (us *UserService) DeleteProfile(userID uint) error {
	if userID == 0 {
		return fmt.Errorf("ID buku tidak valid")
	}
	err := us.m.DeleteProfile(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user dengan ID %v tidak ditemukan", userID)
		}
		log.Printf("terjadi kesalahan saat menghapus data user dengan ID %d: %v", userID, err)
		return errors.New("terdapat masalah pada server")
	}
	return nil
}
