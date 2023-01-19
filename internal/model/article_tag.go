package model

type BlogArticleTag struct {
	*Model
	ArticleId int32 `json:"article_id"` // 文章ID
	TagId     int32 `json:"tag_id"`     // 标签ID
}

func (model BlogArticleTag) TableName() string {
	return "blog_article_tag"
}
