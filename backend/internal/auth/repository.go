package auth

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"fmt"
	"os"
)

type AuthRepository struct {
	storage *common.FileStorage
}

func NewAuthRepository(storage *common.FileStorage) *AuthRepository {
	return &AuthRepository{storage: storage}
}

func (r *AuthRepository) GetUserByUsername(username string) (*model.User, error) {
	var users []model.User
	path := "auth/users.json"

	if err := r.storage.ReadJSON(path, &users); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, nil
}

func (r *AuthRepository) CreateUser(user *model.User) error {
	var users []model.User
	path := "auth/users.json"

	if r.storage.PathExists(path) {
		if err := r.storage.ReadJSON(path, &users); err != nil {
			return err
		}
	}

	users = append(users, *user)
	return r.storage.WriteJSON(path, users)
}

func (r *AuthRepository) UpdateUser(user *model.User) error {
	var users []model.User
	path := "auth/users.json"

	if err := r.storage.ReadJSON(path, &users); err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			return r.storage.WriteJSON(path, users)
		}
	}

	return fmt.Errorf("user not found: %s", user.ID)
}

func (r *AuthRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	path := "auth/users.json"

	if !r.storage.PathExists(path) {
		return []model.User{}, nil
	}

	if err := r.storage.ReadJSON(path, &users); err != nil {
		return nil, err
	}

	return users, nil
}
