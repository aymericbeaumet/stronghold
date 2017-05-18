package stronghold

// Group is the class wrapping the stronghold groups
type Group struct {
	name string
}

type NewOptions struct {
	Name string
}

func New(options NewOptions) (*Group, error) {
	u := new(Group)
	u.name = options.Name
	return u, nil
}
