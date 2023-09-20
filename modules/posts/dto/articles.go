package dto

import (
	"time"

	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/common/consts"
)

type ArticleForm struct {
	TenantID    uint64                   `form:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      uint64                   `form:"user_id" json:"user_id,omitempty"`         // 用户ID
	UUID        string                   `form:"uuid" json:"uuid,omitempty"`               // UUID
	BookID      uint64                   `form:"book_id" json:"book_id,omitempty"`         // 书ID
	ChapterID   uint64                   `form:"chapter_id" json:"chapter_id,omitempty"`   // 章节ID
	CategoryID  uint64                   `form:"category_id" json:"category_id,omitempty"` // 分类
	PublishAt   time.Time                `form:"publish_at" json:"publish_at,omitempty"`   // 发布时间
	Type        consts.ArticleType       `form:"type" json:"type,omitempty"`               // 文章类型
	Format      consts.ArticleFormat     `form:"format" json:"format,omitempty"`           // 文章格式
	Title       string                   `form:"title" json:"title,omitempty"`             // 标题
	Keyword     string                   `form:"keyword" json:"keyword,omitempty"`         // 关键词
	Description string                   `form:"description" json:"description,omitempty"` // 简介
	Thumbnails  common.ArticleThumbnails `form:"thumbnails" json:"thumbnails,omitempty"`   // 缩略图
	Videos      common.ArticleVideos     `form:"videos" json:"videos,omitempty"`           // 视频
	Audios      common.ArticleAudios     `form:"audios" json:"audios,omitempty"`           // 音频
	PostIP      string                   `form:"post_ip" json:"post_ip,omitempty"`         // 发布IP
	Weight      uint64                   `form:"weight" json:"weight,omitempty"`           // 权重
}

type ArticleListQueryFilter struct {
	TenantID    *uint64               `query:"tenant_id" json:"tenant_id,omitempty"`     // 租户ID
	UserID      *uint64               `query:"user_id" json:"user_id,omitempty"`         // 用户ID
	UUID        *string               `query:"uuid" json:"uuid,omitempty"`               // UUID
	BookID      *uint64               `query:"book_id" json:"book_id,omitempty"`         // 书ID
	ChapterID   *uint64               `query:"chapter_id" json:"chapter_id,omitempty"`   // 章节ID
	CategoryID  *uint64               `query:"category_id" json:"category_id,omitempty"` // 分类
	PublishAt   *time.Time            `query:"publish_at" json:"publish_at,omitempty"`   // 发布时间
	Type        *consts.ArticleType   `query:"type" json:"type,omitempty"`               // 文章类型
	Format      *consts.ArticleFormat `query:"format" json:"format,omitempty"`           // 文章格式
	Title       *string               `query:"title" json:"title,omitempty"`             // 标题
	Keyword     *string               `query:"keyword" json:"keyword,omitempty"`         // 关键词
	Description *string               `query:"description" json:"description,omitempty"` // 简介
	PostIP      *string               `query:"post_ip" json:"post_ip,omitempty"`         // 发布IP
	Weight      *uint64               `query:"weight" json:"weight,omitempty"`           // 权重
}

type ArticleItem struct {
	ID            uint64                   `json:"id,omitempty"`             // ID
	CreatedAt     time.Time                `json:"created_at,omitempty"`     // 创建时间
	UpdatedAt     time.Time                `json:"updated_at,omitempty"`     // 更新时间
	TenantID      uint64                   `json:"tenant_id,omitempty"`      // 租户ID
	UserID        uint64                   `json:"user_id,omitempty"`        // 用户ID
	UUID          string                   `json:"uuid,omitempty"`           // UUID
	BookID        uint64                   `json:"book_id,omitempty"`        // 书ID
	ChapterID     uint64                   `json:"chapter_id,omitempty"`     // 章节ID
	CategoryID    uint64                   `json:"category_id,omitempty"`    // 分类
	PublishAt     time.Time                `json:"publish_at,omitempty"`     // 发布时间
	Type          consts.ArticleType       `json:"type,omitempty"`           // 文章类型
	TypeCN        string                   `json:"type_cn,omitempty"`        // 文章类型
	Format        consts.ArticleFormat     `json:"format,omitempty"`         // 文章格式
	FormatCN      string                   `json:"format_cn,omitempty"`      // 文章格式
	Title         string                   `json:"title,omitempty"`          // 标题
	Keyword       string                   `json:"keyword,omitempty"`        // 关键词
	Description   string                   `json:"description,omitempty"`    // 简介
	Thumbnails    common.ArticleThumbnails `json:"thumbnails,omitempty"`     // 缩略图
	Videos        common.ArticleVideos     `json:"videos,omitempty"`         // 视频
	Audios        common.ArticleAudios     `json:"audios,omitempty"`         // 音频
	PostIP        string                   `json:"post_ip,omitempty"`        // 发布IP
	Weight        uint64                   `json:"weight,omitempty"`         // 权重
	Attachments   []ArticleAttachmentItem  `json:"attachments,omitempty"`    // 附件
	Dig           ArticleDigItem           `json:"dig,omitempty"`            // 点赞
	Content       ArticleContentItem       `json:"content,omitempty"`        // 内容
	ForwardSource ArticleForwardSourceItem `json:"forward_source,omitempty"` // 转发源
	Payments      []ArticlePaymentItem     `json:"payment,omitempty"`        // 支付
}

type ArticleDigItem struct {
	Views    uint64 `json:"views,omitempty"`
	Likes    uint64 `json:"likes,omitempty"`
	Dislikes uint64 `json:"dislikes,omitempty"`
}

type ArticleAttachmentItem struct {
	FilesystemID uint64 `json:"filesystem_id,omitempty"` // 文件系统ID
	Description  string `json:"description,omitempty"`   // 文件描述
	Password     string `json:"password,omitempty"`      // 文件密码
}

type ArticleContentItem struct {
	FreeContent  string `json:"free_content,omitempty"`  // 内容
	PriceContent string `json:"price_content,omitempty"` // 付费内容
}

type ArticleForwardSourceItem struct {
	Source       string `json:"source,omitempty"`        // 原文地址
	SourceAuthor string `json:"source_author,omitempty"` // 原文作者
}

type ArticlePaymentItem struct {
	PriceType       int64     `json:"price_type,omitempty"`        // 付费类型
	Token           string    `json:"token,omitempty"`             // 付费密码
	Price           uint64    `json:"price,omitempty"`             // 付费价格
	Discount        uint64    `json:"discount,omitempty"`          // 付费折扣
	DiscountStartAt time.Time `json:"discount_start_at,omitempty"` // 折扣开始时间
	DiscountEndAt   time.Time `json:"discount_end_at,omitempty"`   // 折扣结束时间
}
