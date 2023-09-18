package dto

import (
	"time"
)

type ArticleDigForm struct {
	ArticleID uint64 `form:"article_id" json:"article_id,omitempty"` // 文章ID
	Views     uint64 `form:"views" json:"views,omitempty"`           // 浏览量
	Likes     uint64 `form:"likes" json:"likes,omitempty"`           // 喜欢
	Dislikes  uint64 `form:"dislikes" json:"dislikes,omitempty"`     // 不喜欢
}

type ArticleDigListQueryFilter struct {
	ArticleID *uint64 `query:"article_id" json:"article_id,omitempty"` // 文章ID
	Views     *uint64 `query:"views" json:"views,omitempty"`           // 浏览量
	Likes     *uint64 `query:"likes" json:"likes,omitempty"`           // 喜欢
	Dislikes  *uint64 `query:"dislikes" json:"dislikes,omitempty"`     // 不喜欢
}

type ArticleDigItem struct {
	ID        uint64    `json:"id,omitempty"`         // ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	ArticleID uint64    `json:"article_id,omitempty"` // 文章ID
	Views     uint64    `json:"views,omitempty"`      // 浏览量
	Likes     uint64    `json:"likes,omitempty"`      // 喜欢
	Dislikes  uint64    `json:"dislikes,omitempty"`   // 不喜欢
}
