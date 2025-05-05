package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `json:"username"  gorm:"type:varchar(20);unique;not null"`
	PasswordDigest string `json:"password"  gorm:"type:varchar(20);not null"`
}

const PasswordCost = 12

// 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDisgest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDisgest), []byte(password))
	return err == nil
}
