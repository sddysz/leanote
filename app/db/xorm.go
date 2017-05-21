package db

import (
	"github.com/go-xorm/xorm"
)

var (
	Engine        *xorm.Engine
	CachePageSize int // 允许缓存前几页数据
)

// init 初始化数据库引擎
func Init() {
	//c := revel.Config
	// driver, _ := c.String("db.driver")
	// dbname, _ := c.String("db.dbname")
	// user, _ := c.String("db.username")
	// password, _ := c.String("db.password")
	// host, _ := c.String("db.host")
	// params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)
	Engine, err := xorm.NewEngine("sqlite3", "./data.db")
	// defer Engine.Close()
	if err != nil {
		panic(err)
	}

	// Engine.ShowSQL = revel.DevMode

	err = Engine.Sync()

	if err != nil {
		panic(err)
	}

}
