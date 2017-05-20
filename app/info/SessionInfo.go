package info

import (
	"time"
)

// http://docs.mongodb.org/manual/tutorial/expire-data/
type Session struct {
	Id int64 `xorm:"pk"` // 没有意义

	SessionId string // SessionId

	LoginTimes int    // 登录错误时间
	Captcha    string // 验证码

	UserId string // API时有值UserId

	CreatedTime time.Time `xorm:"createdpk"`
	UpdatedTime time.Time `xorm:"updated"` // 更新时间, expire这个时间会自动清空
}
