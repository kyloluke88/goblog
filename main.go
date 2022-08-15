package main

// 关于包中的init函数  查阅：https://learnku.com/go/t/47178
import (
	"embed"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()

	// 初始化模板
	bootstrap.SetupTemplate(tplFS)

	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
