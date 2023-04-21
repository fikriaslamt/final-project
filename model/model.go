package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type GormModel struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	GormModel
	UserName    string `json:"user_name" gorm:"type:varchar(100);unique_index"`
	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	Photos      []Photo
	Comments    []Comment
	SocialMedia *SocialMedia
}

type Photo struct {
	GormModel
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    int    `json:"user_id"`
	User      *User
	Comments  []Comment      `json:"comments"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Comment struct {
	GormModel
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
	Photo   *Photo
	User    *User
}

type SocialMedia struct {
	GormModel
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID         int    `json:"user_id"`
	User           *User
}

func (l LoginCredentials) Validation() error { // custom validation
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)

}

func (u User) Validation() error { // custom validation
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&u.Age, validation.Required, validation.Min(9)),
	)

}

func (p Photo) Validation() error { // custom validation
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required),
		validation.Field(&p.PhotoUrl, validation.Required),
	)

}

func (s SocialMedia) Validation() error { // custom validation
	return validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.SocialMediaUrl, validation.Required),
	)

}
func (c Comment) Validation() error { // custom validation
	return validation.ValidateStruct(&c,
		validation.Field(&c.Message, validation.Required),
	)

}
