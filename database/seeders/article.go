package seeders

import (
	"github.com/atom-apps/posts/common"
	"github.com/atom-apps/posts/common/consts"
	"github.com/atom-apps/posts/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type ArticleSeeder struct{}

func NewArticleSeeder() contracts.Seeder {
	return &ArticleSeeder{}
}

func (s *ArticleSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Article{}).TableName(nil))

	times := 100
	ms := lo.Times(times, func(index int) *models.Article {
		data := s.Generate(faker, index, true)
		return &data
	})
	// db.CreateInBatches(&ms, 100)
	lo.ForEach(ms, func(item *models.Article, index int) {
		db.Create(item)
	})

	ms = lo.Times(times, func(index int) *models.Article {
		data := s.Generate(faker, index, false)
		return &data
	})
	// db.CreateInBatches(&ms, 100)
	lo.ForEach(ms, func(item *models.Article, index int) {
		db.Create(item)
	})
}

func (s *ArticleSeeder) Generate(faker *gofakeit.Faker, idx int, isBook bool) models.Article {
	bookID := 0
	if isBook {
		bookID = faker.Number(1, 50)
	}
	chapterID := 0
	if isBook {
		chapterID = faker.Number(1, 200)
	}

	return models.Article{
		TenantID:    1,
		UserID:      1,
		UUID:        faker.UUID(),
		BookID:      uint64(bookID),
		ChapterID:   uint64(chapterID),
		CategoryID:  faker.Uint64() % 10,
		PublishAt:   faker.FutureDate(),
		Type:        consts.ArticleTypeValues()[faker.Number(0, len(consts.ArticleTypeValues())-1)],
		Format:      consts.ArticleFormatValues()[faker.Number(0, len(consts.ArticleFormatValues())-1)],
		Title:       faker.Book().Title,
		Keyword:     faker.StreetName(),
		Description: faker.StreetName(),
		Thumbnails: common.ArticleThumbnails{
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
		},
		Videos: common.ArticleVideos{
			{Provider: "", Url: ""},
		},
		Audios: common.ArticleAudios{
			{Provider: "", Url: ""},
		},
		PostIP: faker.IPv4Address(),
		Weight: faker.Uint64(),
	}
}
