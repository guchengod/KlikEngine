// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hashtags is the golang structure for table hashtags.
type Hashtags struct {
	HashtagId int64       `json:"hashtagId" orm:"hashtag_id" description:"话题标签ID，自增主键"`
	Name      string      `json:"name"      orm:"name"       description:"话题标签名称"`
	ViewCount int64       `json:"viewCount" orm:"view_count" description:"话题浏览次数"`
	UseCount  int64       `json:"useCount"  orm:"use_count"  description:"话题被使用次数"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"话题创建时间"`
}
