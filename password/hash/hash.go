package hash

import (
	"crypto/subtle"

	"golang.org/x/crypto/argon2"

	"github.com/go-fox/utils/password"
)

var _ password.Crypto = (*Hash)(nil)

type Hash struct {
	salt string
}

// NewHash 构造函数
func NewHash(salt string) *Hash {
	return &Hash{
		salt: salt,
	}
}

func (h *Hash) Encrypt(password string) (string, error) {
	time := uint32(1)           // 迭代次数
	memory := uint32(64 * 1024) // 内存使用量（以字节为单位）
	threads := uint8(4)         // 并行线程数
	keyLen := uint32(32)        // 输出密钥长度为 32 字节
	hashedPassword := argon2.IDKey([]byte(password), []byte(h.salt), time, memory, threads, keyLen)
	return string(hashedPassword), nil
}

func (h *Hash) Verify(password, hashedPassword string) (bool, error) {
	time := uint32(1)           // 迭代次数
	memory := uint32(64 * 1024) // 内存使用量（以字节为单位）
	threads := uint8(4)         // 并行线程数
	keyLen := uint32(32)        // 输出密钥长度为 32 字节
	derivedKey := argon2.IDKey([]byte(password), []byte(h.salt), time, memory, threads, keyLen)
	return subtle.ConstantTimeCompare([]byte(hashedPassword), derivedKey) == 1, nil
}
