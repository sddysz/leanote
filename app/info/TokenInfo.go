package info

import (
	"time"
)

// 随机token
// 验证邮箱
// 找回密码

// token type
const (
	TokenPwd = iota
	TokenActiveEmail
	TokenUpdateEmail
)

// 过期时间
const (
	PwdOverHours         = 2.0
	ActiveEmailOverHours = 48.0
	UpdateEmailOverHours = 2.0
)

type Token struct {
	UserId      int64
	Email       string
	Token       string
	Type        int
	CreatedTime time.Time `xorm:"created"`
}
