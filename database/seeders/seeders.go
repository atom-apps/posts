package seeders

import "github.com/rogeecn/atom/contracts"

var Seeders = []contracts.SeederProvider{
	NewBookSeeder,
	NewChapterSeeder,
	NewArticleSeeder,
}
