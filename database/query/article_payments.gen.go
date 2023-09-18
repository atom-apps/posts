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

func newArticlePayment(db *gorm.DB, opts ...gen.DOOption) articlePayment {
	_articlePayment := articlePayment{}

	_articlePayment.articlePaymentDo.UseDB(db, opts...)
	_articlePayment.articlePaymentDo.UseModel(&models.ArticlePayment{})

	tableName := _articlePayment.articlePaymentDo.TableName()
	_articlePayment.ALL = field.NewAsterisk(tableName)
	_articlePayment.ID = field.NewUint64(tableName, "id")
	_articlePayment.CreatedAt = field.NewTime(tableName, "created_at")
	_articlePayment.ArticleID = field.NewUint64(tableName, "article_id")
	_articlePayment.PriceType = field.NewInt64(tableName, "price_type")
	_articlePayment.Token = field.NewString(tableName, "token")
	_articlePayment.Price = field.NewUint64(tableName, "price")
	_articlePayment.Discount = field.NewUint64(tableName, "discount")
	_articlePayment.DiscountStartAt = field.NewTime(tableName, "discount_start_at")
	_articlePayment.DiscountEndAt = field.NewTime(tableName, "discount_end_at")

	_articlePayment.fillFieldMap()

	return _articlePayment
}

type articlePayment struct {
	articlePaymentDo articlePaymentDo

	ALL             field.Asterisk
	ID              field.Uint64 // ID
	CreatedAt       field.Time   // 创建时间
	ArticleID       field.Uint64 // 文章ID
	PriceType       field.Int64  // 付费类型
	Token           field.String // 付费密码
	Price           field.Uint64 // 付费价格
	Discount        field.Uint64 // 付费折扣
	DiscountStartAt field.Time   // 折扣开始时间
	DiscountEndAt   field.Time   // 折扣结束时间

	fieldMap map[string]field.Expr
}

func (a articlePayment) Table(newTableName string) *articlePayment {
	a.articlePaymentDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articlePayment) As(alias string) *articlePayment {
	a.articlePaymentDo.DO = *(a.articlePaymentDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articlePayment) updateTableName(table string) *articlePayment {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.ArticleID = field.NewUint64(table, "article_id")
	a.PriceType = field.NewInt64(table, "price_type")
	a.Token = field.NewString(table, "token")
	a.Price = field.NewUint64(table, "price")
	a.Discount = field.NewUint64(table, "discount")
	a.DiscountStartAt = field.NewTime(table, "discount_start_at")
	a.DiscountEndAt = field.NewTime(table, "discount_end_at")

	a.fillFieldMap()

	return a
}

func (a *articlePayment) WithContext(ctx context.Context) IArticlePaymentDo {
	return a.articlePaymentDo.WithContext(ctx)
}

func (a articlePayment) TableName() string { return a.articlePaymentDo.TableName() }

func (a articlePayment) Alias() string { return a.articlePaymentDo.Alias() }

func (a articlePayment) Columns(cols ...field.Expr) gen.Columns {
	return a.articlePaymentDo.Columns(cols...)
}

func (a *articlePayment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articlePayment) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["price_type"] = a.PriceType
	a.fieldMap["token"] = a.Token
	a.fieldMap["price"] = a.Price
	a.fieldMap["discount"] = a.Discount
	a.fieldMap["discount_start_at"] = a.DiscountStartAt
	a.fieldMap["discount_end_at"] = a.DiscountEndAt
}

func (a articlePayment) clone(db *gorm.DB) articlePayment {
	a.articlePaymentDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articlePayment) replaceDB(db *gorm.DB) articlePayment {
	a.articlePaymentDo.ReplaceDB(db)
	return a
}

type articlePaymentDo struct{ gen.DO }

