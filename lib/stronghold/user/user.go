package stronghold

import (
	strongholdGroup "github.com/aymericbeaumet/stronghold/lib/stronghold/group"
)

// User is the class wrapping the stronghold users
type User struct {
	identifier string
	groups     []strongholdGroup.Group
}

// NewOptions wraps the User options
type NewOptions struct {
	Identifier string
	Group      *strongholdGroup.Group
}

// New initializes a new User
func New(options NewOptions) (*User, error) {
	u := new(User)
	u.identifier = options.Identifier
	u.groups = []strongholdGroup.Group{
		*options.Group,
	}
	return u, nil
}
