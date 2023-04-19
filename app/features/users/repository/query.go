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
	var inputUser = User{}
	hashedPassword, err := helper.HashedPassword(newUser.Password)
	if err != nil {
		log.Println("Hashing password error", err.Error())
		return users.Core{}, err
	}

	inputUser.Username = newUser.Username
	inputUser.Email = newUser.Email
	inputUser.Password = hashedPassword

	if err := ar.db.Table("users").Create(&inputUser).Error; err != nil {
		log.Println("Register error, email "+inputUser.Email+" has been registered", err.Error())
		return users.Core{}, err
	}
	return newUser, nil
}

func (ar *UserRepository) Login(email, password string) (users.Core, error) {
	result := User{}
	if err := ar.db.Where("email = ?", email).Find(&result).Error; err != nil {
		log.Println("Email not found", err.Error())
		return users.Core{}, errors.New("Email not found")
	}

	if err := helper.VerifyPassword(result.Password, password); err != nil {
		log.Println("Invalid password")
		return users.Core{}, errors.New("Invalid password")
	}
	return users.Core{Username: result.Username, Email: result.Email}, nil
}

func (ar *UserRepository) GetProfile(userID int) (users.Core, error) {
	tmp := User{}
	tx := ar.db.Where("id = ?", userID).First(&tmp)
	if tx.RowsAffected < 1 {
		log.Println("Terjadi error saat first user (data tidak ditemukan)")
		return users.Core{}, errors.New("user not found")
	}
	if tx.Error != nil {
		log.Println("Terjadi Kesalahan")
		return users.Core{}, tx.Error
	}
	return users.Core{}, nil
}

func (ar *UserRepository) UpdateProfile(userID uint, updatedUser users.Core) error {
	userInput := User{}
	if err := ar.db.Where("id = ?", userID).Find(&userInput).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %v not found", userID)
		}
		log.Print("Failed to query user by ID", err)
		return err
	}

	userInput.Username = updatedUser.Username
	userInput.Email = updatedUser.Email
	userInput.Password = updatedUser.Password
	userInput.UpdatedAt = time.Now()

	if err := ar.db.Save(&userInput).Error; err != nil {
		log.Print("Failed to update user", err)
		return err
	}
	return nil
}

func (ar *UserRepository) DeleteProfile(userID uint) error {
	user := User{}
	if userID == 0 {
		return fmt.Errorf("Terjadi kesalahan input ID")
	}
	if err := ar.db.Find(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("ID user %v tidak ditemukan", userID)
		}
		log.Println("Terjadi error saat mengambil user dengan ID", err)
		return err
	}

	user.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := ar.db.Save(&user).Error; err != nil {
		log.Println("Terjadi error saat melakukan user buku", err)
		return err
	}
	return nil
}
