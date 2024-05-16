// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PostID      int64          `gorm:"column:post_id;not null" json:"post_id"`
	Title       string         `gorm:"column:title;not null" json:"title"`
	Content     string         `gorm:"column:content;not null" json:"content"`
	AuthorID    int64          `gorm:"column:author_id;not null" json:"author_id"`
	CommunityID int32          `gorm:"column:community_id;not null" json:"community_id"`
	Status      int32          `gorm:"column:status;not null" json:"status"`
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}