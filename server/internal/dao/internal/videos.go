// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VideosDao is the data access object for the table hg_videos.
type VideosDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  VideosColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// VideosColumns defines and stores column names for the table hg_videos.
type VideosColumns struct {
	AwemeId         string // 抖音视频唯一ID，主键
	UserId          string // 发布视频的用户ID，关联hg_admin_member表
	Description     string // 视频描述/文案内容
	CreatedAt       string // 视频创建时间（抖音原始时间戳）
	Duration        string // 视频时长，单位毫秒
	ShareUrl        string // 视频分享链接
	DiggCount       string // 点赞数
	CommentCount    string // 评论数
	CollectCount    string // 收藏数
	ShareCount      string // 分享数
	PlayCount       string // 播放数
	IsTop           string // 是否置顶：0-否 1-是
	IsDelete        string // 是否删除：0-正常 1-已删除
	AllowShare      string // 是否允许分享：0-不允许 1-允许
	IsProhibited    string // 是否被禁止：0-正常 1-被禁止
	PreventDownload string // 是否禁止下载：0-允许 1-禁止
	Width           string // 视频宽度（像素）
	Height          string // 视频高度（像素）
	Ratio           string // 视频比例，如1080p,720p等
	VideoUri        string // 视频资源URI
	VideoUrl        string // 视频播放URL
	CoverUri        string // 封面图URI
	CoverUrl        string // 封面图URL
	CreatedAtstamp  string // 系统记录创建时间
	UpdatedAtstamp  string // 系统记录更新时间
}

// videosColumns holds the columns for the table hg_videos.
var videosColumns = VideosColumns{
	AwemeId:         "aweme_id",
	UserId:          "user_id",
	Description:     "description",
	CreatedAt:       "created_at",
	Duration:        "duration",
	ShareUrl:        "share_url",
	DiggCount:       "digg_count",
	CommentCount:    "comment_count",
	CollectCount:    "collect_count",
	ShareCount:      "share_count",
	PlayCount:       "play_count",
	IsTop:           "is_top",
	IsDelete:        "is_delete",
	AllowShare:      "allow_share",
	IsProhibited:    "is_prohibited",
	PreventDownload: "prevent_download",
	Width:           "width",
	Height:          "height",
	Ratio:           "ratio",
	VideoUri:        "video_uri",
	VideoUrl:        "video_url",
	CoverUri:        "cover_uri",
	CoverUrl:        "cover_url",
	CreatedAtstamp:  "created_atstamp",
	UpdatedAtstamp:  "updated_atstamp",
}

// NewVideosDao creates and returns a new DAO object for table data access.
func NewVideosDao(handlers ...gdb.ModelHandler) *VideosDao {
	return &VideosDao{
		group:    "default",
		table:    "hg_videos",
		columns:  videosColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VideosDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VideosDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VideosDao) Columns() VideosColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VideosDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VideosDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *VideosDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
