package controller

import (
	"github.com/atom-apps/posts/modules/posts/service"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		articleContentSvc *service.ArticleContentService,
	) (*ArticleContentController, error) {
		obj := &ArticleContentController{
			articleContentSvc: articleContentSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDigSvc *service.ArticleDigService,
	) (*ArticleDigController, error) {
		obj := &ArticleDigController{
			articleDigSvc: articleDigSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleSvc *service.ArticleService,
	) (*ArticleController, error) {
		obj := &ArticleController{
			articleSvc: articleSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		bookSvc *service.BookService,
	) (*BookController, error) {
		obj := &BookController{
			bookSvc: bookSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		chapterSvc *service.ChapterService,
	) (*ChapterController, error) {
		obj := &ChapterController{
			chapterSvc: chapterSvc,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
