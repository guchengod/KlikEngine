// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoHashtag is the golang structure of table hg_video_hashtag for DAO operations like Where/Data.
type VideoHashtag struct {
	g.Meta    `orm:"table:hg_video_hashtag, do:true"`
	Id        interface{} // 自增主键
	AwemeId   interface{} // 视频ID，关联hg_videos表
	HashtagId interface{} // 话题ID，关联hg_hashtags表
	CreatedAt *gtime.Time // 关联关系创建时间
}
