// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HashtagsDao is the data access object for the table hg_hashtags.
type HashtagsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HashtagsColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HashtagsColumns defines and stores column names for the table hg_hashtags.
type HashtagsColumns struct {
	HashtagId string // 话题标签ID，自增主键
	Name      string // 话题标签名称
	ViewCount string // 话题浏览次数
	UseCount  string // 话题被使用次数
	CreatedAt string // 话题创建时间
}

// hashtagsColumns holds the columns for the table hg_hashtags.
var hashtagsColumns = HashtagsColumns{
	HashtagId: "hashtag_id",
	Name:      "name",
	ViewCount: "view_count",
	UseCount:  "use_count",
	CreatedAt: "created_at",
}

// NewHashtagsDao creates and returns a new DAO object for table data access.
func NewHashtagsDao(handlers ...gdb.ModelHandler) *HashtagsDao {
	return &HashtagsDao{
		group:    "default",
		table:    "hg_hashtags",
		columns:  hashtagsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HashtagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HashtagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HashtagsDao) Columns() HashtagsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HashtagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HashtagsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HashtagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
