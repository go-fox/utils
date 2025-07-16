package password

// Crypto 密码加解密接口
type Crypto interface {
	// Encrypt 加密
	Encrypt(password string) (string, error)
	// Verify 验证密码是否正确
	Verify(password, hashedPassword string) (bool, error)
}
