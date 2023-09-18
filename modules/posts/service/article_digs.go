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
type ArticleDigService struct {
	articleDigDao *dao.ArticleDigDao
}

func (svc *ArticleDigService) DecorateItem(model *models.ArticleDig, id int) *dto.ArticleDigItem {
	return &dto.ArticleDigItem{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		ArticleID: model.ArticleID,
		Views:     model.Views,
		Likes:     model.Likes,
		Dislikes:  model.Dislikes,
	}
}

func (svc *ArticleDigService) GetByID(ctx context.Context, id uint64) (*models.ArticleDig, error) {
	return svc.articleDigDao.GetByID(ctx, id)
}

func (svc *ArticleDigService) FindByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleDigListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleDig, error) {
	return svc.articleDigDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *ArticleDigService) PageByQueryFilter(
	ctx context.Context,
	articleId int,
	queryFilter *dto.ArticleDigListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.ArticleDig, int64, error) {
	return svc.articleDigDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *ArticleDigService) CreateFromModel(ctx context.Context, model *models.ArticleDig) error {
	return svc.articleDigDao.Create(ctx, model)
}

// Create
func (svc *ArticleDigService) Create(ctx context.Context, articleId int, body *dto.ArticleDigForm) error {
	model := &models.ArticleDig{}
	_ = copier.Copy(model, body)
	return svc.articleDigDao.Create(ctx, model)
}

// Update
func (svc *ArticleDigService) Update(ctx context.Context, articleId int, id uint64, body *dto.ArticleDigForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.articleDigDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *ArticleDigService) UpdateFromModel(ctx context.Context, model *models.ArticleDig) error {
	return svc.articleDigDao.Update(ctx, model)
}

// Delete
func (svc *ArticleDigService) Delete(ctx context.Context, articleId int, id uint64) error {
	return svc.articleDigDao.Delete(ctx, id)
}
