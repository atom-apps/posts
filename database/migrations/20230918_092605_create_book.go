package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_092605CreateBook) table() interface{} {
	type Book struct {
		Model
	}

	return &Book{}
}

func (m *Migration20230918_092605CreateBook) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_092605CreateBook) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230918_092605CreateBookMigration)
}

type Migration20230918_092605CreateBook struct {
	id string
}

func New20230918_092605CreateBookMigration() contracts.Migration {
	return &Migration20230918_092605CreateBook{id: "20230918_092605_create_book"}
}

func (m *Migration20230918_092605CreateBook) ID() string {
	return m.id
}
