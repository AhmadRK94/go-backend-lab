package user

import (
	"context"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/auth"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user RegisterUser) (ResponseUser, error) {
	hashedPass, err := auth.HashPassword(user.Password)
	if err != nil {
		return ResponseUser{}, err
	}
	u := &User{
		Username:       user.Username,
		HashedPassword: hashedPass,
		Email:          user.Email,
		DateOfBirth:    user.DateOfBirth,
	}
	err = s.repo.Create(ctx, u)
	if err != nil {
		return ResponseUser{}, err
	}
	return ResponseUser{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		DateOfBirth: u.DateOfBirth,
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (ResponseUser, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return ResponseUser{}, err
	}

	return ResponseUser{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
	}, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (ResponseUser, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return ResponseUser{}, err
	}

	return ResponseUser{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
	}, nil
}

// to do: add pagination
func (s *UserService) GetAll(ctx context.Context) ([]ResponseUser, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	resp := make([]ResponseUser, len(users))
	for i, user := range users {
		resp[i] = ResponseUser{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			DateOfBirth: user.DateOfBirth,
		}
	}
	return resp, nil
}

func (s *UserService) DeleteAccount(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) UpdateProfile(ctx context.Context, id uint, newInfo UpdateUserProfile) (ResponseUser, error) {
	user, err := s.repo.Update(ctx, id, &newInfo)
	if err != nil {
		return ResponseUser{}, err
	}
	return ResponseUser{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
	}, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uint, currentPassword string, newPassword string) error {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	if err := auth.ComparePassword(currentPassword, user.HashedPassword); err != nil {
		return ErrInvalidCurrentPassword
	}
	newHash, err := auth.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return s.repo.ChangePassword(ctx, user.ID, newHash)
}
