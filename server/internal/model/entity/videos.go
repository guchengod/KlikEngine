// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Videos is the golang structure for table videos.
type Videos struct {
	AwemeId         int64       `json:"awemeId"         orm:"aweme_id"         description:"抖音视频唯一ID，主键"`
	UserId          int64       `json:"userId"          orm:"user_id"          description:"发布视频的用户ID，关联hg_admin_member表"`
	Description     string      `json:"description"     orm:"description"      description:"视频描述/文案内容"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"视频创建时间（抖音原始时间戳）"`
	Duration        int         `json:"duration"        orm:"duration"         description:"视频时长，单位毫秒"`
	ShareUrl        string      `json:"shareUrl"        orm:"share_url"        description:"视频分享链接"`
	DiggCount       int         `json:"diggCount"       orm:"digg_count"       description:"点赞数"`
	CommentCount    int         `json:"commentCount"    orm:"comment_count"    description:"评论数"`
	CollectCount    int         `json:"collectCount"    orm:"collect_count"    description:"收藏数"`
	ShareCount      int         `json:"shareCount"      orm:"share_count"      description:"分享数"`
	PlayCount       int         `json:"playCount"       orm:"play_count"       description:"播放数"`
	IsTop           int         `json:"isTop"           orm:"is_top"           description:"是否置顶：0-否 1-是"`
	IsDelete        int         `json:"isDelete"        orm:"is_delete"        description:"是否删除：0-正常 1-已删除"`
	AllowShare      int         `json:"allowShare"      orm:"allow_share"      description:"是否允许分享：0-不允许 1-允许"`
	IsProhibited    int         `json:"isProhibited"    orm:"is_prohibited"    description:"是否被禁止：0-正常 1-被禁止"`
	PreventDownload int         `json:"preventDownload" orm:"prevent_download" description:"是否禁止下载：0-允许 1-禁止"`
	Width           int         `json:"width"           orm:"width"            description:"视频宽度（像素）"`
	Height          int         `json:"height"          orm:"height"           description:"视频高度（像素）"`
	Ratio           string      `json:"ratio"           orm:"ratio"            description:"视频比例，如1080p,720p等"`
	VideoUri        string      `json:"videoUri"        orm:"video_uri"        description:"视频资源URI"`
	VideoUrl        string      `json:"videoUrl"        orm:"video_url"        description:"视频播放URL"`
	CoverUri        string      `json:"coverUri"        orm:"cover_uri"        description:"封面图URI"`
	CoverUrl        string      `json:"coverUrl"        orm:"cover_url"        description:"封面图URL"`
	CreatedAtstamp  *gtime.Time `json:"createdAtstamp"  orm:"created_atstamp"  description:"系统记录创建时间"`
	UpdatedAtstamp  *gtime.Time `json:"updatedAtstamp"  orm:"updated_atstamp"  description:"系统记录更新时间"`
}
