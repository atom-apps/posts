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
type ArticleAttachmentDao struct {
	query *query.Query
}

func (dao *ArticleAttachmentDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *ArticleAttachmentDao) Context(ctx context.Context) query.IArticleAttachmentDo {
	return dao.query.ArticleAttachment.WithContext(ctx)
}

func (dao *ArticleAttachmentDao) decorateSortQueryFilter(query query.IArticleAttachmentDo, sortFilter *common.SortQueryFilter) query.IArticleAttachmentDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.ArticleAttachment.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.ArticleAttachment.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *ArticleAttachmentDao) decorateQueryFilter(query query.IArticleAttachmentDo, queryFilter *dto.ArticleAttachmentListQueryFilter) query.IArticleAttachmentDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.TenantID != nil {
		query = query.Where(dao.query.ArticleAttachment.TenantID.Eq(*queryFilter.TenantID))
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.ArticleAttachment.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.ArticleID != nil {
		query = query.Where(dao.query.ArticleAttachment.ArticleID.Eq(*queryFilter.ArticleID))
	}
	if queryFilter.FilesystemID != nil {
		query = query.Where(dao.query.ArticleAttachment.FilesystemID.Eq(*queryFilter.FilesystemID))
	}
	if queryFilter.Description != nil {
		query = query.Where(dao.query.ArticleAttachment.Description.Eq(*queryFilter.Description))
	}
	if queryFilter.Password != nil {
		query = query.Where(dao.query.ArticleAttachment.Password.Eq(*queryFilter.Password))
	}

	return query
}

func (dao *ArticleAttachmentDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleAttachmentDao) Update(ctx context.Context, model *models.ArticleAttachment) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleAttachmentDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleAttachmentDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.ArticleAttachment.ID.Eq(id)).Delete()
	return err
}

func (dao *ArticleAttachmentDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.ArticleAttachment.ID.Eq(id)).UpdateSimple(dao.query.ArticleAttachment.DeletedAt.Null())
	return err
}

func (dao *ArticleAttachmentDao) Create(ctx context.Context, model *models.ArticleAttachment) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleAttachmentDao) GetByID(ctx context.Context, id uint64) (*models.ArticleAttachment, error) {
	return dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.Eq(id)).First()
}

func (dao *ArticleAttachmentDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleAttachment, error) {
	return dao.Context(ctx).Where(dao.query.ArticleAttachment.ID.In(ids...)).Find()
}

func (dao *ArticleAttachmentDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleAttachmentListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleAttachment, int64, error) {
	query := dao.query.ArticleAttachment
	articleAttachmentQuery := query.WithContext(ctx)
	articleAttachmentQuery = dao.decorateQueryFilter(articleAttachmentQuery, queryFilter)
	articleAttachmentQuery = dao.decorateSortQueryFilter(articleAttachmentQuery, sortFilter)
	return articleAttachmentQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *ArticleAttachmentDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleAttachmentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleAttachment, error) {
	query := dao.query.ArticleAttachment
	articleAttachmentQuery := query.WithContext(ctx)
	articleAttachmentQuery = dao.decorateQueryFilter(articleAttachmentQuery, queryFilter)
	articleAttachmentQuery = dao.decorateSortQueryFilter(articleAttachmentQuery, sortFilter)
	return articleAttachmentQuery.Find()
}

func (dao *ArticleAttachmentDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleAttachmentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.ArticleAttachment, error) {
	query := dao.query.ArticleAttachment
	articleAttachmentQuery := query.WithContext(ctx)
	articleAttachmentQuery = dao.decorateQueryFilter(articleAttachmentQuery, queryFilter)
	articleAttachmentQuery = dao.decorateSortQueryFilter(articleAttachmentQuery, sortFilter)
	return articleAttachmentQuery.First()
}
