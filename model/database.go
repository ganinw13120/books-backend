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

var ReviewTable = struct {
	ID           string
	UserID       string
	Title        string
	BookName     string
	BookAuthor   string
	BookImageUrl string
	Description  string
	CreatedAt    string
	UpdatedAt    string
}{
	"review_id",
	"user_id",
	"title",
	"book_name",
	"book_author",
	"book_img_url",
	"description",
	"created_at",
	"updated_at",
}

type User struct {
	ID        uint      `gorm:"column:user_id" gorm:"primaryKey"`
	name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Review struct {
	ID           uint      `json:"id" gorm:"column:review_id" gorm:"primaryKey"`
	UserID       int       `json:"user_id" gorm:"column:user_id"`
	Title        string    `json:"title" gorm:"column:title"`
	BookName     string    `json:"book_name" gorm:"column:book_name"`
	BookAuthor   string    `json:"book_author" gorm:"column:book_author"`
	BookImageUrl string    `json:"book_img_url" gorm:"column:book_image_url"`
	Description  string    `json:"description" gorm:"column:description"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type React struct {
	UserID    int       `gorm:"primaryKey;column:user_id" json:"user_id"`
	ReviewID  int       `gorm:"primaryKey;column:review_id" json:"review_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
