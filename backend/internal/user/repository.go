package user

import (
	"devops-os/backend/internal/common"
	"devops-os/backend/internal/model"
	"fmt"
	"os"
)

type UserRepository struct {
	storage *common.FileStorage
}

func NewUserRepository(storage *common.FileStorage) *UserRepository {
	return &UserRepository{storage: storage}
}

func (r *UserRepository) GetByID(userID string) (*model.User, error) {
	var users []model.User
	path := "auth/users.json"

	if err := r.storage.ReadJSON(path, &users); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}

	return nil, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
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

func (r *UserRepository) Create(user *model.User) error {
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

func (r *UserRepository) Update(user *model.User) error {
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

func (r *UserRepository) GetAll() ([]model.User, error) {
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

func (r *UserRepository) Delete(userID string) error {
	var users []model.User
	path := "auth/users.json"

	if err := r.storage.ReadJSON(path, &users); err != nil {
		return err
	}

	for i, user := range users {
		if user.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return r.storage.WriteJSON(path, users)
		}
	}

	return fmt.Errorf("user not found: %s", userID)
}
