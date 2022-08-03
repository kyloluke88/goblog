package requests

import (
	"errors"
	"fmt"
	"goblog/pkg/model"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行
func init() {
	// not_exists:users,email
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0]   // user
		dbFiled := rng[1]     // filed name
		val := value.(string) // filed value

		var count int64
		model.DB.Table(tableName).Where(dbFiled+" = ?", val).Count(&count) // Count()函数会直接拿到结果给 count 赋值

		if count != 0 {

			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("%v 已被占用", val)
		}
		return nil
	})
}
