package scrypthash

import (
	"crypto/subtle"

	"golang.org/x/crypto/scrypt"

	"github.com/go-fox/utils/password"
)

var _ password.Crypto = (*ScryptHash)(nil)

type ScryptHash struct {
	salt string
}

// NewScryptHash 创建一个scrypt加密器
func NewScryptHash(salt string) *ScryptHash {
	return &ScryptHash{salt: salt}
}

func (s *ScryptHash) Encrypt(password string) (string, error) {
	N := 16384   // 迭代次数，表示对哈希函数的重复计算次数。值越大，计算成本越高，安全性越好，但性能也会受到影响。通常建议选择一个大于 2^14 的值，例如 16384
	r := 8       // 块大小，表示每次哈希计算中使用的内存块大小。较大的值会增加内存消耗，并增加攻击成本。通常建议选择一个合适的值，例如 8
	p := 1       // 并行度，表示并行计算的线程或处理器数目。较高的值会增加计算成本，适当的值取决于你的硬件配置。通常建议选择一个合适的值，例如 1
	keyLen := 32 // 参数指定生成的密钥的长度
	hashedPassword, err := scrypt.Key([]byte(password), []byte(s.salt), N, r, p, keyLen)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *ScryptHash) Verify(password, hashedPassword string) (bool, error) {
	N := 16384   // 迭代次数，表示对哈希函数的重复计算次数。值越大，计算成本越高，安全性越好，但性能也会受到影响。通常建议选择一个大于 2^14 的值，例如 16384
	r := 8       // 块大小，表示每次哈希计算中使用的内存块大小。较大的值会增加内存消耗，并增加攻击成本。通常建议选择一个合适的值，例如 8
	p := 1       // 并行度，表示并行计算的线程或处理器数目。较高的值会增加计算成本，适当的值取决于你的硬件配置。通常建议选择一个合适的值，例如 1
	keyLen := 32 // 参数指定生成的密钥的长度
	derivedKey, err := scrypt.Key([]byte(password), []byte(s.salt), N, r, p, keyLen)
	if err != nil {
		return false, err
	}
	return subtle.ConstantTimeCompare([]byte(hashedPassword), derivedKey) == 1, nil
}
