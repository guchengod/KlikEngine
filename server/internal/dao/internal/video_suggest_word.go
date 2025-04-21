// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideoSuggestWordDao is the data access object for the table hg_video_suggest_word.
type VideoSuggestWordDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  VideoSuggestWordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// VideoSuggestWordColumns defines and stores column names for the table hg_video_suggest_word.
type VideoSuggestWordColumns struct {
	Id        string // 自增主键
	AwemeId   string // 视频ID，关联hg_videos表
	WordId    string // 关键词ID，关联hg_suggest_words表
	CreatedAt string // 关联关系创建时间
}

// videoSuggestWordColumns holds the columns for the table hg_video_suggest_word.
var videoSuggestWordColumns = VideoSuggestWordColumns{
	Id:        "id",
	AwemeId:   "aweme_id",
	WordId:    "word_id",
	CreatedAt: "created_at",
}

// NewVideoSuggestWordDao creates and returns a new DAO object for table data access.
func NewVideoSuggestWordDao(handlers ...gdb.ModelHandler) *VideoSuggestWordDao {
	return &VideoSuggestWordDao{
		group:    "default",
		table:    "hg_video_suggest_word",
		columns:  videoSuggestWordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VideoSuggestWordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VideoSuggestWordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VideoSuggestWordDao) Columns() VideoSuggestWordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VideoSuggestWordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VideoSuggestWordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VideoSuggestWordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
