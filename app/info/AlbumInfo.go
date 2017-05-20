package info

import (
	"time"
)

type Album struct {
	AlbumId     int64 `xorm:"pk"` //
	UserId      int64
	Name        string // album name
	Type        int    // type, the default is image: 0
	Seq         int
	CreatedTime time.Time `xorm:"Created"`
}
