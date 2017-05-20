package info

import (
	"time"
)

// 主题, 每个用户有多个主题, 这里面有主题的配置信息
// 模板, css, js, images, 都在路径Path下
type Theme struct {
	ThemeId   int64 `xorm:"pk"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	UserId    int64
	Name      string
	Version   string
	Author    string
	AuthorUrl string
	Path      string                 // 文件夹路径, public/upload/54d7620d99c37b030600002c/themes/54d867c799c37b533e000001
	Info      map[string]interface{} // 所有信息
	IsActive  bool                   // 是否在用

	IsDefault bool   // leanote默认主题, 如果用户修改了默认主题, 则先copy之. 也是admin用户的主题
	Style     string // 之前的, 只有default的用户才有blog_default, blog_daqi, blog_left_fixed

	CreatedTime time.Time `xorm:"created"`
	UpdatedTime time.Time `xorm:"updated"`
}
