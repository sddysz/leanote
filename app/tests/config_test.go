package tests

import (
	"github.com/revel/revel"
	"github.com/sddysz/leanote/app/db"
	"testing"
	//	. "github.com/sddysz/leanote/app/lea"
	"github.com/sddysz/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	revel.Init("dev", "github.com/sddysz/leanote", "/Users/life/Documents/Go/package_base/src")
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
