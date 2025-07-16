package bcrypthash

import (
	"github.com/go-fox/utils/password"

	"golang.org/x/crypto/bcrypt"
)

var _ password.Crypto = (*BcryptHash)(nil)

type BcryptHash struct {
}

func NewBcryptHash() *BcryptHash {
	return &BcryptHash{}
}

func (b *BcryptHash) Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *BcryptHash) Verify(password, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil, err
}
