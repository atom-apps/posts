// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/atom-apps/posts/database/models"
)

func newArticleForwardSource(db *gorm.DB, opts ...gen.DOOption) articleForwardSource {
	_articleForwardSource := articleForwardSource{}

	_articleForwardSource.articleForwardSourceDo.UseDB(db, opts...)
	_articleForwardSource.articleForwardSourceDo.UseModel(&models.ArticleForwardSource{})

	tableName := _articleForwardSource.articleForwardSourceDo.TableName()
	_articleForwardSource.ALL = field.NewAsterisk(tableName)
	_articleForwardSource.ID = field.NewUint64(tableName, "id")
	_articleForwardSource.CreatedAt = field.NewTime(tableName, "created_at")
	_articleForwardSource.ArticleID = field.NewUint64(tableName, "article_id")
	_articleForwardSource.Source = field.NewString(tableName, "source")
	_articleForwardSource.SourceAuthor = field.NewString(tableName, "source_author")

	_articleForwardSource.fillFieldMap()

	return _articleForwardSource
}

type articleForwardSource struct {
	articleForwardSourceDo articleForwardSourceDo

	ALL          field.Asterisk
	ID           field.Uint64 // ID
	CreatedAt    field.Time   // 创建时间
	ArticleID    field.Uint64 // 文章ID
	Source       field.String // 原文地址
	SourceAuthor field.String // 原文作者

	fieldMap map[string]field.Expr
}

func (a articleForwardSource) Table(newTableName string) *articleForwardSource {
	a.articleForwardSourceDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleForwardSource) As(alias string) *articleForwardSource {
	a.articleForwardSourceDo.DO = *(a.articleForwardSourceDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleForwardSource) updateTableName(table string) *articleForwardSource {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.ArticleID = field.NewUint64(table, "article_id")
	a.Source = field.NewString(table, "source")
	a.SourceAuthor = field.NewString(table, "source_author")

	a.fillFieldMap()

	return a
}

func (a *articleForwardSource) WithContext(ctx context.Context) IArticleForwardSourceDo {
	return a.articleForwardSourceDo.WithContext(ctx)
}

func (a articleForwardSource) TableName() string { return a.articleForwardSourceDo.TableName() }

func (a articleForwardSource) Alias() string { return a.articleForwardSourceDo.Alias() }

func (a articleForwardSource) Columns(cols ...field.Expr) gen.Columns {
	return a.articleForwardSourceDo.Columns(cols...)
}

func (a *articleForwardSource) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleForwardSource) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["source"] = a.Source
	a.fieldMap["source_author"] = a.SourceAuthor
}

func (a articleForwardSource) clone(db *gorm.DB) articleForwardSource {
	a.articleForwardSourceDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleForwardSource) replaceDB(db *gorm.DB) articleForwardSource {
	a.articleForwardSourceDo.ReplaceDB(db)
	return a
}

type articleForwardSourceDo struct{ gen.DO }

