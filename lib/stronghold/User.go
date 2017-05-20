package stronghold

// User is the class wrapping the stronghold users
type User struct {
	identifier string
	groups     []Group
}

// NewUser initializes a new User
func NewUser(identifier string) *User {
	return &User{
		identifier: identifier,
	}
}
