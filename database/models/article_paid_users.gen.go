// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm/schema"
)

const TableNameArticlePaidUser = "article_paid_users"

// ArticlePaidUser mapped from table <article_paid_users>
type ArticlePaidUser struct {
	ID        uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                 // 创建时间
	TenantID  uint64    `gorm:"column:tenant_id;type:bigint unsigned;comment:租户ID" json:"tenant_id"`               // 租户ID
	UserID    uint64    `gorm:"column:user_id;type:bigint unsigned;comment:用户ID" json:"user_id"`                   // 用户ID
	ArticleID uint64    `gorm:"column:article_id;type:bigint unsigned;comment:文章ID" json:"article_id"`             // 文章ID
	PriceType int64     `gorm:"column:price_type;type:bigint;comment:付费类型" json:"price_type"`                      // 付费类型
	Price     uint64    `gorm:"column:price;type:bigint unsigned;comment:付费价格" json:"price"`                       // 付费价格
	EndAt     uint64    `gorm:"column:end_at;type:bigint unsigned;comment:付费结束时间" json:"end_at"`                   // 付费结束时间
}

func (*ArticlePaidUser) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameArticlePaidUser
	}
	return namer.TableName(TableNameArticlePaidUser)
}