type IArticleForwardSourceDo interface {
	gen.SubQuery
	Debug() IArticleForwardSourceDo
	WithContext(ctx context.Context) IArticleForwardSourceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleForwardSourceDo
	WriteDB() IArticleForwardSourceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleForwardSourceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleForwardSourceDo
	Not(conds ...gen.Condition) IArticleForwardSourceDo
	Or(conds ...gen.Condition) IArticleForwardSourceDo
	Select(conds ...field.Expr) IArticleForwardSourceDo
	Where(conds ...gen.Condition) IArticleForwardSourceDo
	Order(conds ...field.Expr) IArticleForwardSourceDo
	Distinct(cols ...field.Expr) IArticleForwardSourceDo
	Omit(cols ...field.Expr) IArticleForwardSourceDo
	Join(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo
	Group(cols ...field.Expr) IArticleForwardSourceDo
	Having(conds ...gen.Condition) IArticleForwardSourceDo
	Limit(limit int) IArticleForwardSourceDo
	Offset(offset int) IArticleForwardSourceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleForwardSourceDo
	Unscoped() IArticleForwardSourceDo
	Create(values ...*models.ArticleForwardSource) error
	CreateInBatches(values []*models.ArticleForwardSource, batchSize int) error
	Save(values ...*models.ArticleForwardSource) error
	First() (*models.ArticleForwardSource, error)
	Take() (*models.ArticleForwardSource, error)
	Last() (*models.ArticleForwardSource, error)
	Find() ([]*models.ArticleForwardSource, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleForwardSource, err error)
	FindInBatches(result *[]*models.ArticleForwardSource, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ArticleForwardSource) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleForwardSourceDo
	Assign(attrs ...field.AssignExpr) IArticleForwardSourceDo
	Joins(fields ...field.RelationField) IArticleForwardSourceDo
	Preload(fields ...field.RelationField) IArticleForwardSourceDo
	FirstOrInit() (*models.ArticleForwardSource, error)
	FirstOrCreate() (*models.ArticleForwardSource, error)
	FindByPage(offset int, limit int) (result []*models.ArticleForwardSource, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleForwardSourceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleForwardSourceDo) Debug() IArticleForwardSourceDo {
	return a.withDO(a.DO.Debug())
}

func (a articleForwardSourceDo) WithContext(ctx context.Context) IArticleForwardSourceDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleForwardSourceDo) ReadDB() IArticleForwardSourceDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleForwardSourceDo) WriteDB() IArticleForwardSourceDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleForwardSourceDo) Session(config *gorm.Session) IArticleForwardSourceDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleForwardSourceDo) Clauses(conds ...clause.Expression) IArticleForwardSourceDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleForwardSourceDo) Returning(value interface{}, columns ...string) IArticleForwardSourceDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleForwardSourceDo) Not(conds ...gen.Condition) IArticleForwardSourceDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleForwardSourceDo) Or(conds ...gen.Condition) IArticleForwardSourceDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleForwardSourceDo) Select(conds ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleForwardSourceDo) Where(conds ...gen.Condition) IArticleForwardSourceDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleForwardSourceDo) Order(conds ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleForwardSourceDo) Distinct(cols ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleForwardSourceDo) Omit(cols ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleForwardSourceDo) Join(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleForwardSourceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleForwardSourceDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleForwardSourceDo) Group(cols ...field.Expr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleForwardSourceDo) Having(conds ...gen.Condition) IArticleForwardSourceDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleForwardSourceDo) Limit(limit int) IArticleForwardSourceDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleForwardSourceDo) Offset(offset int) IArticleForwardSourceDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleForwardSourceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleForwardSourceDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleForwardSourceDo) Unscoped() IArticleForwardSourceDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleForwardSourceDo) Create(values ...*models.ArticleForwardSource) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleForwardSourceDo) CreateInBatches(values []*models.ArticleForwardSource, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleForwardSourceDo) Save(values ...*models.ArticleForwardSource) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleForwardSourceDo) First() (*models.ArticleForwardSource, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleForwardSource), nil
	}
}

func (a articleForwardSourceDo) Take() (*models.ArticleForwardSource, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleForwardSource), nil
	}
}

func (a articleForwardSourceDo) Last() (*models.ArticleForwardSource, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleForwardSource), nil
	}
}

func (a articleForwardSourceDo) Find() ([]*models.ArticleForwardSource, error) {
	result, err := a.DO.Find()
	return result.([]*models.ArticleForwardSource), err
}

func (a articleForwardSourceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleForwardSource, err error) {
	buf := make([]*models.ArticleForwardSource, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleForwardSourceDo) FindInBatches(result *[]*models.ArticleForwardSource, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleForwardSourceDo) Attrs(attrs ...field.AssignExpr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleForwardSourceDo) Assign(attrs ...field.AssignExpr) IArticleForwardSourceDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleForwardSourceDo) Joins(fields ...field.RelationField) IArticleForwardSourceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleForwardSourceDo) Preload(fields ...field.RelationField) IArticleForwardSourceDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleForwardSourceDo) FirstOrInit() (*models.ArticleForwardSource, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleForwardSource), nil
	}
}

func (a articleForwardSourceDo) FirstOrCreate() (*models.ArticleForwardSource, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleForwardSource), nil
	}
}

func (a articleForwardSourceDo) FindByPage(offset int, limit int) (result []*models.ArticleForwardSource, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleForwardSourceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleForwardSourceDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleForwardSourceDo) Delete(models ...*models.ArticleForwardSource) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleForwardSourceDo) withDO(do gen.Dao) *articleForwardSourceDo {
	a.DO = *do.(*gen.DO)
	return a
}
