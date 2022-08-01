package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	// ID    uint64 // 为什么这里必须要注释掉？？？因为基类中也声明了ID
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.RouteName2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
