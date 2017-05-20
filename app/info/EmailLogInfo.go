package info

import (
	"time"
)

// 发送邮件
type EmailLog struct {
	LogId int64 `xorm:"pk"`

	Email   string // 发送者
	Subject string // 主题
	Body    string // 内容
	Msg     string // 发送失败信息
	Ok      bool   // 发送是否成功

	CreatedTime time.Time `xorm:"created"`
}
