package info

import (
	"time"
)

// 仅仅为了博客的主题

type BlogInfoCustom struct {
	UserId      string
	Username    string
	UserLogo    string
	Title       string
	SubTitle    string
	Logo        string
	OpenComment bool
	CommentType string
	ThemeId     string
	SubDomain   string
	Domain      string
}

type Post struct {
	NoteId      int64
	Title       string
	UrlTitle    string
	ImgSrc      string
	CreatedTime time.Time
	UpdatedTime time.Time
	PublicTime  time.Time
	Desc        string
	Abstract    string
	Content     string
	Tags        []string
	CommentNum  int64
	ReadNum     int64
	LikeNum     int64
	IsMarkdown  bool
}

// 归档
type ArchiveMonth struct {
	Month int
	Posts []*Post
}
type Archive struct {
	Year         int
	MonthAchives []ArchiveMonth
	Posts        []*Post
}

type Cate struct {
	CateId       string
	ParentCateId string
	Title        string
	UrlTitle     string
	Children     []*Cate
}
