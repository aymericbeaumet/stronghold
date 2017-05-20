package crypto

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// PrivateKey is the class wrapping a private key
type PrivateKey struct {
	entity *openpgp.Entity
}

// NewPrivateKey instanciates a new private key.
func NewPrivateKey(identifier string, gpg string, homedir string) *PrivateKey {
	privateKey := &PrivateKey{}
	// spawn process
	name := "gpg"
	if gpg != "" {
		name = gpg
	}
	args := []string{
		"--export-secret-key",
		"--armor",
	}
	if homedir != "" {
		args = append(args, []string{
			"--homedir",
			homedir,
		}...)
	}
	args = append(args, []string{
		"<" + strings.Trim(identifier, "<>") + ">",
	}...)
	command := exec.Command(name, args...)
	// grab stdout
	stdout, err := command.StdoutPipe()
	if err != nil {
		panic(err)
	}
	// start process
	if err = command.Start(); err != nil {
		panic(err)
	}
	// decode armored private key
	block, err := armor.Decode(stdout)
	if err != nil {
		panic(err)
	}
	// check private key type
	if block.Type != openpgp.PrivateKeyType {
		panic(errors.New("Unsupported key type: " + block.Type))
	}
	// parse entity
	reader := packet.NewReader(block.Body)
	entity, err := openpgp.ReadEntity(reader)
	if err != nil {
		panic(err)
	}
	// expose the entity
	privateKey.entity = entity
	return privateKey
}

// Sign reads the given payload and generates the signature.
func (privateKey PrivateKey) Sign(payload io.Reader, signature io.Writer) error {
	return openpgp.DetachSign(signature, privateKey.entity, payload, nil)
}

// SignFile reads the given file and writes the signature on the disk.
func (privateKey PrivateKey) SignFile(filePath string, signaturePath string) error {
	fileHandler, openError := os.Open(filePath)
	if openError != nil {
		return openError
	}
	defer fileHandler.Close()
	signatureHandler, createError := os.Create(signaturePath)
	if createError != nil {
		return createError
	}
	defer signatureHandler.Close()
	return privateKey.Sign(fileHandler, signatureHandler)
}
