package database

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// User for storing users whoever intract with the bot
type User struct {
	gorm.Model
	ID        int            `gorm:"primaryKey"`
	UserId    int64          `gorm:"unique;column:user_id"`
	FirstName sql.NullString `gorm:"column:first_name"`
	LastName  sql.NullString `gorm:"column:last_name"`
	Username  sql.NullString `gorm:"column:username"`
	FirstSeen time.Time      `gorm:"autoCreateTime"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Rss for links, and subscribers
type Rss struct {
	gorm.Model
	ID                int       `gorm:"primaryKey"`
	OwnerId           int64     `gorm:"column:owner_id"`
	Link              string    `gorm:"column:link"`
	LastItemGUID      string    `gorm:"size:255;column:last_item_guid"`
	Subscribers       []User    `gorm:"foreignKey:UserId"`
	NotificationCount int64     `gorm:"default:0;column:notification_count"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
