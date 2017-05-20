package info

import (
	"time"
)

// 配置, 每一个配置一行记录
type Config struct {
	ConfigId    int64 `xorm:"pk"`
	UserId      int64
	Key         string
	ValueStr    string              // "1"
	ValueArr    []string            // ["1","b","c"]
	ValueMap    map[string]string   // {"a":"bb", "CC":"xx"}
	ValueArrMap []map[string]string // [{"a":"B"}, {}, {}]
	IsArr       bool                // 是否是数组
	IsMap       bool                // 是否是Map
	IsArrMap    bool                // 是否是数组Map

	// StringConfigs map[string]string   `StringConfigs` // key => value
	// ArrayConfigs  map[string][]string `ArrayConfigs`  // key => []value

	UpdatedTime time.Time `xorm:"updated"`
}
