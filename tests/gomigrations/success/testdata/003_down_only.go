package gomigrations

import (
	"database/sql"

	"github.com/agrantis/goose/v3"
)

func init() {
	goose.AddMigration(nil, down003)
}

func down003(tx *sql.Tx) error {
	return dropTable(tx, "bravo")
}