type IArticlePaymentDo interface {
	gen.SubQuery
	Debug() IArticlePaymentDo
	WithContext(ctx context.Context) IArticlePaymentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticlePaymentDo
	WriteDB() IArticlePaymentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticlePaymentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticlePaymentDo
	Not(conds ...gen.Condition) IArticlePaymentDo
	Or(conds ...gen.Condition) IArticlePaymentDo
	Select(conds ...field.Expr) IArticlePaymentDo
	Where(conds ...gen.Condition) IArticlePaymentDo
	Order(conds ...field.Expr) IArticlePaymentDo
	Distinct(cols ...field.Expr) IArticlePaymentDo
	Omit(cols ...field.Expr) IArticlePaymentDo
	Join(table schema.Tabler, on ...field.Expr) IArticlePaymentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticlePaymentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticlePaymentDo
	Group(cols ...field.Expr) IArticlePaymentDo
	Having(conds ...gen.Condition) IArticlePaymentDo
	Limit(limit int) IArticlePaymentDo
	Offset(offset int) IArticlePaymentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticlePaymentDo
	Unscoped() IArticlePaymentDo
	Create(values ...*models.ArticlePayment) error
	CreateInBatches(values []*models.ArticlePayment, batchSize int) error
	Save(values ...*models.ArticlePayment) error
	First() (*models.ArticlePayment, error)
	Take() (*models.ArticlePayment, error)
	Last() (*models.ArticlePayment, error)
	Find() ([]*models.ArticlePayment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticlePayment, err error)
	FindInBatches(result *[]*models.ArticlePayment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ArticlePayment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticlePaymentDo
	Assign(attrs ...field.AssignExpr) IArticlePaymentDo
	Joins(fields ...field.RelationField) IArticlePaymentDo
	Preload(fields ...field.RelationField) IArticlePaymentDo
	FirstOrInit() (*models.ArticlePayment, error)
	FirstOrCreate() (*models.ArticlePayment, error)
	FindByPage(offset int, limit int) (result []*models.ArticlePayment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticlePaymentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articlePaymentDo) Debug() IArticlePaymentDo {
	return a.withDO(a.DO.Debug())
}

func (a articlePaymentDo) WithContext(ctx context.Context) IArticlePaymentDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articlePaymentDo) ReadDB() IArticlePaymentDo {
	return a.Clauses(dbresolver.Read)
}

func (a articlePaymentDo) WriteDB() IArticlePaymentDo {
	return a.Clauses(dbresolver.Write)
}

func (a articlePaymentDo) Session(config *gorm.Session) IArticlePaymentDo {
	return a.withDO(a.DO.Session(config))
}

func (a articlePaymentDo) Clauses(conds ...clause.Expression) IArticlePaymentDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articlePaymentDo) Returning(value interface{}, columns ...string) IArticlePaymentDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articlePaymentDo) Not(conds ...gen.Condition) IArticlePaymentDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articlePaymentDo) Or(conds ...gen.Condition) IArticlePaymentDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articlePaymentDo) Select(conds ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articlePaymentDo) Where(conds ...gen.Condition) IArticlePaymentDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articlePaymentDo) Order(conds ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articlePaymentDo) Distinct(cols ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articlePaymentDo) Omit(cols ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articlePaymentDo) Join(table schema.Tabler, on ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articlePaymentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articlePaymentDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articlePaymentDo) Group(cols ...field.Expr) IArticlePaymentDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articlePaymentDo) Having(conds ...gen.Condition) IArticlePaymentDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articlePaymentDo) Limit(limit int) IArticlePaymentDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articlePaymentDo) Offset(offset int) IArticlePaymentDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articlePaymentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticlePaymentDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articlePaymentDo) Unscoped() IArticlePaymentDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articlePaymentDo) Create(values ...*models.ArticlePayment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articlePaymentDo) CreateInBatches(values []*models.ArticlePayment, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articlePaymentDo) Save(values ...*models.ArticlePayment) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articlePaymentDo) First() (*models.ArticlePayment, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePayment), nil
	}
}

func (a articlePaymentDo) Take() (*models.ArticlePayment, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePayment), nil
	}
}

func (a articlePaymentDo) Last() (*models.ArticlePayment, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePayment), nil
	}
}

func (a articlePaymentDo) Find() ([]*models.ArticlePayment, error) {
	result, err := a.DO.Find()
	return result.([]*models.ArticlePayment), err
}

func (a articlePaymentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticlePayment, err error) {
	buf := make([]*models.ArticlePayment, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articlePaymentDo) FindInBatches(result *[]*models.ArticlePayment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articlePaymentDo) Attrs(attrs ...field.AssignExpr) IArticlePaymentDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articlePaymentDo) Assign(attrs ...field.AssignExpr) IArticlePaymentDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articlePaymentDo) Joins(fields ...field.RelationField) IArticlePaymentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articlePaymentDo) Preload(fields ...field.RelationField) IArticlePaymentDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articlePaymentDo) FirstOrInit() (*models.ArticlePayment, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePayment), nil
	}
}

func (a articlePaymentDo) FirstOrCreate() (*models.ArticlePayment, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePayment), nil
	}
}

func (a articlePaymentDo) FindByPage(offset int, limit int) (result []*models.ArticlePayment, count int64, err error) {
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

func (a articlePaymentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articlePaymentDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articlePaymentDo) Delete(models ...*models.ArticlePayment) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articlePaymentDo) withDO(do gen.Dao) *articlePaymentDo {
	a.DO = *do.(*gen.DO)
	return a
}
