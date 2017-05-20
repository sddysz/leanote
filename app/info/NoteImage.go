package info

// 笔记内部图片
type NoteImage struct {
	NoteImageId int64 `xorm:"pk"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	NoteId      int64 // 笔记
	ImageId     int64 // 图片fileId
}
