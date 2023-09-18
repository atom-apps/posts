package dao

import (
	"context"

	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/database/models"
	"github.com/atom-apps/posts/database/query"
	"github.com/atom-apps/posts/modules/posts/dto"

	"gorm.io/gen/field"
)

// @provider
type ArticleDao struct {
	query *query.Query
}

func (dao *ArticleDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleDao) Context(ctx context.Context) query.IArticleDo {
	return dao.query.Article.WithContext(ctx)
}

func (dao *ArticleDao) decorateSortQueryFilter(query query.IArticleDo, sortFilter *common.SortQueryFilter) query.IArticleDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Article.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Article.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *ArticleDao) decorateQueryFilter(query query.IArticleDo, queryFilter *dto.ArticleListQueryFilter) query.IArticleDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.Article.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Article.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.UUID != nil {
		query = query.Where(dao.query.Article.UUID.Eq(*queryFilter.UUID))
	}
	if queryFilter.BookID != nil {
		query = query.Where(dao.query.Article.BookID.Eq(*queryFilter.BookID))
	}
	if queryFilter.ChapterID != nil {
		query = query.Where(dao.query.Article.ChapterID.Eq(*queryFilter.ChapterID))
	}
	if queryFilter.CategoryID != nil {
		query = query.Where(dao.query.Article.CategoryID.Eq(*queryFilter.CategoryID))
	}
	if queryFilter.PublishAt != nil {
		query = query.Where(dao.query.Article.PublishAt.Eq(*queryFilter.PublishAt))
	}
	if queryFilter.Type != nil {
		query = query.Where(dao.query.Article.Type.Eq(*queryFilter.Type))
	}
	if queryFilter.Format != nil {
		query = query.Where(dao.query.Article.Format.Eq(*queryFilter.Format))
	}
	if queryFilter.Title != nil {
		query = query.Where(dao.query.Article.Title.Eq(*queryFilter.Title))
	}
	if queryFilter.Keyword != nil {
		query = query.Where(dao.query.Article.Keyword.Eq(*queryFilter.Keyword))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.Article.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.Thumbnails != nil {
		query = query.Where(dao.query.Article.Thumbnails.Eq(*queryFilter.Thumbnails))
	}
	if queryFilter.Videos != nil {
		query = query.Where(dao.query.Article.Videos.Eq(*queryFilter.Videos))
	}
	if queryFilter.Audios != nil {
		query = query.Where(dao.query.Article.Audios.Eq(*queryFilter.Audios))
	}
	if queryFilter.PostIP != nil {
		query = query.Where(dao.query.Article.PostIP.Eq(*queryFilter.PostIP))
	}
	if queryFilter.Weight != nil {
		query = query.Where(dao.query.Article.Weight.Eq(*queryFilter.Weight))
	}

	return query
}

func (dao *ArticleDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleDao) Update(ctx context.Context, model *models.Article) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Article.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Article.ID.Eq(id)).UpdateSimple(dao.query.Article.DeletedAt.Null())
	return err
}

func (dao *ArticleDao) Create(ctx context.Context, model *models.Article) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleDao) GetByID(ctx context.Context, id uint64) (*models.Article, error) {
	return dao.Context(ctx).Where(dao.query.Article.ID.Eq(id)).First()
}

func (dao *ArticleDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Article, error) {
	return dao.Context(ctx).Where(dao.query.Article.ID.In(ids...)).Find()
}

func (dao *ArticleDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Article, int64, error) {
	query := dao.query.Article
	articleQuery := query.WithContext(ctx)
	articleQuery = dao.decorateQueryFilter(articleQuery, queryFilter)
	articleQuery = dao.decorateSortQueryFilter(articleQuery, sortFilter)
	return articleQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *ArticleDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Article, error) {
	query := dao.query.Article
	articleQuery := query.WithContext(ctx)
	articleQuery = dao.decorateQueryFilter(articleQuery, queryFilter)
	articleQuery = dao.decorateSortQueryFilter(articleQuery, sortFilter)
	return articleQuery.Find()
}

func (dao *ArticleDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Article, error) {
	query := dao.query.Article
	articleQuery := query.WithContext(ctx)
	articleQuery = dao.decorateQueryFilter(articleQuery, queryFilter)
	articleQuery = dao.decorateSortQueryFilter(articleQuery, sortFilter)
	return articleQuery.First()
}
