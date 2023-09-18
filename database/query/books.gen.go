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

func newBook(db *gorm.DB, opts ...gen.DOOption) book {
	_book := book{}

	_book.bookDo.UseDB(db, opts...)
	_book.bookDo.UseModel(&models.Book{})

	tableName := _book.bookDo.TableName()
	_book.ALL = field.NewAsterisk(tableName)
	_book.ID = field.NewUint64(tableName, "id")
	_book.CreatedAt = field.NewTime(tableName, "created_at")
	_book.UpdatedAt = field.NewTime(tableName, "updated_at")
	_book.DeletedAt = field.NewField(tableName, "deleted_at")
	_book.TenantID = field.NewUint64(tableName, "tenant_id")
	_book.UserID = field.NewUint64(tableName, "user_id")
	_book.Title = field.NewString(tableName, "title")
	_book.ThumbnailFileID = field.NewUint64(tableName, "thumbnail_file_id")
	_book.Description = field.NewString(tableName, "description")
	_book.Content = field.NewString(tableName, "content")
	_book.Author = field.NewString(tableName, "author")
	_book.Source = field.NewString(tableName, "source")
	_book.Isbn = field.NewString(tableName, "isbn")
	_book.Price = field.NewUint64(tableName, "price")

	_book.fillFieldMap()

	return _book
}

type book struct {
	bookDo bookDo

	ALL             field.Asterisk
	ID              field.Uint64 // ID
	CreatedAt       field.Time   // 创建时间
	UpdatedAt       field.Time   // 更新时间
	DeletedAt       field.Field  // 删除时间
	TenantID        field.Uint64 // 租户ID
	UserID          field.Uint64 // 用户ID
	Title           field.String // 书名
	ThumbnailFileID field.Uint64 // 封面
	Description     field.String // 简介
	Content         field.String // 详细介绍
	Author          field.String // 原作者
	Source          field.String // 原书地址
	Isbn            field.String // ISBN
	Price           field.Uint64 // 价格

	fieldMap map[string]field.Expr
}

func (b book) Table(newTableName string) *book {
	b.bookDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b book) As(alias string) *book {
	b.bookDo.DO = *(b.bookDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *book) updateTableName(table string) *book {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewUint64(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.TenantID = field.NewUint64(table, "tenant_id")
	b.UserID = field.NewUint64(table, "user_id")
	b.Title = field.NewString(table, "title")
	b.ThumbnailFileID = field.NewUint64(table, "thumbnail_file_id")
	b.Description = field.NewString(table, "description")
	b.Content = field.NewString(table, "content")
	b.Author = field.NewString(table, "author")
	b.Source = field.NewString(table, "source")
	b.Isbn = field.NewString(table, "isbn")
	b.Price = field.NewUint64(table, "price")

	b.fillFieldMap()

	return b
}

func (b *book) WithContext(ctx context.Context) IBookDo { return b.bookDo.WithContext(ctx) }

func (b book) TableName() string { return b.bookDo.TableName() }

func (b book) Alias() string { return b.bookDo.Alias() }

func (b book) Columns(cols ...field.Expr) gen.Columns { return b.bookDo.Columns(cols...) }

func (b *book) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *book) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 14)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["tenant_id"] = b.TenantID
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["title"] = b.Title
	b.fieldMap["thumbnail_file_id"] = b.ThumbnailFileID
	b.fieldMap["description"] = b.Description
	b.fieldMap["content"] = b.Content
	b.fieldMap["author"] = b.Author
	b.fieldMap["source"] = b.Source
	b.fieldMap["isbn"] = b.Isbn
	b.fieldMap["price"] = b.Price
}

func (b book) clone(db *gorm.DB) book {
	b.bookDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b book) replaceDB(db *gorm.DB) book {
	b.bookDo.ReplaceDB(db)
	return b
}

type bookDo struct{ gen.DO }

type IBookDo interface {
	gen.SubQuery
	Debug() IBookDo
	WithContext(ctx context.Context) IBookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBookDo
	WriteDB() IBookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookDo
	Not(conds ...gen.Condition) IBookDo
	Or(conds ...gen.Condition) IBookDo
	Select(conds ...field.Expr) IBookDo
	Where(conds ...gen.Condition) IBookDo
	Order(conds ...field.Expr) IBookDo
	Distinct(cols ...field.Expr) IBookDo
	Omit(cols ...field.Expr) IBookDo
	Join(table schema.Tabler, on ...field.Expr) IBookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookDo
	Group(cols ...field.Expr) IBookDo
	Having(conds ...gen.Condition) IBookDo
	Limit(limit int) IBookDo
	Offset(offset int) IBookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookDo
	Unscoped() IBookDo
	Create(values ...*models.Book) error
	CreateInBatches(values []*models.Book, batchSize int) error
	Save(values ...*models.Book) error
	First() (*models.Book, error)
	Take() (*models.Book, error)
	Last() (*models.Book, error)
	Find() ([]*models.Book, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Book, err error)
	FindInBatches(result *[]*models.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Book) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookDo
	Assign(attrs ...field.AssignExpr) IBookDo
	Joins(fields ...field.RelationField) IBookDo
	Preload(fields ...field.RelationField) IBookDo
	FirstOrInit() (*models.Book, error)
	FirstOrCreate() (*models.Book, error)
	FindByPage(offset int, limit int) (result []*models.Book, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookDo) Debug() IBookDo {
	return b.withDO(b.DO.Debug())
}

func (b bookDo) WithContext(ctx context.Context) IBookDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookDo) ReadDB() IBookDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookDo) WriteDB() IBookDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookDo) Session(config *gorm.Session) IBookDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookDo) Clauses(conds ...clause.Expression) IBookDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookDo) Returning(value interface{}, columns ...string) IBookDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookDo) Not(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookDo) Or(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookDo) Select(conds ...field.Expr) IBookDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookDo) Where(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookDo) Order(conds ...field.Expr) IBookDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookDo) Distinct(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookDo) Omit(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookDo) Join(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookDo) Group(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookDo) Having(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookDo) Limit(limit int) IBookDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookDo) Offset(offset int) IBookDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookDo) Unscoped() IBookDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookDo) Create(values ...*models.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookDo) CreateInBatches(values []*models.Book, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookDo) Save(values ...*models.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookDo) First() (*models.Book, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Book), nil
	}
}

func (b bookDo) Take() (*models.Book, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Book), nil
	}
}

func (b bookDo) Last() (*models.Book, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Book), nil
	}
}

func (b bookDo) Find() ([]*models.Book, error) {
	result, err := b.DO.Find()
	return result.([]*models.Book), err
}

func (b bookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Book, err error) {
	buf := make([]*models.Book, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookDo) FindInBatches(result *[]*models.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookDo) Attrs(attrs ...field.AssignExpr) IBookDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookDo) Assign(attrs ...field.AssignExpr) IBookDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookDo) Joins(fields ...field.RelationField) IBookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookDo) Preload(fields ...field.RelationField) IBookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookDo) FirstOrInit() (*models.Book, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Book), nil
	}
}

func (b bookDo) FirstOrCreate() (*models.Book, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Book), nil
	}
}

func (b bookDo) FindByPage(offset int, limit int) (result []*models.Book, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookDo) Delete(models ...*models.Book) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookDo) withDO(do gen.Dao) *bookDo {
	b.DO = *do.(*gen.DO)
	return b
}
