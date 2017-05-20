package info

// 建议
type Suggestion struct {
	Id         int64 `xorm:"pk"`
	UserId     int64
	Addr       string
	Suggestion string
}
