package model

import "time"

var Tables = struct {
	Review string
	User   string
	React  string
}{
	"Review",
	"User",
	"React",
}

type User struct {
	ID        uint      `gorm:"column:user_id" gorm:"primaryKey"`
	name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Review struct {
	ID           uint      `gorm:"column:review_id" gorm:"primaryKey"`
	UserID       int       `gorm:"column:user_id"`
	Title        string    `gorm:"column:title"`
	BookName     string    `gorm:"column:book_name"`
	BookAuthor   string    `gorm:"column:book_author"`
	BookImageUrl string    `gorm:"column:book_image_url"`
	Description  string    `gorm:"column:description"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type React struct {
	UserID    int       `gorm:"primaryKey;column:user_id" json:"user_id"`
	ReviewID  int       `gorm:"primaryKey;column:review_id" json:"review_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
