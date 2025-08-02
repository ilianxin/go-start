package main

/*
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

*/

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Posts     []Post `gorm:"foreignKey:UserId"`
	PostCount int    `gorm:"default:0"` // 文章数量统计字段
}

type Post struct {
	ID            int `db:"id"`
	UserId        int
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	CreatedAt     time.Time `db:"created_at"`
	Comments      []Comment `gorm:"foreignKey:PostId"`
	CommentStatus string    `gorm:"default:'有评论'"` // 评论状态字段
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", p.UserId).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	return nil
}

type Comment struct {
	ID        int `db:"id"`
	PostId    int
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.First(&post, c.PostId).Error; err != nil {
		return err
	}

	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if count == 0 {
		tx.Model(&post).Update("comment_status", "无评论")
	}

	return nil
}

func main() {
	user := "rwuser"
	pass := "Rwpass@123"
	dsn := user + ":" + pass + "@tcp(47.111.78.104:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	var users []User

	var postInfoMap = make(map[int]int)

	result := db.Model(&User{}).Preload("Posts.Comments").Find(&users)

	if result.Error != nil {
		fmt.Printf("Error fetching users: %v\n", result.Error)
		return

	}

	for _, user := range users {
		fmt.Printf("User: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)
		for _, post := range user.Posts {
			fmt.Printf("  Post: ID=%d, Title=%s, Content=%s\n", post.ID, post.Title, post.Content)

			//result := db.Model(&Post{}).Association("Comments").Find(&post.Comments)

			postInfoMap[post.ID] = len(post.Comments)
			for _, comment := range post.Comments {
				fmt.Printf("    Comment: ID=%d, Content=%s\n", comment.ID, comment.Content)
			}
		}
	}

	var postId int
	var maxComments int

	for index, postCount := range postInfoMap {

		if postCount > maxComments {
			maxComments = postCount
			postId = index
		}
		fmt.Printf("Post ID: %d has %d comments\n", index, postCount)
	}

	fmt.Printf("Post with most comments has ID: %d\n", postId)

	newPost := Post{Title: "New Post", Content: "This is a new post", UserId: 2}

	db.Create(&newPost)

	delComment := Comment{ID: 3}
	if err := db.First(&delComment, 3).Error; err != nil {
		fmt.Println("Comment not found:", err)
	} else {
		db.Delete(&delComment)
	}

	fmt.Println("Post and Comment models created successfully.")
}
