package routes

import (
	"net/http"

	"go-layerd-arc/usecase"

	"go-layerd-arc/domain/repo"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	articleUseCase := usecase.NewArticleUseCase(
		repo.NewArticleRepository(),
	)

	e.GET("/articles", func(c echo.Context) error {
		articles, err := articleUseCase.GetAllArticles(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, articles)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
