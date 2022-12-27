package models

import "time"

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
	Password string `json:"password" validate:"required,min=8" gorm:"unique"`
}

type UserBase struct {
	UserName string `json:"username" gorm:"unique" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email" gorm:"unique"`
}

type UserRequest struct {
	UserBase
	Password             string `json:"password" validate:"required,min=8" gorm:"unique"`
	PasswordConfirmation string `json:"password-confirmation" validate:"required,min=8" gorm:"unique"`
}

type UserResponse struct {
	Id int `json:"id" gorm:"PrimaryKey"`
	UserBase
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	UserResponse
	Token    string  `json:"token" validate:"jwt"`
	IsActive bool    `json:"is-active"  validate:"boolean"`
	Password string  `json:"password" validate:"required,min=8" gorm:"unique"`
	Foods    []Food  `json:"foods" gorm:"foreignKey:UserId"`
	Images   []Image `json:"images" gorm:"foreignKey:UserId"`
}

func ParseToUser(u *User, ur UserRequest) {
	u.UserName = ur.UserName
	u.Email = ur.Email
	u.Password = ur.Password
}

func ParseToUserResponse(u User, ur *UserResponse) {
	ur.Id = u.Id
	ur.UserName = u.UserName
	ur.Email = u.Email
	ur.CreatedAt = u.CreatedAt
	ur.UpdatedAt = u.UpdatedAt
}
