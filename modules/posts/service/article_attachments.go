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
type ArticleAttachmentService struct {
	articleAttachmentDao *dao.ArticleAttachmentDao
}

func (svc *ArticleAttachmentService) DecorateItem(model *models.ArticleAttachment, id int) *dto.ArticleAttachmentItem {
	return &dto.ArticleAttachmentItem{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		TenantID:     model.TenantID,
		UserID:       model.UserID,
		ArticleID:    model.ArticleID,
		FilesystemID: model.FilesystemID,
		Description:  model.Description,
		Password:     model.Password,
	}
}

func (svc *ArticleAttachmentService) GetByID(ctx context.Context, id uint64) (*models.ArticleAttachment, error) {
	return svc.articleAttachmentDao.GetByID(ctx, id)
}

func (svc *ArticleAttachmentService) FindByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleAttachmentListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleAttachment, error) {
	return svc.articleAttachmentDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *ArticleAttachmentService) PageByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleAttachmentListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleAttachment, int64, error) {
	return svc.articleAttachmentDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ArticleAttachmentService) CreateFromModel(ctx context.Context, model *models.ArticleAttachment) error {
	return svc.articleAttachmentDao.Create(ctx, model)
}

// Create
func (svc *ArticleAttachmentService) Create(ctx context.Context, articleId int, body *dto.ArticleAttachmentForm) error {
	model := &models.ArticleAttachment{}
	_ = copier.Copy(model, body)
	return svc.articleAttachmentDao.Create(ctx, model)
}

// Update
func (svc *ArticleAttachmentService) Update(ctx context.Context, articleId int, id uint64, body *dto.ArticleAttachmentForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.articleAttachmentDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ArticleAttachmentService) UpdateFromModel(ctx context.Context, model *models.ArticleAttachment) error {
	return svc.articleAttachmentDao.Update(ctx, model)
}

// Delete
func (svc *ArticleAttachmentService) Delete(ctx context.Context, articleId int, id uint64) error {
	return svc.articleAttachmentDao.Delete(ctx, id)
}
