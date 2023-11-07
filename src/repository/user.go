package repository

import (
	"reflexscale/database/mysql"
	entity "reflexscale/src/entitiy"
)

type (
	UserInterface interface {
		Create(user *entity.User) (*entity.User, error)
		ReadAll() ([]*entity.User, error)
		ReadByID(ID int) (*entity.User, error)
		FindByEmail(email string) (*entity.User, error)
		Update(user *entity.User) (*entity.User, error)
		Delete(user *entity.User) (*entity.User, error)
	}

	user struct {
		db mysql.DB
	}
)

func InitUser(db *mysql.DB) UserInterface {
	return &user{
		db: *db,
	}
}

func (u *user) Create(user *entity.User) (*entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *user) ReadAll() ([]*entity.User, error) {
	var user []*entity.User

	err := u.db.Preload("HasilTests").Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *user) ReadByID(ID int) (*entity.User, error) {
	var user *entity.User

	err := u.db.Preload("HasilTests").First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *user) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User

	err := u.db.Preload("HasilTests").First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *user) Update(user *entity.User) (*entity.User, error) {
	err := u.db.Model(&user).Preload("HasilTests").Updates(&user).First(&user).Error
	if err != nil {
		return nil, err

	}

	return user, err
}

func (u *user) Delete(user *entity.User) (*entity.User, error) {
	err := u.db.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return user, err
}
