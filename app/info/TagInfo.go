package info

import (
	"time"
)

// 这里主要是为了统计每个tag的note数目
// 暂时没用
/*
type TagNote struct {
	TagId   bson.ObjectId `bson:"_id,omitempty"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	UserId  bson.ObjectId `bson:"UserId"`
	Tag   string        `Title`   // 标题
	NoteNum int           `NoteNum` // note数目
}
*/

// 每个用户一条记录, 存储用户的所有tags
type Tag struct {
	UserId int64
	Tags   []string
}

// v2 版标签
type NoteTag struct {
	TagId       int64
	UserId      int64  // 谁的
	Tag         string // UserId, Tag是唯一索引
	Usn         int    // Update Sequence Number
	Count       int    // 笔记数
	CreatedTime time.Time
	UpdatedTime time.Time
	IsDeleted   bool // 删除位
}

type TagCount struct {
	TagCountId int64 `xorm:"pk"`
	UserId     int64 // 谁的
	Tag        string
	IsBlog     bool // 是否是博客的tag统计
	Count      int  // 统计数量
}

/*
type TagsCounts []TagCount
func (this TagsCounts) Len() int {
	return len(this)
}
func (this TagsCounts) Less(i, j int) bool {
	return this[i].Count > this[j].Count
}
func (this TagsCounts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
*/
