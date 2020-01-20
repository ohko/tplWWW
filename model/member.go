package model

import (
	"errors"
	"fmt"
	"strings"
	"tpler/common"
)

// Member 管理员模型
type Member struct {
	User string `gorm:"PRIMARY_KEY"`
	Pass string `json:"-"`
}

// ResetAdmin 重置管理员
func ResetAdmin(newPass string) error {
	if newPass == "" {
		return errors.New("new pass is empty")
	}

	if err := db.Save(&Member{
		User: "admin",
		Pass: string(common.Hash([]byte(newPass))),
	}).Error; err != nil {
		return err
	}

	fmt.Println(strings.Repeat("=", 20) + "\nuser:admin\npassword:******\n" + strings.Repeat("=", 20))
	return nil
}

// NewMember ...
func NewMember() *Member {
	return new(Member)
}

// Check ...
func (Member) Check(user, pass string) error {
	var u Member
	if err := db.First(&u, &Member{
		User: user,
		Pass: string(common.Hash([]byte(pass))),
	}).Error; err != nil {
		return err
	}
	return nil
}

// Save ...
func (Member) Save(m *Member) error {
	if m.User == "" {
		return errors.New("user error")
	}

	return db.Save(m).Error
}
