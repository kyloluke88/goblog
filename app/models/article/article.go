package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	// ID    uint64 // 为什么这里必须要注释掉？？？因为基类中也声明了ID
	Title  string    `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string    `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64    `gorm:"not null;index"`
	User   user.User // Preload("User") 关系调用时的关键字
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
