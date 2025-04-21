// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Musics is the golang structure of table hg_musics for DAO operations like Where/Data.
type Musics struct {
	g.Meta        `orm:"table:hg_musics, do:true"`
	MusicId       interface{} // 音乐唯一ID，主键
	Title         interface{} // 音乐标题
	Author        interface{} // 音乐作者
	CoverUrl      interface{} // 音乐封面URL
	PlayUrl       interface{} // 音乐播放URL
	Duration      interface{} // 音乐时长，单位秒
	IsOriginal    interface{} // 是否原创：0-非原创 1-原创
	UserCount     interface{} // 使用人数统计
	OwnerId       interface{} // 音乐所有者ID
	OwnerNickname interface{} // 音乐所有者昵称
	CreatedAt     *gtime.Time // 记录创建时间
	UpdatedAt     *gtime.Time // 记录更新时间
}
