package model

import "errors"

// User 用户模型
type User struct {
	UID  string
	User string
	Pass string
}

// Check ...
func (o *User) Check(user, pass string) error {
	if user != "demo" || pass != "demo" {
		return errors.New("password error")
	}
	return nil
}
