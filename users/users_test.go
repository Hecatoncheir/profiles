package users

import (
	"testing"
)

func TestUserCanBeCreatedUpdatedDeleted(test *testing.T) {
	var err error
	var user *User

	user, err = NewUser("rasart.pro@gmail.com", "121354")

	if err != nil {
		test.Fail()
	}

	if user.ID == "" {
		test.Fail()
	}

	err = user.Create()
	if err != nil {
		test.Fail()
	}

	err = user.Create()
	// if err.Error() != "Email already used" or:
	if err != ErrorEmailAlreadyUsed {
		test.Fail()
	}

	user.Email = "test@mail"
	updatedUser, err := user.Update()

	if err != nil {
		test.Fail()
	}

	if updatedUser.ID != user.ID {
		test.Fail()
	}

	user.Email = "rasart.pro@gmail.com"
	err = user.Create()
	// if err.Error() != "User already exist" or:
	if err != ErrorUserAlreadyExist {
		test.Fail()
	}

	userFromDataBase, err := GetUserByID(user.ID)

	if userFromDataBase.ID != user.ID {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}

	if userFromDataBase == nil {
		test.Fail()
	}

	err = userFromDataBase.Delete()
	if err != nil {
		test.Fail()
	}
}
