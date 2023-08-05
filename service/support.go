package service

import (
	"context"
	"fmt"
	. "space-web/setting"
	"strconv"
)

type ArticleVO struct {
	ArticleId  int64  `json:"articleId" gorm:"column:articleId"`
	AuthorId   int64  `json:"authorId" gorm:"column:userid"`
	Author     string `json:"author" gorm:"column:username"`
	Title      string `json:"title"`
	CreateTime string `json:"createTime" gorm:"column:createTime"`
	TimeAgo    string `json:"timeAgo"`
	PicUrl     string `json:"picUrl"  gorm:"column:picUrl"`
	Content    string `json:"content"`
}

const articleSupportPrefix = "article"
const commentPrefix = "comment"
const articleScanPrefix = "articleScan"

func AddArticleSupport(articleId int64, userId int64) {
	key := fmt.Sprintf(`%s:%d:%d`, articleSupportPrefix, articleId, userId)
	ctx := context.Background()
	result, _ := Rdb.Exists(ctx, key).Result()
	if result == 0 {
		Rdb.Set(ctx, key, "1", -1)
		return
	}
}

func AddCommentSupport(commentId int64, userId string) {
	commID := commentPrefix + strconv.FormatInt(commentId, 10) + userId
	ctx := context.Background()
	result, _ := Rdb.Exists(ctx, commID).Result()
	if result == 0 {
		Rdb.Set(ctx, commID, 1, -1)
		return
	}
}

func AddArticleScan(articleId int) {
	key := fmt.Sprintf(`%s:%d`, articleScanPrefix, articleId)
	ctx := context.Background()
	result, _ := Rdb.Exists(ctx, key).Result()
	if result == 0 {
		Rdb.Set(ctx, key, 1, -1)
		return
	}
	_, err := Rdb.Incr(context.Background(), key).Result()
	fmt.Println(err)
}
