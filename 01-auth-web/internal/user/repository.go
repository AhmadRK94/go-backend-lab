package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uint) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *UpdateUserProfile) error
	ChangePassword(ctx context.Context, id uint, hash string) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (r *GormUserRepository) Create(ctx context.Context, createInfo *User) error {
	user := User{
		Username:       createInfo.Username,
		HashedPassword: createInfo.HashedPassword,
		DateOfBirth:    createInfo.DateOfBirth,
		Email:          createInfo.Email,
	}
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (r *GormUserRepository) GetUserByID(ctx context.Context, id uint) (User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}
func (r *GormUserRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email=?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}

func (r *GormUserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (r *GormUserRepository) Update(ctx context.Context, id uint, updateInfo *UpdateUserProfile) error {
	var user User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	user.Username = updateInfo.Username
	user.DateOfBirth = updateInfo.DateOfBirth
	user.Email = updateInfo.Email
	err = r.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrDuplicateEmail
		}
		return err
	}
	return nil
}

func (r *GormUserRepository) ChangePassword(ctx context.Context, id uint, hash string) error {
	var user User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	user.HashedPassword = hash
	return r.db.WithContext(ctx).Save(&user).Error
}
