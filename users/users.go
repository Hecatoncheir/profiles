package users

import (
	uuid "github.com/satori/go.uuid"
)

// User is a main structure of user entity
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Groups   []string
}

// NewUser is a constructor for User
func NewUser(email string, password string) (*User, error) {
	id := uuid.NewV4().String()
	user := User{ID: id, Email: email, Password: password}
	return &user, nil
}

// Create method for create user in database
func (user *User) Create() error {
	err := insertUserToDataBase(user)
	return err
}

// Update need for delete user
func (user *User) Update() (*User, error) {
	updatedUser, err := updateUser(user)
	if err != nil {
		return user, err
	}

	return updatedUser, nil
}

// Delete need for delete user and all data
func (user *User) Delete() error {
	var err error

	err = deleteUserByID(user.ID)
	return err
}
