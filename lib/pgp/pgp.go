package pgp

import (
	"errors"
	"os/exec"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// PGP
type PGP struct {
	entity *openpgp.Entity
}

// NewOptions is the structure used to pass the optional parameters
// for the LoadPrivateKey function.
type NewOptions struct {
	Identifier string
	Gpg        string
	Homedir    string
}

// New loads a private key, given an identifier, by spawing a gpg
// process.
func New(options NewOptions) (*PGP, error) {
	p := new(PGP)
	// spawn process
	name := "gpg"
	if len(options.Gpg) > 0 {
		name = options.Gpg
	}
	args := []string{
		"--export-secret-key",
		"--armor",
	}
	if len(options.Homedir) > 0 {
		args = append(args, []string{
			"--homedir", options.Homedir,
		}...)
	}
	args = append(args, []string{
		"<" + strings.Trim(options.Identifier, "<>") + ">",
	}...)
	command := exec.Command(name, args...)
	// grab stdout
	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}
	// start process
	if err = command.Start(); err != nil {
		return nil, err
	}
	// decode armored private key
	block, err := armor.Decode(stdout)
	if err != nil {
		return nil, err
	}
	// check private key type
	if block.Type != openpgp.PrivateKeyType {
		return nil, errors.New("Unsupported key type: " + block.Type)
	}
	// parse entity
	reader := packet.NewReader(block.Body)
	entity, err := openpgp.ReadEntity(reader)
	if err != nil {
		return nil, err
	}
	p.entity = entity
	return p, nil
}
