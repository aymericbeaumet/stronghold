package stronghold

import (
	"errors"
	"os"
	"path/filepath"

	strongholdGroup "github.com/aymericbeaumet/stronghold/lib/stronghold/group"
	strongholdUser "github.com/aymericbeaumet/stronghold/lib/stronghold/user"
)

// Stronghold is the class wrapping the stronghold: including filesystem,
// permissions, etc...
type Stronghold struct {
	path   string
	groups []strongholdGroup.Group
	users  []strongholdUser.User
}

type NewOptions struct {
	WorkingDirectory string
}

// New creates a new Stronghold
func New(options NewOptions) (*Stronghold, error) {
	s := new(Stronghold)
	rootDirectory, err := getRootDirectory(options.WorkingDirectory)
	if err != nil {
		return nil, err
	}
	s.path = filepath.Join(rootDirectory, markers["stronghold"])
	return s, nil
}

// Exists checks whether the corresponding stronghold exists on the disk.
func (s *Stronghold) Exists() (bool, error) {
	_, err := os.Stat(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Dump dumps the content of the stronghold on the disk.
func (s *Stronghold) Dump(overwrite bool) error {
	if !overwrite {
		exists, err := s.Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A stronghold already exists in: " + s.path)
		}
	}
	if err := os.MkdirAll(s.path, 0700); err != nil {
		return err
	}
	return nil
}

// CreateGroup creates a new group in the stronghold
func (s *Stronghold) CreateGroup(groupOptions strongholdGroup.NewOptions) (*strongholdGroup.Group, error) {
	g, err := strongholdGroup.New(groupOptions)
	if err != nil {
		return nil, err
	}
	s.groups = append(s.groups, *g)
	return g, nil
}

// AddUser adds a new user to the stronghold
func (s *Stronghold) AddUser(userOptions strongholdUser.NewOptions) (*strongholdUser.User, error) {
	u, err := strongholdUser.New(userOptions)
	if err != nil {
		return nil, err
	}
	s.users = append(s.users, *u)
	return u, nil
}

var markers = map[string]string{
	"stronghold": ".stronghold",
	"git":        ".git",
	"mercurial":  ".hg",
	"svn":        ".svn",
}

func getRootDirectory(workingDirectory string) (string, error) {
	currentDirectory := workingDirectory
	for {
		for _, marker := range markers {
			stat, err := os.Stat(filepath.Join(currentDirectory, marker))
			if err == nil && stat.IsDir() {
				return currentDirectory, nil
			}
		}
		if currentDirectory == "/" {
			return workingDirectory, nil
		}
		currentDirectory = filepath.Dir(currentDirectory)
	}
}
