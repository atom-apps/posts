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

func newArticleContent(db *gorm.DB, opts ...gen.DOOption) articleContent {
	_articleContent := articleContent{}

	_articleContent.articleContentDo.UseDB(db, opts...)
	_articleContent.articleContentDo.UseModel(&models.ArticleContent{})

	tableName := _articleContent.articleContentDo.TableName()
	_articleContent.ALL = field.NewAsterisk(tableName)
	_articleContent.ID = field.NewUint64(tableName, "id")
	_articleContent.CreatedAt = field.NewTime(tableName, "created_at")
	_articleContent.UpdatedAt = field.NewTime(tableName, "updated_at")
	_articleContent.DeletedAt = field.NewField(tableName, "deleted_at")
	_articleContent.TenantID = field.NewUint64(tableName, "tenant_id")
	_articleContent.UserID = field.NewUint64(tableName, "user_id")
	_articleContent.ArticleID = field.NewUint64(tableName, "article_id")
	_articleContent.FreeContent = field.NewString(tableName, "free_content")
	_articleContent.PriceContent = field.NewString(tableName, "price_content")

	_articleContent.fillFieldMap()

	return _articleContent
}

type articleContent struct {
	articleContentDo articleContentDo

	ALL          field.Asterisk
	ID           field.Uint64 // ID
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间
	TenantID     field.Uint64 // 租户ID
	UserID       field.Uint64 // 用户ID
	ArticleID    field.Uint64 // 文章ID
	FreeContent  field.String // 内容
	PriceContent field.String // 付费内容

	fieldMap map[string]field.Expr
}

func (a articleContent) Table(newTableName string) *articleContent {
	a.articleContentDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleContent) As(alias string) *articleContent {
	a.articleContentDo.DO = *(a.articleContentDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleContent) updateTableName(table string) *articleContent {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.TenantID = field.NewUint64(table, "tenant_id")
	a.UserID = field.NewUint64(table, "user_id")
	a.ArticleID = field.NewUint64(table, "article_id")
	a.FreeContent = field.NewString(table, "free_content")
	a.PriceContent = field.NewString(table, "price_content")

	a.fillFieldMap()

	return a
}

func (a *articleContent) WithContext(ctx context.Context) IArticleContentDo {
	return a.articleContentDo.WithContext(ctx)
}

func (a articleContent) TableName() string { return a.articleContentDo.TableName() }

func (a articleContent) Alias() string { return a.articleContentDo.Alias() }

func (a articleContent) Columns(cols ...field.Expr) gen.Columns {
	return a.articleContentDo.Columns(cols...)
}

func (a *articleContent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleContent) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["tenant_id"] = a.TenantID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["free_content"] = a.FreeContent
	a.fieldMap["price_content"] = a.PriceContent
}

func (a articleContent) clone(db *gorm.DB) articleContent {
	a.articleContentDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleContent) replaceDB(db *gorm.DB) articleContent {
	a.articleContentDo.ReplaceDB(db)
	return a
}

type articleContentDo struct{ gen.DO }

type IArticleContentDo interface {
	gen.SubQuery
	Debug() IArticleContentDo
	WithContext(ctx context.Context) IArticleContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleContentDo
	WriteDB() IArticleContentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleContentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleContentDo
	Not(conds ...gen.Condition) IArticleContentDo
	Or(conds ...gen.Condition) IArticleContentDo
	Select(conds ...field.Expr) IArticleContentDo
	Where(conds ...gen.Condition) IArticleContentDo
	Order(conds ...field.Expr) IArticleContentDo
	Distinct(cols ...field.Expr) IArticleContentDo
	Omit(cols ...field.Expr) IArticleContentDo
	Join(table schema.Tabler, on ...field.Expr) IArticleContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleContentDo
	Group(cols ...field.Expr) IArticleContentDo
	Having(conds ...gen.Condition) IArticleContentDo
	Limit(limit int) IArticleContentDo
	Offset(offset int) IArticleContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleContentDo
	Unscoped() IArticleContentDo
	Create(values ...*models.ArticleContent) error
	CreateInBatches(values []*models.ArticleContent, batchSize int) error
	Save(values ...*models.ArticleContent) error
	First() (*models.ArticleContent, error)
	Take() (*models.ArticleContent, error)
	Last() (*models.ArticleContent, error)
	Find() ([]*models.ArticleContent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleContent, err error)
	FindInBatches(result *[]*models.ArticleContent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ArticleContent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleContentDo
	Assign(attrs ...field.AssignExpr) IArticleContentDo
	Joins(fields ...field.RelationField) IArticleContentDo
	Preload(fields ...field.RelationField) IArticleContentDo
	FirstOrInit() (*models.ArticleContent, error)
	FirstOrCreate() (*models.ArticleContent, error)
	FindByPage(offset int, limit int) (result []*models.ArticleContent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleContentDo) Debug() IArticleContentDo {
	return a.withDO(a.DO.Debug())
}

func (a articleContentDo) WithContext(ctx context.Context) IArticleContentDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleContentDo) ReadDB() IArticleContentDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleContentDo) WriteDB() IArticleContentDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleContentDo) Session(config *gorm.Session) IArticleContentDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleContentDo) Clauses(conds ...clause.Expression) IArticleContentDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleContentDo) Returning(value interface{}, columns ...string) IArticleContentDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleContentDo) Not(conds ...gen.Condition) IArticleContentDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleContentDo) Or(conds ...gen.Condition) IArticleContentDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleContentDo) Select(conds ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleContentDo) Where(conds ...gen.Condition) IArticleContentDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleContentDo) Order(conds ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleContentDo) Distinct(cols ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleContentDo) Omit(cols ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleContentDo) Join(table schema.Tabler, on ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleContentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleContentDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleContentDo) Group(cols ...field.Expr) IArticleContentDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleContentDo) Having(conds ...gen.Condition) IArticleContentDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleContentDo) Limit(limit int) IArticleContentDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleContentDo) Offset(offset int) IArticleContentDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleContentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleContentDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleContentDo) Unscoped() IArticleContentDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleContentDo) Create(values ...*models.ArticleContent) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleContentDo) CreateInBatches(values []*models.ArticleContent, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleContentDo) Save(values ...*models.ArticleContent) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleContentDo) First() (*models.ArticleContent, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleContent), nil
	}
}

func (a articleContentDo) Take() (*models.ArticleContent, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleContent), nil
	}
}

func (a articleContentDo) Last() (*models.ArticleContent, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleContent), nil
	}
}

func (a articleContentDo) Find() ([]*models.ArticleContent, error) {
	result, err := a.DO.Find()
	return result.([]*models.ArticleContent), err
}

func (a articleContentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticleContent, err error) {
	buf := make([]*models.ArticleContent, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleContentDo) FindInBatches(result *[]*models.ArticleContent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleContentDo) Attrs(attrs ...field.AssignExpr) IArticleContentDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleContentDo) Assign(attrs ...field.AssignExpr) IArticleContentDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleContentDo) Joins(fields ...field.RelationField) IArticleContentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleContentDo) Preload(fields ...field.RelationField) IArticleContentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleContentDo) FirstOrInit() (*models.ArticleContent, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleContent), nil
	}
}

func (a articleContentDo) FirstOrCreate() (*models.ArticleContent, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticleContent), nil
	}
}

func (a articleContentDo) FindByPage(offset int, limit int) (result []*models.ArticleContent, count int64, err error) {
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

func (a articleContentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleContentDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleContentDo) Delete(models ...*models.ArticleContent) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleContentDo) withDO(do gen.Dao) *articleContentDo {
	a.DO = *do.(*gen.DO)
	return a
}
