package crypto

// KeyBox wraps the access to a keybox
type KeyBox struct {
	privateKey *PrivateKey
}

// NewKeyBox instanciates a new KeyBox
func NewKeyBox() *KeyBox {
	return &KeyBox{}
}

// GetPrivateKey returns the key box private key
func (keyBox KeyBox) GetPrivateKey() *PrivateKey {
	return keyBox.privateKey
}

// SetPrivateKey loads a private key into the key box
func (keyBox *KeyBox) SetPrivateKey(privateKey *PrivateKey) *KeyBox {
	keyBox.privateKey = privateKey
	return keyBox
}
