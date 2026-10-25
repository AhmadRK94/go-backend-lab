package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id uint) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *User) error
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

func (r *GormUserRepository) Create(ctx context.Context, user *User) error {
	return nil
}

func (r *GormUserRepository) GetUserByID(ctx context.Context, id uint) (User, error) {

	return User{}, nil
}
func (r *GormUserRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return User{}, nil
}

func (r *GormUserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	return nil, nil
}

func (r *GormUserRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (r *GormUserRepository) Update(ctx context.Context, user *User) error {
	return nil
}

func (r *GormUserRepository) ChangePassword(ctx context.Context, id uint, hash string) error {
	return nil
}
