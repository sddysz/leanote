package info

import (
	"time"
)

// 共享笔记本
/*
{
	userId, 谁的共享笔记本
	shared: {
		userId1: { // 哪个用户共享的
			seq:
			"defaultNotebook": {noteId:{Seq:33}, noteId2:{}} // 默认笔记本, 里面是全部的笔记
			"notebookIds":{notebookId:{}, notebookId2:{}}} // 其它笔记本
		},
		userId2: {

		}
	}
}
*/

// 以后可能含有其它信息
type EachSharedNote struct {
	Seq int
}
type EachSharedNotebook struct {
	Seq int
}

// 每一个用户共享给的note, notebook
type EachSharedNotebookAndNotes struct {
	Seq             int                           // 共享给谁的所在序号
	DefaultNotebook map[string]EachSharedNote     // noteId => 共享的note
	Notebooks       map[string]EachSharedNotebook // notebookId => 共享的notebook
}

type SharedNotebookAndNotes struct {
	UserId int64                                 `xorm:"pk"`
	Shared map[string]EachSharedNotebookAndNotes // userId =>
}

/*
{
	UserId
	Notes: {noteId => [userId1, userId2], noteId2: []}
	Notebooks: {notebookId => [], notebookId2 => []}
}
*/

// 用户正在共享的notebook, note
type SharingNotebookAndNotes struct {
	UserId    int64               `xorm:"pk"`
	Notes     map[string][]string // noteId => []string{userId1, userId2}
	Notebooks map[string][]string // notebookId => []string{userId1, userId2}
}

// 以上以后再用, 现不用
//----------------------------------------

/*
每一个sharing一条记录, 这样更好操作
类似Notebook的
{
	pk
	userId
	toUserId
	noteId
	notebookId
	seq
}
*/

type ShareNotebook struct {
	ShareNotebookId int64 `xorm:"pk"` // 必须要设置xorm:"pk" 不然mgo不会认为是主键
	UserId          int64
	ToUserId        int64
	ToGroupId       int64 // 分享给的用户组
	ToGroup         Group // 仅仅为了显示, 不存储, 分组信息
	NotebookId      int64
	Seq             int       // 排序
	Perm            int       `xorm:"Perm"` // 权限, 其下所有notes 0只读, 1可修改
	CreatedTime     time.Time `xorm:"Created"`
	//	IsDefault       bool          `IsDefault` // 是否是默认共享notebook, perm seq=-9999999, NotebookId=null
}

/*
[
	ShareNotebooks,
		Subs: [ShareNotebooks, ]
	ShareNotebooks
]
*/
type SubShareNotebooks []ShareNotebooks
type ShareNotebooks struct {
	Notebook
	ShareNotebook
	Subs SubShareNotebooks

	// Notebook与ShareNotebook公用的不能生成json
	Seq        int
	NotebookId int64
	IsDefault  bool // 是否是默认笔记本
}

// SubShareNotebook sort
func (this SubShareNotebooks) Len() int {
	return len(this)
}
func (this SubShareNotebooks) Less(i, j int) bool {
	return this[i].ShareNotebook.Seq < this[j].ShareNotebook.Seq
}
func (this SubShareNotebooks) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 以用户进行分组
// type ShareNotebooksByUsers []ShareNotebooksByUser
type ShareNotebooksByUser map[string][]ShareNotebooks

/*
type ShareNotebooksByUser struct {
//	User
	UserId int64
	ShareNotebooks []ShareNotebooks // SubShareNotebooks 一样的, 不过用[]更容易理解
}
*/

//----------------------------------

// 唯一: userId-ToUserId-NoteId
// use leanote
// db.leanote.share_notes.ensureIndex({"UserId":1,"ToUserId":1, "NoteId": 1},{"unique":true})
type ShareNote struct {
	ShareNoteId int64 `xorm:"pk"` // 必须要设置xorm:"pk" 不然mgo不会认为是主键
	UserId      int64
	ToUserId    int64
	ToGroupId   int64 // 分享给的用户组
	ToGroup     Group // 仅仅为了显示, 不存储, 分组信息
	NoteId      int64
	Perm        int       `xorm:"Perm"` // 权限, 0只读, 1可修改
	CreatedTime time.Time `xorm:"created"`
}

// 谁共享给了谁note
// 共享了note, notebook都要加!
// 唯一: UserId-ToUserId
// db.leanote.has_share_notes.ensureIndex({"UserId":1,"ToUserId":1},{"unique":true})
type HasShareNote struct {
	HasShareNotebookId int64 `xorm:"pk"` // 必须要设置xorm:"pk" 不然mgo不会认为是主键
	UserId             int64
	ToUserId           int64
	Seq                int // 以后还可以用户排序
}

// 将note与权限结合起来
type ShareNoteWithPerm struct {
	Note
	Perm int
}

// 查看共享状态, 用户的信息
type ShareUserInfo struct {
	ToUserId          int64
	Email             string
	Perm              int  // note或其notebook的权限
	NotebookHasShared bool // 是否其notebook共享了
}
