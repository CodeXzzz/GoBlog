package dao

import (
	"GoBlog/common"
	"GoBlog/internal/model"
)

// GetAllSummaries 用户登录成功后，在主页展示文章摘要
func GetAllSummaries() (articleRes []*model.ArticleRes, err error) {
	//在article表中查询选定字段
	err = common.MDB.Table("articles").
		Select("aid,author,title,cover,summary").Where("deleted_at", nil).Scan(&articleRes).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetArticle 根据aid获取文章详情
func GetArticle(aid int) (*model.Article, error) {
	article := new(model.Article)
	if err := common.MDB.Where("aid=?", aid).First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// PostArticle 用户发布文章
func PostArticle(article *model.Article) (err error) {
	if err = common.MDB.Table("articles").Create(article).Error; err != nil {
		return err
	}
	return nil
}

// UpdateArticle 用户更新文章
func UpdateArticle(article *model.Article) (err error) {
	if err = common.MDB.Table("articles").Where("aid=?", article.Aid).Updates(article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle 用户删除文章（软删除）
func DeleteArticle(aid int) (err error) {
	//var article *model.Article仅仅声明，没有分配值
	article := new(model.Article)
	if err = common.MDB.Table("articles").Where("aid=?", aid).Delete(article).Error; err != nil {
		return err
	}
	return err
}
func GetUserArticle(uid int) (articleList []*model.Article, err error) {
	if err = common.MDB.Table("articles").Where("uid=?", uid).Find(&articleList).Error; err != nil {
		return nil, err
	}
	return articleList, nil
}
