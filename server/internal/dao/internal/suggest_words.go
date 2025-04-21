// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SuggestWordsDao is the data access object for the table hg_suggest_words.
type SuggestWordsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SuggestWordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SuggestWordsColumns defines and stores column names for the table hg_suggest_words.
type SuggestWordsColumns struct {
	WordId    string // 关键词ID，主键
	Word      string // 关键词内容
	Scene     string // 出现场景：comment_top_rec-评论顶部推荐 feed_bottom_rec-信息流底部推荐
	HintText  string // 提示文本，如"大家都在搜："
	CreatedAt string // 记录创建时间
	UpdatedAt string // 记录更新时间
}

// suggestWordsColumns holds the columns for the table hg_suggest_words.
var suggestWordsColumns = SuggestWordsColumns{
	WordId:    "word_id",
	Word:      "word",
	Scene:     "scene",
	HintText:  "hint_text",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSuggestWordsDao creates and returns a new DAO object for table data access.
func NewSuggestWordsDao(handlers ...gdb.ModelHandler) *SuggestWordsDao {
	return &SuggestWordsDao{
		group:    "default",
		table:    "hg_suggest_words",
		columns:  suggestWordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SuggestWordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SuggestWordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SuggestWordsDao) Columns() SuggestWordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SuggestWordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SuggestWordsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SuggestWordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
