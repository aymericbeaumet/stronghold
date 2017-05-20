package stronghold

import (
	"bytes"
	"sort"
)

// Group is the class wrapping the stronghold groups
type Group struct {
	name  string
	users []*User
}

// NewGroup instanciates a new group.
func NewGroup(groupName string) *Group {
	return &Group{
		name: groupName,
	}
}

// AddUser adds a new user to the group
func (group *Group) AddUser(user *User) *Group {
	group.users = append(group.users, user)
	return group
}

// ToString creates a string representation of the group
func (group Group) ToString() string {
	identifiers := []string{}
	for _, user := range group.users {
		identifiers = append(identifiers, user.identifier)
	}
	sort.Strings(identifiers)
	var buffer bytes.Buffer
	for _, identifier := range identifiers {
		buffer.WriteString(identifier)
		buffer.WriteString("\n")
	}
	return buffer.String()
}
