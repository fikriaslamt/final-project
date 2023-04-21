package service

import "final-project/model"

type CommentService interface {
	GetAllComment(userId int, photoId int) ([]model.Comment, error)
	GetOneComment(int, int) (model.Comment, error)
	CreateComment(userId int, photoId int, in model.Comment) (model.Comment, error)
	UpdateComment(userId int, id int, in model.Comment) (model.Comment, error)
	DeleteComment(userId int, id int) error
}

func (s *Service) GetAllComment(userId int, photoId int) ([]model.Comment, error) {
	return s.repo.GetAllComment(userId, photoId)
}

func (s *Service) GetOneComment(userId int, id int) (model.Comment, error) {
	return s.repo.GetOneComment(userId, id)
}

func (s *Service) CreateComment(userId int, photoId int, in model.Comment) (model.Comment, error) {
	in.UserID = userId
	in.PhotoID = photoId
	return s.repo.CreateComment(in)
}

func (s *Service) UpdateComment(userId int, id int, in model.Comment) (model.Comment, error) {
	data, err := s.repo.GetOneComment(userId, id)
	if err != nil {
		return model.Comment{}, err
	}
	photoId := data.PhotoID
	isAvail, err := s.repo.CheckCommentById(userId, photoId, id)
	if !isAvail {
		return model.Comment{}, err
	}
	return s.repo.UpdateComment(userId, photoId, id, in)
}
func (s *Service) DeleteComment(userId int, id int) error {
	data, err := s.repo.GetOneComment(userId, id)
	if err != nil {
		return err
	}
	photoId := data.PhotoID
	isAvail, err := s.repo.CheckCommentById(userId, photoId, id)
	if !isAvail {
		return err
	}
	return s.repo.DeleteComment(userId, photoId, id)

}
