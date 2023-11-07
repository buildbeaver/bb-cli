package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// PublicSharedSecretToken is a shared secret token that is safe to share publicly with the token owner.
type PublicSharedSecretToken struct {
	id   string
	data []byte
}

func (m PublicSharedSecretToken) ID() string {
	return m.id
}

// String returns the string representation of the token that can be shared with the token owner.
func (m PublicSharedSecretToken) String() string {
	str := fmt.Sprintf("%s:%s", m.id, hex.EncodeToString(m.data))
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (m PublicSharedSecretToken) IsValid(salt []byte, hash []byte) (bool, error) {
	computedHash := sha256.Sum256(append(m.data, salt...))
	return bytes.Equal(computedHash[:], hash), nil
}
