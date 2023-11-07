package usecase

import (
	"errors"
	entity "reflexscale/src/entitiy"
	"reflexscale/src/repository"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserInterface interface {
		// Login(userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error)
		Register(userInput *entity.UserRegisterInputParam) (*entity.User, error)
		ReadAll() ([]*entity.User, error)
	}

	User struct {
		userRepo repository.UserInterface
	}
)

func InitUser(ur repository.UserInterface) UserInterface {
	return &User{
		userRepo: ur,
	}
}

func (u *User) Register(userInput *entity.UserRegisterInputParam) (*entity.User, error)  {
	if len(userInput.Password) < 8 {
		return nil, errors.New("password minimal 8 karakter")
	}

	if userInput.VerifPassword != userInput.Password {
		return nil, errors.New("konfirmasi password tidak sesuai")
	}

	hash, _ := HashPassword(userInput.Password)

	user := &entity.User{
		NamaLengkap:  userInput.NamaLengkap,
		Email:        userInput.Email,
		Usia:         userInput.Usia,
		JenisKelamin: userInput.JenisKelamin,
		Alamat:       userInput.Alamat,
		Password:     hash,
	}

	user, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) ReadAll() ([]*entity.User, error) {
	users, err := u.userRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}