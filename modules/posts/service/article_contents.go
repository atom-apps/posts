package service

import (
	"context"

	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/database/models"
	"github.com/atom-apps/posts/modules/posts/dao"
	"github.com/atom-apps/posts/modules/posts/dto"

	"github.com/jinzhu/copier"
)

// @provider
type ArticleContentService struct {
	articleContentDao *dao.ArticleContentDao
}

func (svc *ArticleContentService) DecorateItem(model *models.ArticleContent, id int) *dto.ArticleContentItem {
	return &dto.ArticleContentItem{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		TenantID:     model.TenantID,
		UserID:       model.UserID,
		ArticleID:    model.ArticleID,
		FreeContent:  model.FreeContent,
		PriceContent: model.PriceContent,
	}
}

func (svc *ArticleContentService) GetByID(ctx context.Context, id uint64) (*models.ArticleContent, error) {
	return svc.articleContentDao.GetByID(ctx, id)
}

func (svc *ArticleContentService) FindByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleContentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleContent, error) {
	return svc.articleContentDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *ArticleContentService) PageByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleContentListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleContent, int64, error) {
	return svc.articleContentDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ArticleContentService) CreateFromModel(ctx context.Context, model *models.ArticleContent) error {
	return svc.articleContentDao.Create(ctx, model)
}

// Create
func (svc *ArticleContentService) Create(ctx context.Context, articleId int, body *dto.ArticleContentForm) error {
	model := &models.ArticleContent{}
	_ = copier.Copy(model, body)
	return svc.articleContentDao.Create(ctx, model)
}

// Update
func (svc *ArticleContentService) Update(ctx context.Context, articleId int, id uint64, body *dto.ArticleContentForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.articleContentDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ArticleContentService) UpdateFromModel(ctx context.Context, model *models.ArticleContent) error {
	return svc.articleContentDao.Update(ctx, model)
}

// Delete
func (svc *ArticleContentService) Delete(ctx context.Context, articleId int, id uint64) error {
	return svc.articleContentDao.Delete(ctx, id)
}
