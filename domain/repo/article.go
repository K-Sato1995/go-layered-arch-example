package repo

import (
	"context"
	"go-layerd-arc/domain/model"
)

type IArticleRepository interface {
	FindAll(context.Context) ([]*model.Article, error)
}

func FindAllArticles(ctx context.Context) ([]*model.Article, error) {
	article := model.Article{
		ID:      1,
		Title:   "Hello",
		Content: "World",
	}

	return []*model.Article{&article}, nil
}
