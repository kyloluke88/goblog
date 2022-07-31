package main

// 关于包中的init函数  查阅：https://learnku.com/go/t/47178
import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
