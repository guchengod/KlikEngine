// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MusicsDao is the data access object for the table hg_musics.
type MusicsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MusicsColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MusicsColumns defines and stores column names for the table hg_musics.
type MusicsColumns struct {
	MusicId       string // 音乐唯一ID，主键
	Title         string // 音乐标题
	Author        string // 音乐作者
	CoverUrl      string // 音乐封面URL
	PlayUrl       string // 音乐播放URL
	Duration      string // 音乐时长，单位秒
	IsOriginal    string // 是否原创：0-非原创 1-原创
	UserCount     string // 使用人数统计
	OwnerId       string // 音乐所有者ID
	OwnerNickname string // 音乐所有者昵称
	CreatedAt     string // 记录创建时间
	UpdatedAt     string // 记录更新时间
}

// musicsColumns holds the columns for the table hg_musics.
var musicsColumns = MusicsColumns{
	MusicId:       "music_id",
	Title:         "title",
	Author:        "author",
	CoverUrl:      "cover_url",
	PlayUrl:       "play_url",
	Duration:      "duration",
	IsOriginal:    "is_original",
	UserCount:     "user_count",
	OwnerId:       "owner_id",
	OwnerNickname: "owner_nickname",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewMusicsDao creates and returns a new DAO object for table data access.
func NewMusicsDao(handlers ...gdb.ModelHandler) *MusicsDao {
	return &MusicsDao{
		group:    "default",
		table:    "hg_musics",
		columns:  musicsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MusicsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MusicsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MusicsDao) Columns() MusicsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MusicsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MusicsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MusicsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
