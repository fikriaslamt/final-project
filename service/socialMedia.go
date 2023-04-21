package service

import (
	"errors"
	"final-project/model"
)

type SocialMediaService interface {
	GetAllSocialMedia(int) ([]model.SocialMedia, error)
	GetOneSocialMedia(int, int) (model.SocialMedia, error)
	CreateSocialMedia(userId int, in model.SocialMedia) (model.SocialMedia, error)
	UpdateSocialMedia(userId int, id int, in model.SocialMedia) (model.SocialMedia, error)
	DeleteSocialMedia(int, int) error
}

func (s *Service) GetAllSocialMedia(id int) ([]model.SocialMedia, error) {
	return s.repo.GetAllSocialMedia(id)
}

func (s *Service) GetOneSocialMedia(userId int, id int) (model.SocialMedia, error) {
	return s.repo.GetOneSocialMedia(userId, id)
}

func (s *Service) CreateSocialMedia(userId int, in model.SocialMedia) (model.SocialMedia, error) {
	isAvail, _ := s.repo.CheckSocialMediaById(userId)
	if isAvail {
		return model.SocialMedia{}, errors.New("can't create more than one sosmed")
	}
	in.UserID = userId
	return s.repo.CreateSocialMedia(in)
}

func (s *Service) UpdateSocialMedia(userId int, id int, in model.SocialMedia) (model.SocialMedia, error) {
	isAvail, err := s.repo.CheckSocialMediaById(id)
	if !isAvail {
		return model.SocialMedia{}, err
	}
	return s.repo.UpdateSocialMedia(userId, id, in)
}
func (s *Service) DeleteSocialMedia(userId int, id int) error {
	isAvail, err := s.repo.CheckSocialMediaById(id)
	if !isAvail {
		return err
	}
	return s.repo.DeleteSocialMedia(userId, id)
}
