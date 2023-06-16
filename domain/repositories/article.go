package repositories

import (
	"context"
	"go-layerd-arc/domain/models"
)

type IArticleRepository interface {
	FindAll(context.Context) ([]*models.Article, error)
}
