package repository

import (
	"Event-Planning-App/app/features/users"
	"Event-Planning-App/helper"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ar *UserRepository) Register(newUser users.Core) (users.Core, error) {
	var input = User{}
	hashedPassword, err := helper.HashedPassword(newUser.Password)
	if err != nil {
		log.Println("Hashing password error", err.Error())
		return users.Core{}, err
	}

	input.Username = newUser.Username
	input.Email = newUser.Email
	input.Password = hashedPassword

	if err := ar.db.Table("users").Create(&input).Error; err != nil {
		log.Println("Register error, email "+input.Email+" has been registered", err.Error())
		return users.Core{}, err
	}
	return newUser, nil
}

func (ar *UserRepository) Login(email, password string) (users.Core, error) {
	var input User
	if err := ar.db.Where("email = ?", email).Find(&input).Error; err != nil {
		return users.Core{}, errors.New("Email not found")
	}

	if err := helper.VerifyPassword(input.Password, password); err != nil {
		return users.Core{}, errors.New("Invalid password")
	}

	return users.Core{Email: input.Email}, nil
}

func (ar *UserRepository) GetProfile(email string) (users.Core, error) {
	var input User
	result := ar.db.Where("email = ?", email).Find(&input)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return users.Core{}, errors.New("user not found")
	}
	return users.Core{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}, nil
}

func (ar *UserRepository) UpdateProfile(email string, updatedUser users.Core) error {
	input := User{}
	if err := ar.db.Where("email = ?", email).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with email %v not found", email)
		}
		log.Print("Failed to query user by email", err)
		return err
	}

	input.Username = updatedUser.Username
	input.Email = updatedUser.Email
	input.Password = updatedUser.Password
	input.UpdatedAt = time.Now()

	if err := ar.db.Save(&input).Error; err != nil {
		log.Print("Failed to update user", err)
		return err
	}
	return nil
}

func (ar *UserRepository) DeleteProfile(email string) error {
	input := User{}
	if err := ar.db.Where("email = ?", email).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with email %v not found", email)
		}
		log.Print("Failed to query user by email", err)
		return err
	}

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := ar.db.Save(&input).Error; err != nil {
		log.Println("Terjadi error saat melakukan user buku", err)
		return err
	}
	return nil
}
