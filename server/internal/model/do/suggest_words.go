// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SuggestWords is the golang structure of table hg_suggest_words for DAO operations like Where/Data.
type SuggestWords struct {
	g.Meta    `orm:"table:hg_suggest_words, do:true"`
	WordId    interface{} // 关键词ID，主键
	Word      interface{} // 关键词内容
	Scene     interface{} // 出现场景：comment_top_rec-评论顶部推荐 feed_bottom_rec-信息流底部推荐
	HintText  interface{} // 提示文本，如"大家都在搜："
	CreatedAt *gtime.Time // 记录创建时间
	UpdatedAt *gtime.Time // 记录更新时间
}
