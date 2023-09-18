package controller

import (
	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/modules/posts/dto"
	"github.com/atom-apps/posts/modules/posts/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type BookController struct {
	bookSvc *service.BookService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"BookID"
//	@Success		200	{object}	dto.BookItem
//	@Router			/v1/posts/books/{id} [get]
func (c *BookController) Show(ctx *fiber.Ctx, id uint64) (*dto.BookItem, error) {
	item, err := c.bookSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.bookSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.BookListQueryFilter	true	"BookListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.BookItem}
//	@Router			/v1/posts/books [get]
func (c *BookController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.BookListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.bookSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.bookSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.BookForm	true	"BookForm"
//	@Success		200		{string}	BookID
//	@Router			/v1/posts/books [post]
func (c *BookController) Create(ctx *fiber.Ctx, body *dto.BookForm) error {
	return c.bookSvc.Create(ctx.Context(), body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"BookID"
//	@Param			body	body		dto.BookForm	true	"BookForm"
//	@Success		200		{string}	BookID
//	@Router			/v1/posts/books/{id} [put]
func (c *BookController) Update(ctx *fiber.Ctx, id uint64, body *dto.BookForm) error {
	return c.bookSvc.Update(ctx.Context(), id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"BookID"
//	@Success		200	{string}	BookID
//	@Router			/v1/posts/books/{id} [delete]
func (c *BookController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.bookSvc.Delete(ctx.Context(), id)
}
