package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	// First() 是 gorm.DB 提供的用以从结果集中获取第一条数据的查询方法，需要注意的是第二个参数可以传参整型或者字符串 ID，
	// 但是传字符串会有 SQL 注入的风险，所以安全起见，我们使用 StringToUint64 做类型转换。
	// 在 GORM 中，当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.RouteName2URL("home"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

// Create 创建文章
func (article *Article) Create() (err error) {
	// create() 方法有以下三种返回，通过article.ID 来判断是否创建成功
	// article.ID             // 返回插入数据的主键
	// result.Error           // Create结果返回 error
	// result.RowsAffected    // 返回插入记录的条数
	result := model.DB.Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	// save() 方法有两个返回
	// result.RowsAffected // 更新的记录数
	// result.Error        // 更新的错误
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}
