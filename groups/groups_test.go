package groups

import (
	"testing"
)

func TestGroupCanBeCreatedUpdatedDeleted(test *testing.T) {
	var err error
	var group *Group

	group, err = NewGroup("root")

	if err != nil {
		test.Fail()
	}

	if group.ID == "" {
		test.Fail()
	}

	err = group.Create()
	if err != nil {
		test.Fail()
	}

	// err = user.Create()
	// // if err.Error() != "Email already used" or:
	// if err != ErrorEmailAlreadyUsed {
	// 	test.Fail()
	// }

	group.Name = "test"
	updatedGroup, err := group.Update()

	if err != nil {
		test.Fail()
	}

	if updatedGroup.ID != group.ID {
		test.Fail()
	}

	groupFromDataBase, err := GetGroupByID(group.ID)

	if groupFromDataBase == nil {
		test.Fail()
	}

	if groupFromDataBase.ID != group.ID {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}

	err = groupFromDataBase.Delete()

	if err != nil {
		test.Fail()
	}
}
