package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)
	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println(sql.ErrNoRows)
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示文章
		// tmpl, err := template.ParseFiles("resources/views/articles/show.gohtml")
		// 下方代码参考 https://learnku.com/courses/go-basic/1.17/delete-article/11513#827331
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
			"Name2URL":       route.Name2URL,
			"Uint64ToString": types.Uint64ToString,
		}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
		// fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 1获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
        tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
        logger.LogError(err)

        // 3. 渲染模板，将所有文章的数据传输进去
        err = tmpl.Execute(w, articles)
        logger.LogError(err)
	}
}
