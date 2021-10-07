package jauth

import (
	"fmt"

	"github.com/dkkyeremateng/jauth/keystore"
)

func NewAuth(privatePEM []byte, fileName, activeKID string) (*Auth, error) {
	// Construct a key store based on the key files stored in
	// the specified directory.
	ks, err := keystore.NewFS(privatePEM, fileName)
	if err != nil {
		return nil, fmt.Errorf("reading keys: %w", err)
	}

	fmt.Println(activeKID)

	auth, err := new(activeKID, ks)
	if err != nil {
		return nil, fmt.Errorf("constructing auth: %w", err)
	}

	return auth, nil
}
