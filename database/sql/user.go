package sql

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12
)

type User struct {
	gorm.Model
	UserName string `gorm:"primaryKey;column:user_name" `
	Password string
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
