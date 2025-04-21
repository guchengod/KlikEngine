// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoMusic is the golang structure of table hg_video_music for DAO operations like Where/Data.
type VideoMusic struct {
	g.Meta    `orm:"table:hg_video_music, do:true"`
	Id        interface{} // 自增主键
	AwemeId   interface{} // 视频ID，关联hg_videos表
	MusicId   interface{} // 音乐ID，关联hg_musics表
	CreatedAt *gtime.Time // 关联关系创建时间
}
