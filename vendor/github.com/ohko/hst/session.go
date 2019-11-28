package hst

import (
	"time"
)

type sessionData struct {
	Data   interface{}
	Expire time.Time
}

// Session ...
type Session interface {
	GetCookieExpire() time.Duration
	Set(c *Context, key string, value interface{}, expire time.Duration) error
	Get(c *Context, key string) (interface{}, error)
	Destory(c *Context) error
}
