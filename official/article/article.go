package article

import (
	"github.com/dysodeng/wx/kernel/contracts"
)

// Article 公众号文章管理
type Article struct {
	accessToken contracts.AccessTokenInterface
}

func NewArticle(accessToken contracts.AccessTokenInterface) *Article {
	return &Article{accessToken: accessToken}
}
