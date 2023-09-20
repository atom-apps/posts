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
type ArticleService struct {
	articleDao       *dao.ArticleDao
	attachmentDao    *dao.ArticleAttachmentDao
	contentDao       *dao.ArticleContentDao
	digDao           *dao.ArticleDigDao
	forwardSourceDao *dao.ArticleForwardSourceDao
	paidUserDao      *dao.ArticlePaidUserDao
	paymentDao       *dao.ArticlePaymentDao
}

func (svc *ArticleService) DecorateItem(model *models.Article, id int) *dto.ArticleItem {
	return &dto.ArticleItem{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		TenantID:    model.TenantID,
		UserID:      model.UserID,
		UUID:        model.UUID,
		BookID:      model.BookID,
		ChapterID:   model.ChapterID,
		CategoryID:  model.CategoryID,
		PublishAt:   model.PublishAt,
		Type:        model.Type,
		TypeCN:      model.Type.Cn(),
		Format:      model.Format,
		FormatCN:    model.Format.Cn(),
		Title:       model.Title,
		Keyword:     model.Keyword,
		Description: model.Description,
		Thumbnails:  model.Thumbnails,
		Videos:      model.Videos,
		Audios:      model.Audios,
		PostIP:      model.PostIP,
		Weight:      model.Weight,
	}
}

func (svc *ArticleService) GetByID(ctx context.Context, id uint64) (*models.Article, error) {
	return svc.articleDao.GetByID(ctx, id)
}

func (svc *ArticleService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Article, error) {
	return svc.articleDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *ArticleService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.ArticleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Article, int64, error) {
	return svc.articleDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ArticleService) CreateFromModel(ctx context.Context, model *models.Article) error {
	return svc.articleDao.Create(ctx, model)
}

// Create
func (svc *ArticleService) Create(ctx context.Context, body *dto.ArticleForm) error {
	model := &models.Article{}
	_ = copier.Copy(model, body)
	return svc.articleDao.Create(ctx, model)
}

// Update
func (svc *ArticleService) Update(ctx context.Context, id uint64, body *dto.ArticleForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.articleDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ArticleService) UpdateFromModel(ctx context.Context, model *models.Article) error {
	return svc.articleDao.Update(ctx, model)
}

// Delete
func (svc *ArticleService) Delete(ctx context.Context, id uint64) error {
	return svc.articleDao.Delete(ctx, id)
}
