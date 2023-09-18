package dto

import (
	"time"
)

type ArticleAttachmentForm struct {
	TenantID     uint64 `form:"tenant_id" json:"tenant_id,omitempty"`         // 租户ID
	UserID       uint64 `form:"user_id" json:"user_id,omitempty"`             // 用户ID
	ArticleID    uint64 `form:"article_id" json:"article_id,omitempty"`       // 文章ID
	FilesystemID uint64 `form:"filesystem_id" json:"filesystem_id,omitempty"` // 文件系统ID
	Description  string `form:"description" json:"description,omitempty"`     // 文件描述
	Password     string `form:"password" json:"password,omitempty"`           // 文件密码
}

type ArticleAttachmentListQueryFilter struct {
	TenantID     *uint64 `query:"tenant_id" json:"tenant_id,omitempty"`         // 租户ID
	UserID       *uint64 `query:"user_id" json:"user_id,omitempty"`             // 用户ID
	ArticleID    *uint64 `query:"article_id" json:"article_id,omitempty"`       // 文章ID
	FilesystemID *uint64 `query:"filesystem_id" json:"filesystem_id,omitempty"` // 文件系统ID
	Description  *string `query:"description" json:"description,omitempty"`     // 文件描述
	Password     *string `query:"password" json:"password,omitempty"`           // 文件密码
}

type ArticleAttachmentItem struct {
	ID           uint64    `json:"id,omitempty"`            // ID
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at,omitempty"`    // 更新时间
	TenantID     uint64    `json:"tenant_id,omitempty"`     // 租户ID
	UserID       uint64    `json:"user_id,omitempty"`       // 用户ID
	ArticleID    uint64    `json:"article_id,omitempty"`    // 文章ID
	FilesystemID uint64    `json:"filesystem_id,omitempty"` // 文件系统ID
	Description  string    `json:"description,omitempty"`   // 文件描述
	Password     string    `json:"password,omitempty"`      // 文件密码
}
