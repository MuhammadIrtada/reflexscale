package usecase

import (
	"errors"
	"reflexscale/sdk/jwt"
	"reflexscale/src/entitiy"
	"reflexscale/src/repository"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserInterface interface {
		Login(userInput *entity.UserLoginInputParam) (*entity.User, string, error)
		Register(userInput *entity.UserRegisterInputParam) (*entity.User, error)
		ReadAll() ([]*entity.User, error)
		ReadByID(ID int) (*entity.User, error)
		Update(ID int, userRequest *entity.UserUpdateInputParam) (*entity.User, error)
		Delete(ID int) (*entity.User, error)
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

func (u *User) Register(userInput *entity.UserRegisterInputParam) (*entity.User, error) {
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

func (u *User) Login(userInput *entity.UserLoginInputParam) (*entity.User, string, error) {
	user, err := u.userRepo.FindByEmail(userInput.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	tokenJwt, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("error generating token")
	}

	return user, tokenJwt, nil
}

func (u *User) ReadAll() ([]*entity.User, error) {
	users, err := u.userRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) ReadByID(ID int) (*entity.User, error) {
	user, err := u.userRepo.ReadByID(ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetByEmail(email string) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Update(ID int, userRequest *entity.UserUpdateInputParam) (*entity.User, error) {
	newUser, err := u.userRepo.Update(&entity.User{
		ID:           uint(ID),
		NamaLengkap:  userRequest.NamaLengkap,
		Email:        userRequest.Email,
		Usia:         userRequest.Usia,
		JenisKelamin: userRequest.JenisKelamin,
		Alamat:       userRequest.Alamat,
	})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *User) Delete(ID int) (*entity.User, error) {
	user, err := u.userRepo.ReadByID(ID)
	if err != nil {
		return nil, err
	}

	user, err = u.userRepo.Delete(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
