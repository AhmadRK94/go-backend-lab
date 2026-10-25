package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string    `gorm:"not null"`
	HashedPassword string    `gorm:"not null"`
	DateOfBirth    time.Time `gorm:"not null"`
	Email          string    `gorm:"not null;uniqueIndex;size:254"`
}

type RegisterUser struct {
	Username        string    `json:"username" validate:"required,min=2,max=100"`
	Password        string    `json:"password" validate:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" validate:"required,eqfield=Password"`
	Email           string    `json:"email" validate:"required,email"`
	DateOfBirth     time.Time `json:"date_of_birth" validate:"required"`
}

type UpdateUserProfile struct {
	Username    string `json:"username" validate:"required,min=2,max=100"`
	DateOfBirth time.Time
}

type ResponseUser struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

// type LoginUser struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8"`
// }
