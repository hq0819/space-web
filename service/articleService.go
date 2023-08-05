package service

import (
	"context"
	"fmt"
	"space-web/model"
	"space-web/result"
	"space-web/setting"
)

func GetArticleListByUserId(id int) ([]model.Article, error) {
	list := make([]model.Article, 0)
	err := setting.DB.Debug().Model(model.Article{}).Where("user_id", id).Scan(&list).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return list, nil
}

func QueryRecommendArticle(page *model.PageInfo) *result.Message {
	return result.Success(model.StartPage[model.Article](page))
}

func ArticleList(pageInfo *model.PageInfo) *result.Message {
	db := setting.DB.Debug().Select(`t.row_id as articleId,tt.row_id as userid,tt.username,t.title,t.content,t.pic_url as picUrl,date_format(t.create_time,'%Y/%m/%d %H:%i:%s') as createTime`).
		Table("t_article t").
		Joins("left join t_user tt on t.user_id = tt.row_id").
		Order(`t.` + pageInfo.OrderBy)
	page := model.ToPage[ArticleVO](pageInfo, db)
	return result.Success(page)
}

func AddArticle(article *model.Article) error {
	return setting.DB.Debug().Model(model.Article{}).Create(article).Error
}

func QueryScanCount(artiID int) *result.Message {
	key := fmt.Sprintf(`%s:%d`, articleScanPrefix, artiID)
	res, _ := setting.Rdb.Get(context.Background(), key).Result()
	return result.Success(res)
}
func QueryCommentCount(artiID int) *result.Message {
	var count int64
	setting.DB.Debug().Model(model.Comment{}).Where("article_id=?", artiID).Count(&count)
	return result.Success(count)
}

func QuerySupportCount(artiID int) *result.Message {
	key := fmt.Sprintf(`%s:%d:*`, articleScanPrefix, artiID)
	res, _ := setting.Rdb.Keys(context.Background(), key).Result()
	return result.Success(len(res))
}
