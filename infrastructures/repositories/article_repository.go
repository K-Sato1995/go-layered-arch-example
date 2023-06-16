package repositoryImpl

import (
	"context"
	"go-layerd-arc/domain/models"
	"go-layerd-arc/domain/repositories"
)

type ArticleRepository struct{}

func NewArticleRepositoryImpl() repositories.IArticleRepository {
	return &ArticleRepository{}
}

func (repository ArticleRepository) FindAll(ctx context.Context) ([]*models.Article, error) {
	// データストアから取得
	article := models.Article{
		ID:      1,
		Title:   "Hello",
		Content: "World",
	}
	return []*models.Article{&article}, nil
}
