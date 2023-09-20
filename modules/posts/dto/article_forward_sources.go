package dto

import (
	"time"
)

type ArticleForwardSourceForm struct {
	ArticleID    uint64 `form:"article_id" json:"article_id,omitempty"`       // 文章ID
	Source       string `form:"source" json:"source,omitempty"`               // 原文地址
	SourceAuthor string `form:"source_author" json:"source_author,omitempty"` // 原文作者
}

type ArticleForwardSourceListQueryFilter struct {
	ArticleID    *uint64 `query:"article_id" json:"article_id,omitempty"`       // 文章ID
	Source       *string `query:"source" json:"source,omitempty"`               // 原文地址
	SourceAuthor *string `query:"source_author" json:"source_author,omitempty"` // 原文作者
}

type ArticleForwardSourceItem struct {
	ID           uint64    `json:"id,omitempty"`            // ID
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 创建时间
	ArticleID    uint64    `json:"article_id,omitempty"`    // 文章ID
	Source       string    `json:"source,omitempty"`        // 原文地址
	SourceAuthor string    `json:"source_author,omitempty"` // 原文作者
}
