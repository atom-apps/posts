package dto

import (
	"time"
)

type ArticlePaymentForm struct {
	ArticleID       uint64    `form:"article_id" json:"article_id,omitempty"`               // 文章ID
	PriceType       int64     `form:"price_type" json:"price_type,omitempty"`               // 付费类型
	Token           string    `form:"token" json:"token,omitempty"`                         // 付费密码
	Price           uint64    `form:"price" json:"price,omitempty"`                         // 付费价格
	Discount        uint64    `form:"discount" json:"discount,omitempty"`                   // 付费折扣
	DiscountStartAt time.Time `form:"discount_start_at" json:"discount_start_at,omitempty"` // 折扣开始时间
	DiscountEndAt   time.Time `form:"discount_end_at" json:"discount_end_at,omitempty"`     // 折扣结束时间
}

type ArticlePaymentListQueryFilter struct {
	ArticleID       *uint64    `query:"article_id" json:"article_id,omitempty"`               // 文章ID
	PriceType       *int64     `query:"price_type" json:"price_type,omitempty"`               // 付费类型
	Token           *string    `query:"token" json:"token,omitempty"`                         // 付费密码
	Price           *uint64    `query:"price" json:"price,omitempty"`                         // 付费价格
	Discount        *uint64    `query:"discount" json:"discount,omitempty"`                   // 付费折扣
	DiscountStartAt *time.Time `query:"discount_start_at" json:"discount_start_at,omitempty"` // 折扣开始时间
	DiscountEndAt   *time.Time `query:"discount_end_at" json:"discount_end_at,omitempty"`     // 折扣结束时间
}

type ArticlePaymentItem struct {
	ID              uint64    `json:"id,omitempty"`                // ID
	CreatedAt       time.Time `json:"created_at,omitempty"`        // 创建时间
	ArticleID       uint64    `json:"article_id,omitempty"`        // 文章ID
	PriceType       int64     `json:"price_type,omitempty"`        // 付费类型
	Token           string    `json:"token,omitempty"`             // 付费密码
	Price           uint64    `json:"price,omitempty"`             // 付费价格
	Discount        uint64    `json:"discount,omitempty"`          // 付费折扣
	DiscountStartAt time.Time `json:"discount_start_at,omitempty"` // 折扣开始时间
	DiscountEndAt   time.Time `json:"discount_end_at,omitempty"`   // 折扣结束时间
}
