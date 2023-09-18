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

func newArticleDig(db *gorm.DB, opts ...gen.DOOption) articleDig {
	_articleDig := articleDig{}

	_articleDig.articleDigDo.UseDB(db, opts...)
	_articleDig.articleDigDo.UseModel(&models.ArticleDig{})

	tableName := _articleDig.articleDigDo.TableName()
	_articleDig.ALL = field.NewAsterisk(tableName)
	_articleDig.ID = field.NewUint64(tableName, "id")
	_articleDig.CreatedAt = field.NewTime(tableName, "created_at")
	_articleDig.ArticleID = field.NewUint64(tableName, "article_id")
	_articleDig.Views = field.NewUint64(tableName, "views")
	_articleDig.Likes = field.NewUint64(tableName, "likes")
	_articleDig.Dislikes = field.NewUint64(tableName, "dislikes")

	_articleDig.fillFieldMap()

	return _articleDig
}

type articleDig struct {
	articleDigDo articleDigDo

	ALL       field.Asterisk
	ID        field.Uint64 // ID
	CreatedAt field.Time   // 创建时间
	ArticleID field.Uint64 // 文章ID
	Views     field.Uint64 // 浏览量
	Likes     field.Uint64 // 喜欢
	Dislikes  field.Uint64 // 不喜欢

	fieldMap map[string]field.Expr
}

func (a articleDig) Table(newTableName string) *articleDig {
	a.articleDigDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleDig) As(alias string) *articleDig {
	a.articleDigDo.DO = *(a.articleDigDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleDig) updateTableName(table string) *articleDig {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.ArticleID = field.NewUint64(table, "article_id")
	a.Views = field.NewUint64(table, "views")
	a.Likes = field.NewUint64(table, "likes")
	a.Dislikes = field.NewUint64(table, "dislikes")

	a.fillFieldMap()

	return a
}

func (a *articleDig) WithContext(ctx context.Context) IArticleDigDo {
	return a.articleDigDo.WithContext(ctx)
}

func (a articleDig) TableName() string { return a.articleDigDo.TableName() }

func (a articleDig) Alias() string { return a.articleDigDo.Alias() }

func (a articleDig) Columns(cols ...field.Expr) gen.Columns { return a.articleDigDo.Columns(cols...) }

func (a *articleDig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleDig) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["views"] = a.Views
	a.fieldMap["likes"] = a.Likes
	a.fieldMap["dislikes"] = a.Dislikes
}

func (a articleDig) clone(db *gorm.DB) articleDig {
	a.articleDigDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleDig) replaceDB(db *gorm.DB) articleDig {
	a.articleDigDo.ReplaceDB(db)
	return a
}

type articleDigDo struct{ gen.DO }

type IArticleDigDo interface {
	gen.SubQuery
	Debug() IArticleDigDo
	WithContext(ctx context.Context) IArticleDigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleDigDo
	WriteDB() IArticleDigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleDigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleDigDo
	Not(conds ...gen.Condition) IArticleDigDo
	Or(conds ...gen.Condition) IArticleDigDo
	Select(conds ...field.Expr) IArticleDigDo
	Where(conds ...gen.Condition) IArticleDigDo
	Order(conds ...field.Expr) IArticleDigDo
	Distinct(cols ...field.Expr) IArticleDigDo
	Omit(cols ...field.Expr) IArticleDigDo
	Join(table schema.Tabler, on ...field.Expr) IArticleDigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleDigDo
	Group(cols ...field.Expr) IArticleDigDo
	Having(conds ...gen.Condition) IArticleDigDo
	Limit(limit int) IArticleDigDo
	Offset(offset int) IArticleDigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDigDo
	Unscoped() IArticleDigDo
	Create(values ...*models.ArticleDig) error
	CreateInBatches(values []*models.ArticleDig, batchSize int) error
	Save(values ...*models.ArticleDig) error
	First() (*models.ArticleDig, error)
	Take() (*models.ArticleDig, error)
	Last() (*models.ArticleDig, error)
	Find() ([]*models.ArticleDig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleDig, err error)
	FindInBatches(result *[]*models.ArticleDig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ArticleDig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleDigDo
	Assign(attrs ...field.AssignExpr) IArticleDigDo
	Joins(fields ...field.RelationField) IArticleDigDo
	Preload(fields ...field.RelationField) IArticleDigDo
	FirstOrInit() (*models.ArticleDig, error)
	FirstOrCreate() (*models.ArticleDig, error)
	FindByPage(offset int, limit int) (result []*models.ArticleDig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleDigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleDigDo) Debug() IArticleDigDo {
	return a.withDO(a.DO.Debug())
}

func (a articleDigDo) WithContext(ctx context.Context) IArticleDigDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleDigDo) ReadDB() IArticleDigDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleDigDo) WriteDB() IArticleDigDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleDigDo) Session(config *gorm.Session) IArticleDigDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleDigDo) Clauses(conds ...clause.Expression) IArticleDigDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleDigDo) Returning(value interface{}, columns ...string) IArticleDigDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleDigDo) Not(conds ...gen.Condition) IArticleDigDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleDigDo) Or(conds ...gen.Condition) IArticleDigDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleDigDo) Select(conds ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleDigDo) Where(conds ...gen.Condition) IArticleDigDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleDigDo) Order(conds ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleDigDo) Distinct(cols ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleDigDo) Omit(cols ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleDigDo) Join(table schema.Tabler, on ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleDigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleDigDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleDigDo) Group(cols ...field.Expr) IArticleDigDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleDigDo) Having(conds ...gen.Condition) IArticleDigDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleDigDo) Limit(limit int) IArticleDigDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleDigDo) Offset(offset int) IArticleDigDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleDigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleDigDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleDigDo) Unscoped() IArticleDigDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleDigDo) Create(values ...*models.ArticleDig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleDigDo) CreateInBatches(values []*models.ArticleDig, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleDigDo) Save(values ...*models.ArticleDig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleDigDo) First() (*models.ArticleDig, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleDig), nil
	}
}

func (a articleDigDo) Take() (*models.ArticleDig, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleDig), nil
	}
}

func (a articleDigDo) Last() (*models.ArticleDig, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleDig), nil
	}
}

func (a articleDigDo) Find() ([]*models.ArticleDig, error) {
	result, err := a.DO.Find()
	return result.([]*models.ArticleDig), err
}

func (a articleDigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleDig, err error) {
	buf := make([]*models.ArticleDig, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleDigDo) FindInBatches(result *[]*models.ArticleDig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleDigDo) Attrs(attrs ...field.AssignExpr) IArticleDigDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleDigDo) Assign(attrs ...field.AssignExpr) IArticleDigDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleDigDo) Joins(fields ...field.RelationField) IArticleDigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleDigDo) Preload(fields ...field.RelationField) IArticleDigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleDigDo) FirstOrInit() (*models.ArticleDig, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleDig), nil
	}
}

func (a articleDigDo) FirstOrCreate() (*models.ArticleDig, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleDig), nil
	}
}

func (a articleDigDo) FindByPage(offset int, limit int) (result []*models.ArticleDig, count int64, err error) {
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

func (a articleDigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleDigDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleDigDo) Delete(models ...*models.ArticleDig) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleDigDo) withDO(do gen.Dao) *articleDigDo {
	a.DO = *do.(*gen.DO)
	return a
}