package models

type Post struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID"`
}

type PostDetails struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}