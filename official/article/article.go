package article

import (
	"github.com/dysodeng/wx/kernel/contracts"
)

// Article 公众号文章管理
type Article struct {
	account contracts.AccountInterface
}

func NewArticle(account contracts.AccountInterface) *Article {
	return &Article{account: account}
}
