package model

import (
	"os"
	"path/filepath"
	"time"
	"tpler/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // ...

	// _ "github.com/jinzhu/gorm/dialects/postgres" // ...
	"github.com/ohko/logger"
)

var (
	ll  *logger.Logger
	db  *gorm.DB
	loc *time.Location
)

// Init ...
func Init(lll *logger.Logger, dbPath string) error {
	ll = lll

	var err error
	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	if err := initDB(dbPath); err != nil {
		return err
	}

	return nil
}

// initDB 初始化数据库
func initDB(dbPath string) error {
	var err error

	os.MkdirAll(filepath.Dir(dbPath), 0755)
	if db, err = gorm.Open("sqlite3", dbPath); err != nil {
		// if db, err = gorm.Open("postgres", "postgres://user:pass@host/database?sslmode=disable"); err != nil {
		return err
	}
	if os.Getenv("DEBUG") != "" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.SetNowFuncOverride(func() time.Time {
		return time.Now().In(loc)
	})

	if err := db.AutoMigrate(&Member{}, &User{}, &Setting{}).Error; err != nil {
		return err
	}

	var m Member
	if err := db.First(&m).Error; err != nil {
		if err := db.Save(&Member{User: "admin", Pass: string(util.Hash([]byte("admin")))}).Error; err != nil {
			return err
		}
	}

	// 初始化系统配置
	defaultSetting := []Setting{
		Setting{Key: "Int1", Int: 1},
		Setting{Key: "String2", String: "string2"},
		Setting{Key: "Bool3", Bool: true},
	}
	for _, v := range defaultSetting {
		var d Setting
		if err := db.First(&d, &Setting{Key: v.Key}).Error; err != nil {
			ll.Log0Debug(v.Key)
			if err := db.Save(&v).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Close ...
func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
