package info

import (
	"time"
)

// 分组
type Group struct {
	GroupId     int64     `xorm:"pk"` // 谁的
	UserId      int64     // 所有者Id
	Title       string    // 标题
	UserCount   int       // 用户数
	CreatedTime time.Time `xorm:"created"`

	Users []User // 分组下的用户, 不保存, 仅查看
}

// 分组好友
type GroupUser struct {
	GroupUserId int64     `xorm:"pk"` // 谁的
	GroupId     int64     // 分组
	UserId      int64     //  用户
	CreatedTime time.Time `xorm:"created"`
}
