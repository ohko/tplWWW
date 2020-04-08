package model

import (
	"database/sql/driver"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"tpler/common"
	"unicode"

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

type model struct {
	ll *logger.Logger
}

func (o model) Print(arg ...interface{}) {
	o.ll.Log0Debug(LogFormatter(arg...)...)
}

var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	sqlRegexp := regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp := regexp.MustCompile(`\$\d+`)
	isPrintable := func(s string) bool {
		for _, r := range s {
			if !unicode.IsPrint(r) {
				return false
			}
		}
		return true
	}
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			// currentTime     = "\033[33m" + time.Now().In(common.TimeLocation).Format("2006-01-02 15:04:05") + "\033[0m"
			currentTime = time.Now().In(common.TimeLocation).Format("2006-01-02 15:04:05")
			// source      = fmt.Sprintf("\033[35m%v\033[0m", values[1])
			source = filepath.Base(fmt.Sprintf("%v", values[1].(string)))
		)

		messages = []interface{}{currentTime, source}

		if len(values) == 2 {
			//remove the line break
			currentTime = currentTime[1:]
			//remove the brackets
			// source = fmt.Sprintf("\033[35m%v\033[0m", values[1])
			source = fmt.Sprintf("%v", values[1])

			messages = []interface{}{currentTime, source}
		}

		if level == "sql" {
			// duration
			// messages = append(messages, fmt.Sprintf("\033[36;1m[%.2fms]\033[0m ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			messages = append(messages, fmt.Sprintf("[%.2fms]", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						if t.IsZero() {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", "0000-00-00 00:00:00"))
						} else {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
						}
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			// messages = append(messages, fmt.Sprintf("\033[36;31m[%v]\033[0m ", strconv.FormatInt(values[5].(int64), 10)+" rows"))
			messages = append(messages, fmt.Sprintf("[%v]", strconv.FormatInt(values[5].(int64), 10)+" rows"))
		} else {
			messages = append(messages, "\033[31;1m")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\033[0m")
		}
	}

	return
}

// Init 初始化数据库
func Init(dbPath string) error {
	var err error

	_ = os.MkdirAll(filepath.Dir(dbPath), 0755)
	if db, err = gorm.Open("sqlite3", dbPath); err != nil {
		// if db, err = gorm.Open("postgres", "postgres://user:pass@host/database?sslmode=disable"); err != nil {
		return err
	}
	sqlLog := common.LL.Fork("SQL")
	sqlLog.SetFlags(0)
	db.SetLogger(&model{ll: sqlLog})
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
