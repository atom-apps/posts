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
type ChapterService struct {
	chapterDao *dao.ChapterDao
}

func (svc *ChapterService) DecorateItem(model *models.Chapter, id int) *dto.ChapterItem {
	return &dto.ChapterItem{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		TenantID:    model.TenantID,
		UserID:      model.UserID,
		BookID:      model.BookID,
		Title:       model.Title,
		Description: model.Description,
		Content:     model.Content,
	}
}

func (svc *ChapterService) GetByID(ctx context.Context, id uint64) (*models.Chapter, error) {
	return svc.chapterDao.GetByID(ctx, id)
}

func (svc *ChapterService) FindByQueryFilter(
	ctx context.Context,
	bookId int,
	queryFilter *dto.ChapterListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Chapter, error) {
	return svc.chapterDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *ChapterService) PageByQueryFilter(
	ctx context.Context,
	bookId int,
	queryFilter *dto.ChapterListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Chapter, int64, error) {
	return svc.chapterDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ChapterService) CreateFromModel(ctx context.Context, model *models.Chapter) error {
	return svc.chapterDao.Create(ctx, model)
}

// Create
func (svc *ChapterService) Create(ctx context.Context, bookId int, body *dto.ChapterForm) error {
	model := &models.Chapter{}
	_ = copier.Copy(model, body)
	return svc.chapterDao.Create(ctx, model)
}

// Update
func (svc *ChapterService) Update(ctx context.Context, bookId int, id uint64, body *dto.ChapterForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.chapterDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ChapterService) UpdateFromModel(ctx context.Context, model *models.Chapter) error {
	return svc.chapterDao.Update(ctx, model)
}

// Delete
func (svc *ChapterService) Delete(ctx context.Context, bookId int, id uint64) error {
	return svc.chapterDao.Delete(ctx, id)
}
