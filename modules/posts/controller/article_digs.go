package controller

import (
	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/modules/posts/dto"
	"github.com/atom-apps/posts/modules/posts/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type ArticleDigController struct {
	articleDigSvc *service.ArticleDigService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			id	path		int	true	"ArticleDigID"
//	@Success		200	{object}	dto.ArticleDigItem
//	@Router			/v1/posts/articles/{article_id}/dig/{id} [get]
func (c *ArticleDigController) Show(ctx *fiber.Ctx, articleId int, id uint64) (*dto.ArticleDigItem, error) {
	item, err := c.articleDigSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.articleDigSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			queryFilter	query		dto.ArticleDigListQueryFilter	true	"ArticleDigListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.ArticleDigItem}
//	@Router			/v1/posts/articles/{article_id}/dig [get]
func (c *ArticleDigController) List(
	ctx *fiber.Ctx,
	articleId int,
	queryFilter *dto.ArticleDigListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.articleDigSvc.PageByQueryFilter(ctx.Context(), articleId, queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.articleDigSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			body	body		dto.ArticleDigForm	true	"ArticleDigForm"
//	@Success		200		{string}	ArticleDigID
//	@Router			/v1/posts/articles/{article_id}/dig [post]
func (c *ArticleDigController) Create(ctx *fiber.Ctx, articleId int, body *dto.ArticleDigForm) error {
	return c.articleDigSvc.Create(ctx.Context(), articleId, body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			id		path		int				true	"ArticleDigID"
//	@Param			body	body		dto.ArticleDigForm	true	"ArticleDigForm"
//	@Success		200		{string}	ArticleDigID
//	@Router			/v1/posts/articles/{article_id}/dig/{id} [put]
func (c *ArticleDigController) Update(ctx *fiber.Ctx, articleId int, id uint64, body *dto.ArticleDigForm) error {
	return c.articleDigSvc.Update(ctx.Context(), articleId, id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			articleId	path		int	true	"ArticleId"
//	@Param			id	path		int	true	"ArticleDigID"
//	@Success		200	{string}	ArticleDigID
//	@Router			/v1/posts/articles/{article_id}/dig/{id} [delete]
func (c *ArticleDigController) Delete(ctx *fiber.Ctx, articleId int, id uint64) error {
	return c.articleDigSvc.Delete(ctx.Context(), articleId, id)
}
