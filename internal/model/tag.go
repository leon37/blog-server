package model

type BlogTag struct {
	*Model
	Name  string `json:"name"`  // 标签名称
	State int8   `json:"state"` // 状态 0为禁用、1为启用
}

func (model BlogTag) TableName() string {
	return "blog_tag"
}
