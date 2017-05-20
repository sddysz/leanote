package info

import (
	"time"
)

// 举报
type Report struct {
	ReportId int64 `xorm:"pk"`
	NoteId   int64

	UserId int64  // UserId回复ToUserId
	Reason string // 评论内容

	CommentId int64 // 对某条评论进行回复

	CreatedTime time.Time `xorm:"Created"`
}
