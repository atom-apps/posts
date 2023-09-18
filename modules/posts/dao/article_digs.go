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
type ArticleDigDao struct {
	query *query.Query
}

func (dao *ArticleDigDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleDigDao) Context(ctx context.Context) query.IArticleDigDo {
	return dao.query.ArticleDig.WithContext(ctx)
}

func (dao *ArticleDigDao) decorateSortQueryFilter(query query.IArticleDigDo, sortFilter *common.SortQueryFilter) query.IArticleDigDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.ArticleDig.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.ArticleDig.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *ArticleDigDao) decorateQueryFilter(query query.IArticleDigDo, queryFilter *dto.ArticleDigListQueryFilter) query.IArticleDigDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.ArticleID != nil {
		query = query.Where(dao.query.ArticleDig.ArticleID.Eq(*queryFilter.ArticleID))
	}
	if queryFilter.Views != nil {
		query = query.Where(dao.query.ArticleDig.Views.Eq(*queryFilter.Views))
	}
	if queryFilter.Likes != nil {
		query = query.Where(dao.query.ArticleDig.Likes.Eq(*queryFilter.Likes))
	}
	if queryFilter.Dislikes != nil {
		query = query.Where(dao.query.ArticleDig.Dislikes.Eq(*queryFilter.Dislikes))
	}

	return query
}

func (dao *ArticleDigDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleDigDao) Update(ctx context.Context, model *models.ArticleDig) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleDigDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDigDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.ArticleDig.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleDigDao) Create(ctx context.Context, model *models.ArticleDig) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleDigDao) GetByID(ctx context.Context, id uint64) (*models.ArticleDig, error) {
	return dao.Context(ctx).Where(dao.query.ArticleDig.ID.Eq(id)).First()
}

func (dao *ArticleDigDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleDig, error) {
	return dao.Context(ctx).Where(dao.query.ArticleDig.ID.In(ids...)).Find()
}

func (dao *ArticleDigDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleDigListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleDig, int64, error) {
	query := dao.query.ArticleDig
	articleDigQuery := query.WithContext(ctx)
	articleDigQuery = dao.decorateQueryFilter(articleDigQuery, queryFilter)
	articleDigQuery = dao.decorateSortQueryFilter(articleDigQuery, sortFilter)
	return articleDigQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *ArticleDigDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleDigListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleDig, error) {
	query := dao.query.ArticleDig
	articleDigQuery := query.WithContext(ctx)
	articleDigQuery = dao.decorateQueryFilter(articleDigQuery, queryFilter)
	articleDigQuery = dao.decorateSortQueryFilter(articleDigQuery, sortFilter)
	return articleDigQuery.Find()
}

func (dao *ArticleDigDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleDigListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.ArticleDig, error) {
	query := dao.query.ArticleDig
	articleDigQuery := query.WithContext(ctx)
	articleDigQuery = dao.decorateQueryFilter(articleDigQuery, queryFilter)
	articleDigQuery = dao.decorateSortQueryFilter(articleDigQuery, sortFilter)
	return articleDigQuery.First()
}
