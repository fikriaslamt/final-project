package repository

import (
	"final-project/model"
)

// clean architectures -> handler->service->repo

// interface employee
type PhotoRepo interface {
	CreatePhoto(model.Photo) (res model.Photo, err error)
	GetAllPhoto(int) (res []model.Photo, err error)
	GetOnePhoto(int, int) (res model.Photo, err error)
	CheckPhotoById(id int) (bool, error)
	UpdatePhoto(int, int, model.Photo) (res model.Photo, err error)
	DeletePhoto(int, int) (err error)
}

func (r Repo) CreatePhoto(in model.Photo) (res model.Photo, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllPhoto(id int) (res []model.Photo, err error) {
	err = r.gorm.Where("user_id", id).Preload("User", id).Preload("Comments").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOnePhoto(userId int, id int) (res model.Photo, err error) {
	err = r.gorm.Where("user_id", userId).Preload("User", userId).First(&res, id).Error
	if err != nil {

		return res, err
	}

	return res, nil
}

func (r Repo) CheckPhotoById(id int) (isAvail bool, err error) {
	err = r.gorm.First(&model.Photo{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r Repo) UpdatePhoto(userId int, id int, in model.Photo) (res model.Photo, err error) {

	err = r.gorm.Model(&res).Where("user_id", id).Preload("User", userId).Where("id = ?", id).Updates(model.Photo{
		Title:    in.Title,
		Caption:  in.Caption,
		PhotoUrl: in.PhotoUrl,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeletePhoto(userId int, id int) (err error) {

	err = r.gorm.Debug().Where("user_id = ?", userId).Model(&model.Photo{}).Delete(&model.Photo{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
