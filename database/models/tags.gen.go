// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID        uint64         `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`      // ID
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);comment:创建时间" json:"created_at"`                      // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);comment:更新时间" json:"updated_at"`                      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);comment:删除时间" json:"deleted_at" swaggertype:"string"` // 删除时间
	Name      string         `gorm:"column:name;type:varchar(12);not null;comment:名称" json:"name"`                           // 名称
	Count     uint64         `gorm:"column:count;type:bigint unsigned;comment:引用次数" json:"count"`                            // 引用次数
}

func (*Tag) TableName(namer schema.Namer) string {
	if namer == nil {
		return TableNameTag
	}
	return namer.TableName(TableNameTag)
}
