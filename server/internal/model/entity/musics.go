// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Musics is the golang structure for table musics.
type Musics struct {
	MusicId       int64       `json:"musicId"       orm:"music_id"       description:"音乐唯一ID，主键"`
	Title         string      `json:"title"         orm:"title"          description:"音乐标题"`
	Author        string      `json:"author"        orm:"author"         description:"音乐作者"`
	CoverUrl      string      `json:"coverUrl"      orm:"cover_url"      description:"音乐封面URL"`
	PlayUrl       string      `json:"playUrl"       orm:"play_url"       description:"音乐播放URL"`
	Duration      int         `json:"duration"      orm:"duration"       description:"音乐时长，单位秒"`
	IsOriginal    int         `json:"isOriginal"    orm:"is_original"    description:"是否原创：0-非原创 1-原创"`
	UserCount     int         `json:"userCount"     orm:"user_count"     description:"使用人数统计"`
	OwnerId       string      `json:"ownerId"       orm:"owner_id"       description:"音乐所有者ID"`
	OwnerNickname string      `json:"ownerNickname" orm:"owner_nickname" description:"音乐所有者昵称"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"记录创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"记录更新时间"`
}
