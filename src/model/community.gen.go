// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCommunity = "community"

// Community mapped from table <community>
type Community struct {
	ID            int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CommunityID   int32          `gorm:"column:community_id;not null" json:"community_id"`
	CommunityName string         `gorm:"column:community_name;not null" json:"community_name"`
	Introduction  string         `gorm:"column:introduction;not null" json:"introduction"`
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Community's table name
func (*Community) TableName() string {
	return TableNameCommunity
}