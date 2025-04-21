// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SuggestWords is the golang structure for table suggest_words.
type SuggestWords struct {
	WordId    int64       `json:"wordId"    orm:"word_id"    description:"关键词ID，主键"`
	Word      string      `json:"word"      orm:"word"       description:"关键词内容"`
	Scene     string      `json:"scene"     orm:"scene"      description:"出现场景：comment_top_rec-评论顶部推荐 feed_bottom_rec-信息流底部推荐"`
	HintText  string      `json:"hintText"  orm:"hint_text"  description:"提示文本，如\"大家都在搜：\""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"记录创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"记录更新时间"`
}
