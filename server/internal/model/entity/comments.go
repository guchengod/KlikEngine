// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure for table comments.
type Comments struct {
	CommentId  int64       `json:"commentId"  orm:"comment_id"  description:"评论ID，主键"`
	AwemeId    int64       `json:"awemeId"    orm:"aweme_id"    description:"所属视频ID，关联hg_videos表"`
	UserId     int64       `json:"userId"     orm:"user_id"     description:"评论用户ID，关联hg_admin_member表"`
	Content    string      `json:"content"    orm:"content"     description:"评论内容"`
	DiggCount  int         `json:"diggCount"  orm:"digg_count"  description:"评论点赞数"`
	ReplyCount int         `json:"replyCount" orm:"reply_count" description:"回复数"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"评论创建时间"`
	IsAuthor   int         `json:"isAuthor"   orm:"is_author"   description:"是否作者回复：0-否 1-是"`
	IsDelete   int         `json:"isDelete"   orm:"is_delete"   description:"是否删除：0-正常 1-已删除"`
}
