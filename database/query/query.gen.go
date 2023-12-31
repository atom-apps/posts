// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	Article              *article
	ArticleAttachment    *articleAttachment
	ArticleContent       *articleContent
	ArticleDig           *articleDig
	ArticleForwardSource *articleForwardSource
	ArticlePaidUser      *articlePaidUser
	ArticlePayment       *articlePayment
	ArticleTag           *articleTag
	Book                 *book
	Chapter              *chapter
	Migration            *migration
	Tag                  *tag
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Article = &Q.Article
	ArticleAttachment = &Q.ArticleAttachment
	ArticleContent = &Q.ArticleContent
	ArticleDig = &Q.ArticleDig
	ArticleForwardSource = &Q.ArticleForwardSource
	ArticlePaidUser = &Q.ArticlePaidUser
	ArticlePayment = &Q.ArticlePayment
	ArticleTag = &Q.ArticleTag
	Book = &Q.Book
	Chapter = &Q.Chapter
	Migration = &Q.Migration
	Tag = &Q.Tag
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		Article:              newArticle(db, opts...),
		ArticleAttachment:    newArticleAttachment(db, opts...),
		ArticleContent:       newArticleContent(db, opts...),
		ArticleDig:           newArticleDig(db, opts...),
		ArticleForwardSource: newArticleForwardSource(db, opts...),
		ArticlePaidUser:      newArticlePaidUser(db, opts...),
		ArticlePayment:       newArticlePayment(db, opts...),
		ArticleTag:           newArticleTag(db, opts...),
		Book:                 newBook(db, opts...),
		Chapter:              newChapter(db, opts...),
		Migration:            newMigration(db, opts...),
		Tag:                  newTag(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Article              article
	ArticleAttachment    articleAttachment
	ArticleContent       articleContent
	ArticleDig           articleDig
	ArticleForwardSource articleForwardSource
	ArticlePaidUser      articlePaidUser
	ArticlePayment       articlePayment
	ArticleTag           articleTag
	Book                 book
	Chapter              chapter
	Migration            migration
	Tag                  tag
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		Article:              q.Article.clone(db),
		ArticleAttachment:    q.ArticleAttachment.clone(db),
		ArticleContent:       q.ArticleContent.clone(db),
		ArticleDig:           q.ArticleDig.clone(db),
		ArticleForwardSource: q.ArticleForwardSource.clone(db),
		ArticlePaidUser:      q.ArticlePaidUser.clone(db),
		ArticlePayment:       q.ArticlePayment.clone(db),
		ArticleTag:           q.ArticleTag.clone(db),
		Book:                 q.Book.clone(db),
		Chapter:              q.Chapter.clone(db),
		Migration:            q.Migration.clone(db),
		Tag:                  q.Tag.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		Article:              q.Article.replaceDB(db),
		ArticleAttachment:    q.ArticleAttachment.replaceDB(db),
		ArticleContent:       q.ArticleContent.replaceDB(db),
		ArticleDig:           q.ArticleDig.replaceDB(db),
		ArticleForwardSource: q.ArticleForwardSource.replaceDB(db),
		ArticlePaidUser:      q.ArticlePaidUser.replaceDB(db),
		ArticlePayment:       q.ArticlePayment.replaceDB(db),
		ArticleTag:           q.ArticleTag.replaceDB(db),
		Book:                 q.Book.replaceDB(db),
		Chapter:              q.Chapter.replaceDB(db),
		Migration:            q.Migration.replaceDB(db),
		Tag:                  q.Tag.replaceDB(db),
	}
}

type queryCtx struct {
	Article              IArticleDo
	ArticleAttachment    IArticleAttachmentDo
	ArticleContent       IArticleContentDo
	ArticleDig           IArticleDigDo
	ArticleForwardSource IArticleForwardSourceDo
	ArticlePaidUser      IArticlePaidUserDo
	ArticlePayment       IArticlePaymentDo
	ArticleTag           IArticleTagDo
	Book                 IBookDo
	Chapter              IChapterDo
	Migration            IMigrationDo
	Tag                  ITagDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Article:              q.Article.WithContext(ctx),
		ArticleAttachment:    q.ArticleAttachment.WithContext(ctx),
		ArticleContent:       q.ArticleContent.WithContext(ctx),
		ArticleDig:           q.ArticleDig.WithContext(ctx),
		ArticleForwardSource: q.ArticleForwardSource.WithContext(ctx),
		ArticlePaidUser:      q.ArticlePaidUser.WithContext(ctx),
		ArticlePayment:       q.ArticlePayment.WithContext(ctx),
		ArticleTag:           q.ArticleTag.WithContext(ctx),
		Book:                 q.Book.WithContext(ctx),
		Chapter:              q.Chapter.WithContext(ctx),
		Migration:            q.Migration.WithContext(ctx),
		Tag:                  q.Tag.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
