package repository

import (
	"final-project/model"
)

// clean architectures -> handler->service->repo

// interface employee
type SocialMediaRepo interface {
	CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error)
	GetAllSocialMedia(id int) (res []model.SocialMedia, err error)
	GetOneSocialMedia(int, int) (res model.SocialMedia, err error)
	CheckSocialMediaById(id int) (bool, error)
	UpdateSocialMedia(userId int, id int, in model.SocialMedia) (res model.SocialMedia, err error)
	DeleteSocialMedia(int, int) (err error)
}

func (r Repo) CreateSocialMedia(in model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllSocialMedia(id int) (res []model.SocialMedia, err error) {
	err = r.gorm.Where("user_id", id).Preload("User", id).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOneSocialMedia(userId int, id int) (res model.SocialMedia, err error) {
	err = r.gorm.Where("user_id", userId).Preload("User", userId).First(&res, id).Error
	if err != nil {

		return res, err
	}

	return res, nil
}

func (r Repo) CheckSocialMediaById(id int) (isAvail bool, err error) {
	err = r.gorm.Where("id = ?", id).First(&model.SocialMedia{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r Repo) UpdateSocialMedia(userId int, id int, in model.SocialMedia) (res model.SocialMedia, err error) {

	err = r.gorm.Model(&res).Where("user_id", userId).Where("id = ?", id).Updates(model.SocialMedia{
		Name:           in.Name,
		SocialMediaUrl: in.SocialMediaUrl,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteSocialMedia(userId int, id int) (err error) {

	err = r.gorm.Debug().Where("user_id = ?", userId).Model(&model.SocialMedia{}).Delete(&model.SocialMedia{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
