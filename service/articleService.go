package service

import (
	"fmt"
	"space-web/model"
	"space-web/setting"
)

func GetArticleListByUserId(id int) ([]model.Article, error) {
	list := make([]model.Article, 0)
	err := setting.DB.Debug().Model(model.Article{}).Where("user_id", id).Scan(list).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return list, nil
}

func AddArticle(article *model.Article) error {
	return setting.DB.Debug().Model(model.Article{}).Create(article).Error
}
