package model

import (
	"bytes"
	"errors"

	"tpler/util"
)

// User 用户模型
type User struct {
	User  string `gorm:"PRIMARY_KEY"`
	Pass  string `json:"-"`
	Email string
}

// NewUser ...
func NewUser() *User {
	return new(User)
}

// Check ...
func (o *User) Check(user, pass string) error {
	u, err := o.Get(user)
	if err != nil {
		return err
	}

	if bytes.Compare([]byte(u.Pass), util.Hash([]byte(pass))) != 0 {
		return errors.New("password error")
	}

	return nil
}

// List ...
func (o *User) List() ([]*User, error) {
	var us []*User
	if err := db.Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// ListPageDemo ...
func (o *User) ListPageDemo(offset, limit int) (int, []*User, error) {
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}
	var us []*User
	if err := db.Order("user").Offset(offset).Limit(limit).Find(&us).Error; err != nil {
		return 0, nil, err
	}
	return count, us, nil
}

// Get ...
func (o *User) Get(user string) (*User, error) {
	var u User

	if user == "" {
		return nil, errors.New("user error")
	}

	if err := db.Find(&u, &User{User: user}).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// Save ...
func (o *User) Save(u *User) error {
	if u.User == "" {
		return errors.New("user error")
	}

	return db.Save(u).Error
}

// Delete ...
func (o *User) Delete(user string) error {
	if user == "" {
		return errors.New("user error")
	}
	return db.Delete(&User{User: user}).Error
}
