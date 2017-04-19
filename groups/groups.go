package groups

import uuid "github.com/satori/go.uuid"

// Group entity
type Group struct {
	ID   string
	Name string
}

// NewGroup constructor for Group entity
func NewGroup(name string) (*Group, error) {
	id := uuid.NewV4().String()
	group := Group{ID: id, Name: name}
	return &group, nil
}

// Create need for save user object in database
func (group *Group) Create() error {
	err := insertGroupToDataBase(group)
	return err
}

// Update need for update group
func (group *Group) Update() (*Group, error) {
	updatedGroup, err := updateGroup(group)
	if err != nil {
		return group, err
	}

	return updatedGroup, nil
}

// Delete method for remove group object from database
func (group *Group) Delete() error {
	err := deleteGroupByID(group.ID)
	return err
}
