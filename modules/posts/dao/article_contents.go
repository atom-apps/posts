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
type ArticleContentDao struct {
	query *query.Query
}

func (dao *ArticleContentDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleContentDao) Context(ctx context.Context) query.IArticleContentDo {
	return dao.query.ArticleContent.WithContext(ctx)
}

func (dao *ArticleContentDao) decorateSortQueryFilter(query query.IArticleContentDo, sortFilter *common.SortQueryFilter) query.IArticleContentDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.ArticleContent.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.ArticleContent.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *ArticleContentDao) decorateQueryFilter(query query.IArticleContentDo, queryFilter *dto.ArticleContentListQueryFilter) query.IArticleContentDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.ArticleContent.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.ArticleContent.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.ArticleID != nil {
		query = query.Where(dao.query.ArticleContent.ArticleID.Eq(*queryFilter.ArticleID))
	}
	if queryFilter.FreeContent != nil {
		query = query.Where(dao.query.ArticleContent.FreeContent.Eq(*queryFilter.FreeContent))
	}
	if queryFilter.PriceContent != nil {
		query = query.Where(dao.query.ArticleContent.PriceContent.Eq(*queryFilter.PriceContent))
	}

	return query
}

func (dao *ArticleContentDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleContentDao) Update(ctx context.Context, model *models.ArticleContent) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleContentDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleContentDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.ArticleContent.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleContentDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.ArticleContent.ID.Eq(id)).UpdateSimple(dao.query.ArticleContent.DeletedAt.Null())
	return err
}

func (dao *ArticleContentDao) Create(ctx context.Context, model *models.ArticleContent) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleContentDao) GetByID(ctx context.Context, id uint64) (*models.ArticleContent, error) {
	return dao.Context(ctx).Where(dao.query.ArticleContent.ID.Eq(id)).First()
}

func (dao *ArticleContentDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleContent, error) {
	return dao.Context(ctx).Where(dao.query.ArticleContent.ID.In(ids...)).Find()
}

func (dao *ArticleContentDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleContentListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleContent, int64, error) {
	query := dao.query.ArticleContent
	articleContentQuery := query.WithContext(ctx)
	articleContentQuery = dao.decorateQueryFilter(articleContentQuery, queryFilter)
	articleContentQuery = dao.decorateSortQueryFilter(articleContentQuery, sortFilter)
	return articleContentQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *ArticleContentDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleContentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleContent, error) {
	query := dao.query.ArticleContent
	articleContentQuery := query.WithContext(ctx)
	articleContentQuery = dao.decorateQueryFilter(articleContentQuery, queryFilter)
	articleContentQuery = dao.decorateSortQueryFilter(articleContentQuery, sortFilter)
	return articleContentQuery.Find()
}

func (dao *ArticleContentDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleContentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.ArticleContent, error) {
	query := dao.query.ArticleContent
	articleContentQuery := query.WithContext(ctx)
	articleContentQuery = dao.decorateQueryFilter(articleContentQuery, queryFilter)
	articleContentQuery = dao.decorateSortQueryFilter(articleContentQuery, sortFilter)
	return articleContentQuery.First()
}
