// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoMusic is the golang structure for table video_music.
type VideoMusic struct {
	Id        int         `json:"id"        orm:"id"         description:"自增主键"`
	AwemeId   int64       `json:"awemeId"   orm:"aweme_id"   description:"视频ID，关联hg_videos表"`
	MusicId   int64       `json:"musicId"   orm:"music_id"   description:"音乐ID，关联hg_musics表"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"关联关系创建时间"`
}
