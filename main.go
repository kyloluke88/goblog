package main

// 关于包中的init函数  查阅：https://learnku.com/go/t/47178
import (
	"database/sql"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"net/http"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" // 匿名导入， 当导入了一个数据库驱动后，此驱动会自行初始化（利用 init() 函数）并注册自己到 Golang 的 database/sql 上下文中
	"github.com/gorilla/mux"
)

var router *mux.Router
var db *sql.DB

func initDB() {

	// 该内容，请参考  https://learnku.com/courses/go-basic/1.17/connect-to-database/11507
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3308",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())

	logger.LogError(err)
	// 以下三个配置信息参考： https://learnku.com/courses/go-basic/1.17/connect-to-database/11507#88a495
	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := db.Exec(createArticlesSQL)
	logger.LogError(err)
}

func main() {

	database.Initialize()
	db = database.DB

	router = bootstrap.SetupRoute()
	bootstrap.SetupDB()

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	// homeURL, _ := router.Get("home").URL()
	// fmt.Println("homeURL: ", homeURL)
	// articleURL, _ := router.Get("articles.show").URL("id", "1")
	// fmt.Println("articleURL: ", articleURL)

	// 1.Gorilla Mux 会先匹配路由，再执行中间件，故使用中间件永远会返回 404.
	// 2.所以写一个函数把 Gorilla Mux 包起来，在这个函数中我们先对进来的请求做处理，然后再传给 Gorilla Mux 去解析。
	http.ListenAndServe(":3000", removeTrailingSlash(router))
}

// 中间件
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

// 请求进来时，删除掉  /
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		// 2. 将请求传递下去
		// fmt.Print(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
