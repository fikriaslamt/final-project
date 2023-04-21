package repository

import (
	"final-project/model"
)

// clean architectures -> handler->service->repo

// interface employee
type UserRepo interface {
	Register(model.User) (model.User, error)
	Login(id int64) error
	FindByEmail(string) (model.User, bool, error)
	FindByUsername(string) (model.User, bool, error)
}

func (r Repo) Register(user model.User) (res model.User, err error) {
	err = r.gorm.Create(&user).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) Login(id int64) error {
	return nil
}

func (r Repo) FindByEmail(email string) (user model.User, res bool, err error) {
	err = r.gorm.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, false, err
	}

	return user, true, nil
}
func (r Repo) FindByUsername(username string) (user model.User, res bool, err error) {
	err = r.gorm.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return model.User{}, false, err
	}

	return user, true, nil
}
