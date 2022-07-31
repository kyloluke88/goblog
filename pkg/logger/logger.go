package logger

import "log"

// 错误记录到日志
func LogError(err error) {
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
}
