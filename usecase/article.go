package usecase

import (
	"context"
	"go-layerd-arc/domain/model"
	"go-layerd-arc/domain/repo"
)

type ArticleUseCase struct {
	articleRepository repo.IArticleRepository
}

// 抽象に依存
func NewArticleUseCase(repo repo.IArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		articleRepository: repo,
	}
}

func (usecase ArticleUseCase) GetAllArticles(ctx context.Context) ([]*model.Article, error) {
	articles, err := usecase.articleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
