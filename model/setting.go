package model

import (
	"errors"
)

// Setting 配置模型
type Setting struct {
	Key    string `gorm:"PRIMARY_KEY"`
	Type   int    // 0=int/1=string/2=bool
	Desc   string // 说明
	Int    int
	String string
	Bool   bool
}

// NewSetting ...
func NewSetting() *Setting {
	return new(Setting)
}

// List ...
func (Setting) List() ([]*Setting, error) {
	var us []*Setting
	if err := db.Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// GetSetting 获取配置，错误或没有返回设定的默认值
func (Setting) GetSetting(key string, defaultValue *Setting) *Setting {
	var u Setting

	if err := db.Find(&u, &Setting{Key: key}).Error; err != nil {
		return defaultValue
	}

	return &u
}

// Get ...
func (Setting) Get(key string) (*Setting, error) {
	var u Setting

	if err := db.Find(&u, &Setting{Key: key}).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

// Save ...
func (Setting) Save(u *Setting) error {
	if u.Key == "" {
		return errors.New("key error")
	}

	return db.Save(u).Error
}

// Delete ...
func (Setting) Delete(key string) error {
	if key == "" {
		return errors.New("key error")
	}
	return db.Delete(&Setting{Key: key}).Error
}
