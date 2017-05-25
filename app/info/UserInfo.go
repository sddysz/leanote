package info

import (
	"time"
)

// 第三方类型
const (
	ThirdGithub = iota
	ThirdQQ
)

type User struct {
	UserId      int64  `xorm:"pk"`                          // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Email       string `xorm:"varchar(225) notnull unique"` // 全是小写
	Verified    bool   // Email是否已验证过?
	Username    string // 不区分大小写, 全是小写
	UsernameRaw string // 可能有大小写
	Pwd         string
	CreatedTime time.Time `xorm:"created"`

	Logo string // 9-24
	// 主题
	Theme string

	// 用户配置
	NotebookWidth int  // 笔记本宽度
	NoteListWidth int  // 笔记列表宽度
	MdEditorWidth int  // markdown 左侧编辑器宽度
	LeftIsMin     bool // 左侧是否是隐藏的, 默认是打开的

	// 这里 第三方登录
	ThirdUserId   string // 用户Id, 在第三方中唯一可识别
	ThirdUsername string // 第三方中username, 为了显示
	ThirdType     int    // 第三方类型

	// 用户的帐户类型

	ImageNum   int   // 图片数量
	ImageSize  int   // 图片大小
	AttachNum  int   // 附件数量
	AttachSize int   // 附件大小
	FromUserId int64 // 邀请的用户

	AccountType      string    // normal(为空), premium
	AccountStartTime time.Time // 开始日期
	AccountEndTime   time.Time // 结束日期
	// 阈值
	MaxImageNum      int // 图片数量
	MaxImageSize     int // 图片大小
	MaxAttachNum     int // 图片数量
	MaxAttachSize    int // 图片大小
	MaxPerAttachSize int // 单个附件大小

	// 2015/1/15, 更新序号
	Usn            int64     // UpdateSequenceNum , 全局的
	FullSyncBefore time.Time // 需要全量同步的时间, 如果 > 客户端的LastSyncTime, 则需要全量更新
}

type UserAccount struct {
	AccountType      string    // normal(为空), premium
	AccountStartTime time.Time // 开始日期
	AccountEndTime   time.Time // 结束日期
	// 阈值
	MaxImageNum      int // 图片数量
	MaxImageSize     int // 图片大小
	MaxAttachNum     int // 图片数量
	MaxAttachSize    int // 图片大小
	MaxPerAttachSize int // 单个附件大小
}

// note主页需要
type UserAndBlogUrl struct {
	User
	BlogUrl string
	PostUrl string
}

// 用户与博客信息结合, 公开
type UserAndBlog struct {
	UserId    int64  // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Email     string // 全是小写
	Username  string // 不区分大小写, 全是小写
	Logo      string
	BlogTitle string // 博客标题
	BlogLogo  string // 博客Logo
	BlogUrl   string // 博客链接, 主页

	BlogUrls // 各个页面
}
