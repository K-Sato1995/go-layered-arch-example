package repo

import (
	"context"
	"go-layerd-arc/domain/model"
)

type IArticleRepository interface {
	FindAll(context.Context) ([]*model.Article, error)
}

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (repository ArticleRepository) FindAll(ctx context.Context) ([]*model.Article, error) {
	// データストアから取得
	article := model.Article{
		ID:      1,
		Title:   "Hello",
		Content: "World",
	}
	return []*model.Article{&article}, nil
}
