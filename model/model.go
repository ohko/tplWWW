package model

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"tpler/common"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // ...
	"github.com/ohko/logger"
	// _ "github.com/jinzhu/gorm/dialects/postgres" // ...
)

// ...
var (
	db        *gorm.DB
	DBUser    = NewUser()
	DBMember  = NewMember()
	DBSetting = NewSetting()
)

type model struct{}

func (o *model) Print(arg ...interface{}) {
	common.LL.LogCalldepth(2, logger.LoggerLevel0Debug, arg[3:]...)
}

// Init 初始化数据库
func Init(dbPath string) error {
	var err error

	os.MkdirAll(filepath.Dir(dbPath), 0755)
	if db, err = gorm.Open("sqlite3", dbPath); err != nil {
		// if db, err = gorm.Open("postgres", "postgres://user:pass@host/database?sslmode=disable"); err != nil {
		return err
	}
	db.SetLogger(&model{})
	db.LogMode(os.Getenv("DEBUG") != "")
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(2)
	db.SetNowFuncOverride(func() time.Time {
		return time.Now().In(common.TimeLocation)
	})

	if err := db.AutoMigrate(&Member{}, &User{}, &Setting{}).Error; err != nil {
		return err
	}

	var m Member
	if err := db.First(&m).Error; err != nil {
		if err := db.Save(&Member{User: "admin", Pass: string(common.Hash([]byte("admin")))}).Error; err != nil {
			return err
		}

		// 创建50个测试账号
		for i := 0; i < 50; i++ {
			if err := db.Save(&User{User: fmt.Sprintf("user-%d", i), Email: fmt.Sprintf("email-%d@xx.com", i)}).Error; err != nil {
				return err
			}
		}
	}

	// 初始化系统配置
	defaultSetting := []Setting{
		Setting{Key: "Int1", Type: 0, Int: 1},
		Setting{Key: "String2", Type: 1, String: "string2"},
		Setting{Key: "Bool3", Type: 2, Bool: true},
	}
	for _, v := range defaultSetting {
		var d Setting
		if err := db.First(&d, &Setting{Key: v.Key}).Error; err != nil {
			common.LL.Log0Debug(v.Key)
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
