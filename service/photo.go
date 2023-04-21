package service

import "final-project/model"

type PhotoService interface {
	GetAllPhoto(int) ([]model.Photo, error)
	GetOnePhoto(int, int) (model.Photo, error)
	CreatePhoto(int, model.Photo) (model.Photo, error)
	UpdatePhoto(int, int, model.Photo) (model.Photo, error)
	DeletePhoto(int, int) error
}

func (s *Service) GetAllPhoto(id int) ([]model.Photo, error) {
	return s.repo.GetAllPhoto(id)
}

func (s *Service) GetOnePhoto(userId int, id int) (model.Photo, error) {
	return s.repo.GetOnePhoto(userId, id)
}

func (s *Service) CreatePhoto(userId int, in model.Photo) (model.Photo, error) {
	in.UserID = userId
	return s.repo.CreatePhoto(in)
}

func (s *Service) UpdatePhoto(userId int, id int, in model.Photo) (model.Photo, error) {
	isAvail, err := s.repo.CheckPhotoById(id)
	if !isAvail {
		return model.Photo{}, err
	}
	return s.repo.UpdatePhoto(userId, id, in)

}
func (s *Service) DeletePhoto(userId int, id int) error {
	isAvail, err := s.repo.CheckPhotoById(id)
	if !isAvail {
		return err
	}
	return s.repo.DeletePhoto(userId, id)

}
