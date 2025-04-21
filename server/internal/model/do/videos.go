// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Videos is the golang structure of table hg_videos for DAO operations like Where/Data.
type Videos struct {
	g.Meta          `orm:"table:hg_videos, do:true"`
	AwemeId         interface{} // 抖音视频唯一ID，主键
	UserId          interface{} // 发布视频的用户ID，关联hg_admin_member表
	Description     interface{} // 视频描述/文案内容
	CreatedAt       *gtime.Time // 视频创建时间（抖音原始时间戳）
	Duration        interface{} // 视频时长，单位毫秒
	ShareUrl        interface{} // 视频分享链接
	DiggCount       interface{} // 点赞数
	CommentCount    interface{} // 评论数
	CollectCount    interface{} // 收藏数
	ShareCount      interface{} // 分享数
	PlayCount       interface{} // 播放数
	IsTop           interface{} // 是否置顶：0-否 1-是
	IsDelete        interface{} // 是否删除：0-正常 1-已删除
	AllowShare      interface{} // 是否允许分享：0-不允许 1-允许
	IsProhibited    interface{} // 是否被禁止：0-正常 1-被禁止
	PreventDownload interface{} // 是否禁止下载：0-允许 1-禁止
	Width           interface{} // 视频宽度（像素）
	Height          interface{} // 视频高度（像素）
	Ratio           interface{} // 视频比例，如1080p,720p等
	VideoUri        interface{} // 视频资源URI
	VideoUrl        interface{} // 视频播放URL
	CoverUri        interface{} // 封面图URI
	CoverUrl        interface{} // 封面图URL
	CreatedAtstamp  *gtime.Time // 系统记录创建时间
	UpdatedAtstamp  *gtime.Time // 系统记录更新时间
}
