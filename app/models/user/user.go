package user

import "goblog/app/models"

// User 用户模型，数据验证同时可以用
type User struct {
	models.BaseModel

	// 默认将键小写化作为字段名，所以下方的  column:name 等可以去除
	// 默认是可以为null的，所以 default:NULL 也可以去除
	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);default:NULL;unique;" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)" valid:"password"`
	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(password string) bool {
	return user.Password == password
}
