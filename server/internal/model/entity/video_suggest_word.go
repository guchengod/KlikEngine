// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoSuggestWord is the golang structure for table video_suggest_word.
type VideoSuggestWord struct {
	Id        int         `json:"id"        orm:"id"         description:"自增主键"`
	AwemeId   int64       `json:"awemeId"   orm:"aweme_id"   description:"视频ID，关联hg_videos表"`
	WordId    int64       `json:"wordId"    orm:"word_id"    description:"关键词ID，关联hg_suggest_words表"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"关联关系创建时间"`
}
