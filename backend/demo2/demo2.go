package demo2

import (
	"fmt"
	"time"
	"tpler/common"

	"github.com/ohko/logger"
)

// Demo2 ...
type Demo2 struct {
	start time.Time
	last  time.Time
}

// Register ...
func (o *Demo2) Register() error {
	o.start = time.Now()
	go o.do()
	go o.do()

	return nil
}

// Status ...
func (o *Demo2) Status() string {
	return fmt.Sprintf("demo2: %s ~ %s", o.start, o.last)
}

func (o *Demo2) do() {
	for {
		time.Sleep(time.Hour)
		o.last = time.Now()

		loger := logger.NewLogger(common.LLFile)
		loger.SetPrefix(time.Now().Format(time.RFC3339Nano))
		o.step1(loger)
	}
}

func (o *Demo2) step1(loger *logger.Logger) {
	loger.Log0Debug("Demo2 :: step1")
	o.step2(loger)
}

func (o *Demo2) step2(loger *logger.Logger) {
	loger.Log0Debug("Demo2 :: step2")
	o.step3(loger)
}

func (o *Demo2) step3(loger *logger.Logger) {
	loger.Log0Debug("Demo2 :: step3")
}
