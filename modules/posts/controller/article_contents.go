package controller

import (
	"github.com/atom-apps/posts/modules/posts/dto"
	"github.com/atom-apps/posts/modules/posts/service"

	"github.com/gofiber/fiber/v2"
)

// @provider
type ArticleContentController struct {
	articleContentSvc *service.ArticleContentService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			id	path		int	true	"ArticleContentID"
//	@Success		200	{object}	dto.ArticleContentItem
//	@Router			/v1/posts/articles/{article_id}/contents/{id} [get]
func (c *ArticleContentController) Show(ctx *fiber.Ctx, articleId int, id uint64) (*dto.ArticleContentItem, error) {
	item, err := c.articleContentSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.articleContentSvc.DecorateItem(item, 0), nil
}
