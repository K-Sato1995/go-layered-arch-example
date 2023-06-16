package usecase

import (
	"context"
	"go-layerd-arc/domain/models"
	"go-layerd-arc/domain/repositories"
)

type ArticleUseCase struct {
	articleRepository repositories.IArticleRepository
}

// 抽象に依存
func NewArticleUseCase(repo repositories.IArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepository: repo,
	}
}

func (usecase ArticleUseCase) GetAllArticles(ctx context.Context) ([]*models.Article, error) {
	articles, err := usecase.articleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
