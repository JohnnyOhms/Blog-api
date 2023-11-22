package models

import (
	"github.com/JohnnyOhms/BlogPost-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Blogs struct {
	gorm.Model
	BlogId int64  `gorm:"int;unique"`
	Arthur string `gorm:"int;unique"`
	Title  string `gorm:"type:varchar(100)"`
	Blog   string `gorm:"type:varchar(1000)"`
	Likes  int64  `gorm:"int"`
}

func Init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Blogs{})
}

func (b *Blogs) CreateBlog() *Blogs {
	db.Create(&b)
	return b
}

func GetBlogs() []Blogs {
	var blogs []Blogs
	db.Find(&blogs)
	return blogs
}

func GetBlog(Id int64) (*Blogs, *gorm.DB) {
	var blog Blogs
	db := db.Where("BlogId=?", Id).Find(blog)
	return &blog, db
}

func (b *Blogs) UpdateBlog(Id int64, field string) *Blogs {
	db.Where("BlogId=?", Id).Update(field, &b)
	return b
}

func DeleteBlog(Id int64) *Blogs {
	var blog Blogs
	db.Where("BlogId=?", Id).Delete(&blog)
	return &blog
}
