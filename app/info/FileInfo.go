package info

import (
	"time"
)

type File struct {
	FileId         int64 `xorm:"pk"` //
	UserId         int64
	AlbumId        int64
	Name           string // file name
	Title          string // file name or user defined for search
	Size           int64  // file size (byte)
	Type           string // file type, "" = image, "doc" = word
	Path           string // the file path
	IsDefaultAlbum bool
	CreatedTime    time.Time `xorm:"created"`

	FromFileId int64 // copy from fileId, for collaboration
}
