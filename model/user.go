package model

import (
	"errors"

	"tpler/common"
)

// User 用户模型
type User struct {
	ID    int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	User  string `gorm:"UNIQUE"`
	Pass  string `json:"-"`
	Email string
}

// NewUser ...
func NewUser() *User {
	return new(User)
}

// Check ...
func (o User) Check(user, pass string) error {
	if user == "" {
		return errors.New("user not found")
	}

	var u User
	if err := db.Find(&u, &User{User: user, Pass: string(common.Hash([]byte(pass)))}).Error; err != nil {
		return errors.New("user not found")
	}

	return nil
}

// List ...
func (User) List(offset, limit int) (int, []*User, error) {
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, nil, err
	}

	var us []*User
	if err := db.Offset(offset).Limit(limit).Find(&us).Error; err != nil {
		return 0, nil, err
	}
	return count, us, nil
}

// Get ...
func (User) Get(id int) (*User, error) {
	var u User

	if id == 0 {
		return nil, errors.New("id error")
	}

	if err := db.Find(&u, &User{ID: id}).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// Create ...
func (User) Create(u *User) error {
	if u.User == "" {
		return errors.New("user error")
	}
	return db.Create(u).Error
}

// Save ...
func (User) Save(u *User) error {
	if u.User == "" {
		return errors.New("user error")
	}

	return db.Save(u).Error
}

// Delete ...
func (User) Delete(ids []int) error {
	if len(ids) == 0 {
		return errors.New("ids error")
	}
	return db.Where("ID in (?)", ids).Delete(&User{}).Error
}
