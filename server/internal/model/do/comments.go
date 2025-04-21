// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure of table hg_comments for DAO operations like Where/Data.
type Comments struct {
	g.Meta     `orm:"table:hg_comments, do:true"`
	CommentId  interface{} // 评论ID，主键
	AwemeId    interface{} // 所属视频ID，关联hg_videos表
	UserId     interface{} // 评论用户ID，关联hg_admin_member表
	Content    interface{} // 评论内容
	DiggCount  interface{} // 评论点赞数
	ReplyCount interface{} // 回复数
	CreatedAt  *gtime.Time // 评论创建时间
	IsAuthor   interface{} // 是否作者回复：0-否 1-是
	IsDelete   interface{} // 是否删除：0-正常 1-已删除
}
