package requests

import (
	"errors"
	"fmt"
	"goblog/pkg/model"
	"strconv"
	"strings"
	"unicode/utf8"

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

	// max_cn:8
	govalidator.AddCustomRule("max_cn", func(filed string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string)) // 表单传来的值

		length, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:")) // 规定的长度

		if valLength > length {
			// 超长，返回错误
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", length)
		}
		return nil
	})

	// min_cn:2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:")) //handle other error
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})
}
