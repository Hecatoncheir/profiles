package users

import (
	"errors"
	"log"

	rethink "gopkg.in/gorethink/gorethink.v3"
)

// ErrorUserAlreadyExist emit when user already exist
var ErrorUserAlreadyExist = errors.New("User already exist")

// ErrorEmailAlreadyUsed emit when email for new user elready used
var ErrorEmailAlreadyUsed = errors.New("Email already used")

var userDataBaseSession rethink.QueryExecutor

func init() {
	var err error

	userDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address: "localhost:28015",
	})

	err = checkDataBase("Users", userDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	userDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address:  "localhost:28015",
		Database: "Users",
	})

	err = checkTable("Users", userDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = rethink.Table("Users").IndexCreate("Email").Run(userDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// GetUserByID return user from database
func GetUserByID(id string) (*User, error) {
	var err error
	cursor, err := rethink.Table("Users").Get(id).Run(userDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return nil, err
	}

	user := User{}
	err = cursor.One(&user)

	if err == rethink.ErrEmptyResult {
		// row not found
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func checkUserEmail(email string) error {
	cursor, _ := rethink.Table("Users").Filter(rethink.Row.Field("Email").Eq(email)).Run(userDataBaseSession)
	usersWithExistEmails := []User{}

	cursor.All(&usersWithExistEmails)

	if len(usersWithExistEmails) > 0 {
		return ErrorEmailAlreadyUsed
	}

	return nil
}

func insertUserToDataBase(user *User) error {
	var err error

	err = checkUserEmail(user.Email)
	if err != nil {
		return err
	}

	_, err = rethink.Table("Users").Insert(user).RunWrite(userDataBaseSession)

	if err != nil {
		return ErrorUserAlreadyExist
	}

	return nil
}

func updateUser(user *User) (*User, error) {
	if user.ID == "" {
		return user, errors.New("User id is empty")
	}

	cursor, err := rethink.Table("Users").Get(user.ID).Update(user).Run(userDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return user, err
	}

	updatedUser, err := GetUserByID(user.ID)

	if err != nil {
		return user, err
	}

	return updatedUser, nil

}

func deleteUserByID(id string) error {
	if id == "" {
		return errors.New("User id is empty")
	}

	cursor, err := rethink.Table("Users").Get(id).Delete().Run(userDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return err
	}

	return nil
}
