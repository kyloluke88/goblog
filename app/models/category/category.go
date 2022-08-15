package category

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255); not null;" valid:"name"`
}

// Link 获取URL
func (c *Category) Link() string {
	// return route.RouteName2URL("categories.show", "id", strconv.FormatUint(category.ID, 10))
	return route.RouteName2URL("categories.show", "id", c.GetStringID())
}
