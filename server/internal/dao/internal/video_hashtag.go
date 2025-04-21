// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoHashtagDao is the data access object for the table hg_video_hashtag.
type VideoHashtagDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  VideoHashtagColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// VideoHashtagColumns defines and stores column names for the table hg_video_hashtag.
type VideoHashtagColumns struct {
	Id        string // 自增主键
	AwemeId   string // 视频ID，关联hg_videos表
	HashtagId string // 话题ID，关联hg_hashtags表
	CreatedAt string // 关联关系创建时间
}

// videoHashtagColumns holds the columns for the table hg_video_hashtag.
var videoHashtagColumns = VideoHashtagColumns{
	Id:        "id",
	AwemeId:   "aweme_id",
	HashtagId: "hashtag_id",
	CreatedAt: "created_at",
}

// NewVideoHashtagDao creates and returns a new DAO object for table data access.
func NewVideoHashtagDao(handlers ...gdb.ModelHandler) *VideoHashtagDao {
	return &VideoHashtagDao{
		group:    "default",
		table:    "hg_video_hashtag",
		columns:  videoHashtagColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VideoHashtagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VideoHashtagDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VideoHashtagDao) Columns() VideoHashtagColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VideoHashtagDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VideoHashtagDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VideoHashtagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
