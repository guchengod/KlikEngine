// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoSuggestWord is the golang structure of table hg_video_suggest_word for DAO operations like Where/Data.
type VideoSuggestWord struct {
	g.Meta    `orm:"table:hg_video_suggest_word, do:true"`
	Id        interface{} // 自增主键
	AwemeId   interface{} // 视频ID，关联hg_videos表
	WordId    interface{} // 关键词ID，关联hg_suggest_words表
	CreatedAt *gtime.Time // 关联关系创建时间
}
