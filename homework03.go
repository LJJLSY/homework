package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Posts     []Post // Has Many: One user has many posts
	PostCount uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID            uint
	Title         string
	Content       string
	UserID        uint      // Foreign key to user
	Comments      []Comment // Has Many: One post has many comments
	CommentCount  uint
	CommentStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Comment struct {
	ID            uint
	Content       string
	CommentUserID uint
	PostID        uint // Foreign key to post
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	if p.UserID == 0 {
		return fmt.Errorf("userid不能为null")
	}

	if err := db.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count+1")).Error; err != nil {
		log.Fatalf("更新用户post_count失败: %v", err)
	}

	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) error {
	if err := db.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count-1")).Error; err != nil {
		log.Fatalf("更新文章comment_count失败: %v", err)
	}

	var post Post
	if err := db.Model(&Post{}).
		Where("id = ?", c.PostID).
		Find(&post).Error; err != nil {
		log.Fatalf("获取最新post数据: %v", err)
	}

	if post.CommentCount == 0 {
		if err := db.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error; err != nil {
			log.Fatalf("更新文章comment_status失败: %v", err)
		}
	}

	return nil
}

// 连接数据库
func AutoMigrate() *gorm.DB {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/TEST?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取DB失败: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("数据库连接测试失败: %v", err)
	}
	log.Println("数据库连接成功")

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("表创建失败: %v", err)
	}
	return db
}

// 初始化数据
func create() {
	db := AutoMigrate()
	//清空数据
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM comments")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	//插入数据
	u := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Posts: []Post{
			{
				Title:   "关于gorm框架简介",
				Content: "GORM 是 Go 语言中最流行的 ORM 框架，用于在面向对象编程语言和关系型数据库之间建立映射关系",
				Comments: []Comment{
					{
						Content:       "言简意赅，通俗易懂",
						CommentUserID: 1,
					},
					{
						Content:       "讲得真清楚",
						CommentUserID: 2,
					},
				},
			},
		},
	}

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&u).Error; err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}

	u1 := User{
		Name:  "Bob",
		Email: "Bob@example.com",
		Posts: []Post{
			{
				Title:   "关于gorm框架",
				Content: "GORM 是 Go 语言中最流行的 ORM 框架",
				Comments: []Comment{
					{
						Content:       "清晰易懂",
						CommentUserID: 10,
					},
					{
						Content:       "一目了然",
						CommentUserID: 20,
					},
				},
			},
			{
				Title:   "关于gin框架简介",
				Content: "Gin 是 Go 语言中最流行的 Web 框架，它提供了一系列工具和库，帮助开发者处理路由、请求解析、响应生成等常见任务",
				Comments: []Comment{
					{
						Content:       "理解透彻",
						CommentUserID: 10,
					},
					{
						Content:       "讲得很好",
						CommentUserID: 20,
					},
					{
						Content:       "思路清晰",
						CommentUserID: 30,
					},
				},
			},
		},
	}

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&u1).Error; err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}
}

// 关联查询
func associations(userid uint) {
	db := AutoMigrate()
	var post []Post
	if err := db.Preload("Comments").Where("user_id = ?", userid).Find(&post).Error; err != nil {
		log.Fatalf("查询文章失败: %v", err)
	}
	fmt.Println(post)
}

// 评论数量最多的文章信息
func mostcomments() {
	db := AutoMigrate()
	type commentcount struct {
		Post_id int64
		Total   int64
	}
	comment := []commentcount{}

	if err := db.Model(&Comment{}).Select("post_id,count(*) as total").
		Group("post_id").Order("total desc").Scan(&comment).Error; err != nil {
		log.Fatalf("统计文章评论数量失败: %v", err)
	}

	var post []Post
	if err := db.Preload("Comments").Where("id = ?", comment[0].Post_id).Find(&post).Error; err != nil {
		log.Fatalf("获取评论数量最多的文章失败: %v", err)
	}
	fmt.Println(post)
}

// 新增文章数据
func PostCreate() {
	db := AutoMigrate()

	Posts1 := Post{
		Title:   "测试更新用户1文章数量",
		Content: "用户1每次新增一篇文章，PostCount应该增加相应数量",
		UserID:  1,
	}

	if err := db.Create(&Posts1).Error; err != nil {
		log.Fatalf("用户1新增文章: %v", err)
	}

	Posts2 := []Post{
		{
			Title:   "测试更新用户2文章数量-1",
			Content: "用户2每次新增两篇文章，PostCount应该增加相应数量",
			UserID:  2,
		},
		{
			Title:   "测试更新用户2文章数量-2",
			Content: "用户2每次新增两篇文章，PostCount应该增加相应数量",
			UserID:  2,
		},
	}

	if err := db.Create(&Posts2).Error; err != nil {
		log.Fatalf("用户2新增文章: %v", err)
	}
}

// 删除评论
func CommentDelete() {
	db := AutoMigrate()

	var comment []Comment
	if err := db.Where("id between ? and ?", 1, 3).
		Find(&comment).Error; err != nil {
		log.Fatalf("评论查找: %v", err)
	}

	for _, c := range comment {
		if err := db.Delete(&c).Error; err != nil {
			log.Fatalf("删除评论: %v", err)
		}
	}
}

func main() {
	//create()
	associations(2)
	mostcomments()
	//PostCreate()
	//CommentDelete()
}
