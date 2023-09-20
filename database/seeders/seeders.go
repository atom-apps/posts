package seeders

import "github.com/rogeecn/atom/contracts"

var Seeders = []contracts.SeederProvider{
	NewTagSeeder,
	NewBookSeeder,
	NewChapterSeeder,
	NewArticleSeeder,
}
