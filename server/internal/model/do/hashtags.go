// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hashtags is the golang structure of table hg_hashtags for DAO operations like Where/Data.
type Hashtags struct {
	g.Meta    `orm:"table:hg_hashtags, do:true"`
	HashtagId interface{} // 话题标签ID，自增主键
	Name      interface{} // 话题标签名称
	ViewCount interface{} // 话题浏览次数
	UseCount  interface{} // 话题被使用次数
	CreatedAt *gtime.Time // 话题创建时间
}
