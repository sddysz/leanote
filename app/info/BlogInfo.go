package info

import (
	"time"
)

// 只为blog, 不为note

type BlogItem struct {
	Note
	Abstract string
	Content  string // 可能是content的一部分, 截取. 点击more后就是整个信息了
	HasMore  bool   // 是否是否还有
	User     User   // 用户信息
}

type UserBlogBase struct {
	Logo     string
	Title    string // 标题
	SubTitle string // 副标题
	//	AboutMe  string `AboutMe`  // 关于我
}

type UserBlogComment struct {
	CanComment  bool   // 是否可以评论
	CommentType string // default 或 disqus
	DisqusId    string
}

type UserBlogStyle struct {
	Style string // 风格
	Css   string // 自定义css
}

// 每个用户一份博客设置信息
type UserBlog struct {
	UserId   int64 `xorm:"pk"` // 谁的
	Logo     string
	Title    string // 标题
	SubTitle string // 副标题
	AboutMe  string // 关于我, 弃用

	CanComment bool // 是否可以评论

	CommentType string // default 或 disqus
	DisqusId    string

	Style string // 风格
	Css   string // 自定义css

	ThemeId   int64  // 主题Id
	ThemePath string // 不存值, 从Theme中获取, 相对路径 public/

	CateIds []string            // 分类Id, 排序好的
	Singles []map[string]string // 单页, 排序好的, map包含: ["Title"], ["SingleId"]

	PerPageSize int
	SortField   string // 排序字段
	IsAsc       bool   // 排序类型, 降序, 升序, 默认是false, 表示降序

	SubDomain string // 二级域名
	Domain    string // 自定义域名

}

// 博客统计信息
type BlogStat struct {
	NoteId     int64 `xorm:"pk"`
	ReadNum    int64 // 阅读次数 2014/9/28
	LikeNum    int64 // 点赞次数 2014/9/28
	CommentNum int64 // 评论次数 2014/9/28
}

// 单页
type BlogSingle struct {
	SingleId    int64 `xorm:"pk"`
	UserId      int64
	Title       string
	UrlTitle    string // 2014/11/11
	Content     string
	UpdatedTime time.Time `xorm:"updated"`
	CreatedTime time.Time `xorm:"created"`
}

//------------------------
// 社交功能, 点赞, 分享, 评论

// 点赞记录
type BlogLike struct {
	LikeId      int64 `xorm:"pk"`
	NoteId      int64
	UserId      int64
	CreatedTime time.Time `xorm:"created"`
}

// 评论
type BlogComment struct {
	CommentId int64 `xorm:"pk"`
	NoteId    int64

	UserId  int64  // UserId回复ToUserId
	Content string `xorm:"text"` // 评论内容

	ToCommentId int64 // 对某条评论进行回复
	ToUserId    int64 // 为空表示直接评论, 不回空表示回复某人

	LikeNum     int      // 点赞次数, 评论也可以点赞
	LikeUserIds []string // 点赞的用户ids

	CreatedTime time.Time `xorm:"created"`
}

type BlogCommentPublic struct {
	BlogComment
	IsILikeIt bool
}

type BlogUrls struct {
	IndexUrl    string
	CateUrl     string
	SearchUrl   string
	SingleUrl   string
	PostUrl     string
	ArchiveUrl  string
	TagsUrl     string
	TagPostsUrl string
}
