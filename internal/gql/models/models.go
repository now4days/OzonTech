package models

type Post struct {
	ID            string `gorm:"primary_key"`
	Title         string
	Content       string
	AllowComments bool
	Comments      []Comment
}

type Comment struct {
	ID       string `gorm:"primary_key"`
	Content  string
	PostID   string
	ParentID *string
	Replies  []Comment `gorm:"foreignkey:ParentID"`
}
