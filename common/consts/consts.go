package consts

// swagger:enum ArticleType
// ENUM(
// Text = 0,
// Image = 1,
// Audio = 2,
// Video = 3,
// )
type ArticleType int

// swagger:enum ArticleFormat
// ENUM(
// Html = 0,
// Markdown = 1,
// )
type ArticleFormat int

// swagger:enum PriceType
// ENUM(
// Content = 0,
// Attachment = 1,
// )
type ArticlePriceType int

// swagger:enum VideoProvider
// ENUM(url, bilibili, tencent, youku, youtube, iqiyi, sohu, xigua, douyin, weishi, weibo)
type VideoProvider string

// swagger:enum AudioProvider
// ENUM(url, ximalaya)
type AudioProvider string
