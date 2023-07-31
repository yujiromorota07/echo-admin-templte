package util

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash - password hash interface
type PasswordHash interface {
	GeneratePasswordHash() string
	CompareHashAndPassword(hash string, password string) error
}

// Password - password util
type Password struct {
	Password string
	Hash     string
}

// ランダムなパスワードの生成（招待用）
func NewRandomPassword() Password {
	rand := RandomString(16)
	password := Password{
		Password: rand,
	}
	password.Hash = password.GeneratePasswordHash()
	return password
}

// 新規パスワード生成（パスワードを指定）
func NewPassword(newPassword string) Password {
	password := Password{
		Password: newPassword,
	}
	password.Hash = password.GeneratePasswordHash()
	return password
}

// GeneratePasswordHash - generate password hash
func (p Password) GeneratePasswordHash() string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// CompareHashAndPassword - compare hash and password
func (p Password) CompareHashAndPassword() error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(p.Password))
	return err
}
