package groups

import (
	"errors"
	"log"

	rethink "gopkg.in/gorethink/gorethink.v3"
)

// ErrorGroupAlreadyExist emit when user already exist
var ErrorGroupAlreadyExist = errors.New("Group already exist")

var groupsDataBaseSession rethink.QueryExecutor

func init() {
	var err error

	groupsDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address: "localhost:28015",
	})

	err = checkDataBase("Users", groupsDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}

	groupsDataBaseSession, err = rethink.Connect(rethink.ConnectOpts{
		Address:  "localhost:28015",
		Database: "Users",
	})

	err = checkTable("Groups", groupsDataBaseSession)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// GetGroupByID return user from database
func GetGroupByID(id string) (*Group, error) {
	var err error
	cursor, err := rethink.Table("Groups").Get(id).Run(groupsDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return nil, err
	}

	group := &Group{}
	err = cursor.One(group)

	if err == rethink.ErrEmptyResult {
		// row not found
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return group, nil
}

func insertGroupToDataBase(group *Group) error {
	var err error

	_, err = rethink.Table("Groups").Insert(group).RunWrite(groupsDataBaseSession)

	if err != nil {
		return ErrorGroupAlreadyExist
	}

	return nil
}

func updateGroup(group *Group) (*Group, error) {
	if group.ID == "" {
		return group, errors.New("Group id is empty")
	}

	cursor, err := rethink.Table("Groups").Get(group.ID).Update(group).Run(groupsDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return group, err
	}

	updatedGroup, err := GetGroupByID(group.ID)

	if err != nil {
		return group, err
	}

	return updatedGroup, nil

}

func deleteGroupByID(id string) error {
	if id == "" {
		return errors.New("Group id is empty")
	}

	cursor, err := rethink.Table("Groups").Get(id).Delete().Run(groupsDataBaseSession)
	defer cursor.Close()

	if err != nil {
		return err
	}

	return nil
}
