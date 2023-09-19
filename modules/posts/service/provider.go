package service

import (
	"github.com/atom-apps/posts/modules/posts/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		articleAttachmentDao *dao.ArticleAttachmentDao,
	) (*ArticleAttachmentService, error) {
		obj := &ArticleAttachmentService{
			articleAttachmentDao: articleAttachmentDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleContentDao *dao.ArticleContentDao,
	) (*ArticleContentService, error) {
		obj := &ArticleContentService{
			articleContentDao: articleContentDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDigDao *dao.ArticleDigDao,
	) (*ArticleDigService, error) {
		obj := &ArticleDigService{
			articleDigDao: articleDigDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDao *dao.ArticleDao,
	) (*ArticleService, error) {
		obj := &ArticleService{
			articleDao: articleDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDao *dao.ArticleDao,
		bookDao *dao.BookDao,
		chapterDao *dao.ChapterDao,
	) (*BookService, error) {
		obj := &BookService{
			articleDao: articleDao,
			bookDao:    bookDao,
			chapterDao: chapterDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		chapterDao *dao.ChapterDao,
	) (*ChapterService, error) {
		obj := &ChapterService{
			chapterDao: chapterDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
