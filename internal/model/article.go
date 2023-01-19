package model

type BlogArticle struct {
	*Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址
	Content       string `json:"content"`         // 文章内容
	State         int8   `json:"state"`           // 状态：0为禁用、1为启用
}

func (model BlogArticle) TableName() string {
	return "blog_article"
}
