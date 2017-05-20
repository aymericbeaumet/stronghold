package stronghold

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymericbeaumet/stronghold/lib/crypto"
)

// Vault is the class wrapping the access to the secret vault.
type Vault struct {
	path       string
	groupsPath string
	keyBox     *crypto.KeyBox
	groups     []*Group
	users      []*User
}

// NewVault instanciates a new Vault based on the given workingDirectory.
func NewVault(workingDirectory string, keyBox *crypto.KeyBox) *Vault {
	path := findVaultDirectory(workingDirectory)
	return &Vault{
		path:       path,
		groupsPath: filepath.Join(path, "groups"),
		keyBox:     keyBox,
	}
}

// Exists checks whether the corresponding vault exists on the disk.
func (vault *Vault) Exists() bool {
	_, err := os.Stat(vault.path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

// AddGroup creates a new group in the vault
func (vault *Vault) AddGroup(group *Group) *Vault {
	vault.groups = append(vault.groups, group)
	return vault
}

// AddUser adds a new user to the vault
func (vault *Vault) AddUser(user *User) *Vault {
	vault.users = append(vault.users, user)
	return vault
}

// Commit dumps the content of the vault on the disk, including:
//   - the groups
func (vault Vault) Commit() error {
	if vault.Exists() {
		return errors.New("A vault already exists in: " + vault.path)
	}
	if err := vault.commitGroups(); err != nil {
		return err
	}
	return nil
}

func (vault Vault) commitGroups() error {
	dirErr := os.MkdirAll(vault.groupsPath, 0755)
	if dirErr != nil {
		panic(dirErr)
	}
	for _, group := range vault.groups {
		vault.commitGroup(group)
	}
	return nil
}

func (vault Vault) commitGroup(group *Group) error {
	groupPath := filepath.Join(vault.groupsPath, group.name)
	groupContent := []byte(group.ToString())
	writeFileError := ioutil.WriteFile(groupPath, groupContent, 0644)
	if writeFileError != nil {
		panic(writeFileError)
	}
	signaturePath := groupPath + ".sig"
	signFileError := vault.keyBox.GetPrivateKey().SignFile(groupPath, signaturePath)
	if signFileError != nil {
		panic(signFileError)
	}
	return nil
}

var markers = map[string]string{
	"stronghold": ".stronghold",
	"git":        ".git",
	"mercurial":  ".hg",
	"svn":        ".svn",
}

func findVaultDirectory(workingDirectory string) string {
	currentDirectory := workingDirectory
	for {
		for _, marker := range markers {
			_, err := os.Stat(filepath.Join(currentDirectory, marker))
			if err != nil {
				continue
			}
			break
		}
		if currentDirectory == "/" {
			currentDirectory = workingDirectory
			break
		}
		currentDirectory = filepath.Dir(currentDirectory)
	}
	return filepath.Join(currentDirectory, markers["stronghold"])
}
