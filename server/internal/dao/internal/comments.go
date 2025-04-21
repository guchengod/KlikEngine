// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentsDao is the data access object for the table hg_comments.
type CommentsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CommentsColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CommentsColumns defines and stores column names for the table hg_comments.
type CommentsColumns struct {
	CommentId  string // 评论ID，主键
	AwemeId    string // 所属视频ID，关联hg_videos表
	UserId     string // 评论用户ID，关联hg_admin_member表
	Content    string // 评论内容
	DiggCount  string // 评论点赞数
	ReplyCount string // 回复数
	CreatedAt  string // 评论创建时间
	IsAuthor   string // 是否作者回复：0-否 1-是
	IsDelete   string // 是否删除：0-正常 1-已删除
}

// commentsColumns holds the columns for the table hg_comments.
var commentsColumns = CommentsColumns{
	CommentId:  "comment_id",
	AwemeId:    "aweme_id",
	UserId:     "user_id",
	Content:    "content",
	DiggCount:  "digg_count",
	ReplyCount: "reply_count",
	CreatedAt:  "created_at",
	IsAuthor:   "is_author",
	IsDelete:   "is_delete",
}

// NewCommentsDao creates and returns a new DAO object for table data access.
func NewCommentsDao(handlers ...gdb.ModelHandler) *CommentsDao {
	return &CommentsDao{
		group:    "default",
		table:    "hg_comments",
		columns:  commentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommentsDao) Columns() CommentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommentsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
