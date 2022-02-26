package article

import "github.com/dyaodeng/wx/base"

// Article 公众号文章管理
type Article struct {
	accessToken base.AccessTokenInterface
}

func NewArticle(accessToken base.AccessTokenInterface) *Article {
	return &Article{accessToken: accessToken}
}
