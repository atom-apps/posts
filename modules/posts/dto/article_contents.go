package dto

import (
	"time"
)

type ArticleContentForm struct {
	TenantID     uint64 `form:"tenant_id" json:"tenant_id,omitempty"`         // 租户ID
	UserID       uint64 `form:"user_id" json:"user_id,omitempty"`             // 用户ID
	ArticleID    uint64 `form:"article_id" json:"article_id,omitempty"`       // 文章ID
	FreeContent  string `form:"free_content" json:"free_content,omitempty"`   // 内容
	PriceContent string `form:"price_content" json:"price_content,omitempty"` // 付费内容
}

type ArticleContentListQueryFilter struct {
	TenantID     *uint64 `query:"tenant_id" json:"tenant_id,omitempty"`         // 租户ID
	UserID       *uint64 `query:"user_id" json:"user_id,omitempty"`             // 用户ID
	ArticleID    *uint64 `query:"article_id" json:"article_id,omitempty"`       // 文章ID
	FreeContent  *string `query:"free_content" json:"free_content,omitempty"`   // 内容
	PriceContent *string `query:"price_content" json:"price_content,omitempty"` // 付费内容
}

type ArticleContentItem struct {
	ID           uint64    `json:"id,omitempty"`            // ID
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at,omitempty"`    // 更新时间
	TenantID     uint64    `json:"tenant_id,omitempty"`     // 租户ID
	UserID       uint64    `json:"user_id,omitempty"`       // 用户ID
	ArticleID    uint64    `json:"article_id,omitempty"`    // 文章ID
	FreeContent  string    `json:"free_content,omitempty"`  // 内容
	PriceContent string    `json:"price_content,omitempty"` // 付费内容
}
