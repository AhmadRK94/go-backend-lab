package user

import "context"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user RegisterUser) (*ResponseUser, error) {
	return nil, nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*ResponseUser, error) {
	return nil, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*ResponseUser, error) {
	return nil, nil
}

func (s *UserService) DeleteAccount(ctx context.Context, id uint) error {
	return nil
}

func (s *UserService) UpdateProfile(ctx context.Context, id uint, newInfo UpdateUserProfile) (*ResponseUser, error) {
	return nil, nil
}

func (s *UserService) ChangePassword(ctx context.Context, id uint, currentPassword string, newPassword string) error {
	return nil
}
