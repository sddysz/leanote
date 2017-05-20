package info

import (
	"time"
)

// Attach belongs to note
type Attach struct {
	AttachId     int64     `xorm:"pk"` //
	NoteId       int64     //
	UploadUserId int64     // 可以不是note owner, 协作者userId
	Name         string    // file name, md5, such as 13232312.doc
	Title        string    // raw file name
	Size         int64     // file size (byte)
	Type         string    // file type, "doc" = word
	Path         string    // the file path such as: files/userId/attachs/adfadf.doc
	CreatedTime  time.Time `xorm:"created"`

	// FromFileId int64 `bson:"FromFileId,omitempty"` // copy from fileId, for collaboration
}
