package controller

import (
	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/modules/posts/dto"
	"github.com/atom-apps/posts/modules/posts/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type ArticleController struct {
	articleSvc *service.ArticleService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ArticleID"
//	@Success		200	{object}	dto.ArticleItem
//	@Router			/v1/posts/articles/{id} [get]
func (c *ArticleController) Show(ctx *fiber.Ctx, id uint64) (*dto.ArticleItem, error) {
	item, err := c.articleSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.articleSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.ArticleListQueryFilter	true	"ArticleListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.ArticleItem}
//	@Router			/v1/posts/articles [get]
func (c *ArticleController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.ArticleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.articleSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.articleSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.ArticleForm	true	"ArticleForm"
//	@Success		200		{string}	ArticleID
//	@Router			/v1/posts/articles [post]
func (c *ArticleController) Create(ctx *fiber.Ctx, body *dto.ArticleForm) error {
	return c.articleSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"ArticleID"
//	@Param			body	body		dto.ArticleForm	true	"ArticleForm"
//	@Success		200		{string}	ArticleID
//	@Router			/v1/posts/articles/{id} [put]
func (c *ArticleController) Update(ctx *fiber.Ctx, id uint64, body *dto.ArticleForm) error {
	return c.articleSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ArticleID"
//	@Success		200	{string}	ArticleID
//	@Router			/v1/posts/articles/{id} [delete]
func (c *ArticleController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.articleSvc.Delete(ctx.Context(), id)
}
