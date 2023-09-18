// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm/schema"
)

const TableNameArticleDig = "article_digs"

// ArticleDig mapped from table <article_digs>
type ArticleDig struct {
	ID        uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                 // 创建时间
	ArticleID uint64    `gorm:"column:article_id;type:bigint unsigned;comment:文章ID" json:"article_id"`             // 文章ID
	Views     uint64    `gorm:"column:views;type:bigint unsigned;comment:浏览量" json:"views"`                        // 浏览量
	Likes     uint64    `gorm:"column:likes;type:bigint unsigned;comment:喜欢" json:"likes"`                         // 喜欢
	Dislikes  uint64    `gorm:"column:dislikes;type:bigint unsigned;comment:不喜欢" json:"dislikes"`                  // 不喜欢
}

func (*ArticleDig) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameArticleDig
	}
	return namer.TableName(TableNameArticleDig)
}
