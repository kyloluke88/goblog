package requests

import (
	"goblog/app/models/category"

	"github.com/thedevsaddam/govalidator"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	// 1. 定制验证规则

	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	// 2. 自定义错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度至少2个字",
			"max_cn:分类名称长度不能超过8个子",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data, // struct
		Rules:         rules,
		TagIdentifier: "valid", // 模型中Struct标签标识符
		Messages:      messages,
	}

	// 4. 开始验证
	return govalidator.New(opts).ValidateStruct()
}
