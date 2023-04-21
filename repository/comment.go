package repository

import (
	"final-project/model"
)

type CommentRepo interface {
	CreateComment(in model.Comment) (res model.Comment, err error)
	GetAllComment(int, int) (res []model.Comment, err error)
	GetOneComment(int, int) (res model.Comment, err error)
	CheckCommentById(userId int, photoId int, id int) (bool, error)
	UpdateComment(userId int, photoId int, id int, in model.Comment) (res model.Comment, err error)
	DeleteComment(userId int, photoId int, id int) (err error)
}

func (r Repo) CreateComment(in model.Comment) (res model.Comment, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllComment(userId int, photoId int) (res []model.Comment, err error) {
	err = r.gorm.Preload("User", userId).Where("user_id  = ?", userId).Where("photo_id  = ?", photoId).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetOneComment(userId int, id int) (res model.Comment, err error) {
	err = r.gorm.Preload("User", userId).First(&res, id, userId).Error
	if err != nil {

		return res, err
	}

	return res, nil
}

func (r Repo) CheckCommentById(userId int, photoId int, id int) (isAvail bool, err error) {
	err = r.gorm.Where("id = ?", id).Where("user_id = ?", userId).Where("photo_id = ?", photoId).First(&model.Comment{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r Repo) UpdateComment(userId int, photoId int, id int, in model.Comment) (res model.Comment, err error) {

	err = r.gorm.Model(&res).Where("id = ?", id).Where("user_id = ?", userId).Where("photo_id = ?", photoId).Updates(model.Comment{
		Message: in.Message,
	}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteComment(userId int, photoId int, id int) (err error) {

	err = r.gorm.Debug().Model(&model.Comment{}).Where("id = ?", id).Where("user_id = ?", userId).Where("photo_id = ?", photoId).Delete(&model.Comment{}).Error
	if err != nil {
		return err
	}

	return nil
}
