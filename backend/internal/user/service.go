package user

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	existingUser, err := s.repo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:         common.GenerateID(),
		Username:   req.Username,
		Password:   string(hashedPassword),
		Email:      req.Email,
		SystemRole: req.SystemRole,
		Status:     "active",
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	if user.SystemRole == "" {
		user.SystemRole = "user"
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(userID string) (*model.User, error) {
	return s.repo.GetByID(userID)
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.GetByUsername(username)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(userID string, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.SystemRole != "" {
		user.SystemRole = req.SystemRole
	}

	if req.Status != "" {
		user.Status = req.Status
	}

	user.UpdatedAt = time.Now().Format(time.RFC3339)

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(userID string) error {
	return s.repo.Delete(userID)
}

func (s *UserService) IsSuperAdmin(userID string) (bool, error) {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("user not found")
	}

	return user.SystemRole == "super_admin", nil
}
