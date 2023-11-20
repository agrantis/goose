package gomigrations

import (
	"github.com/agrantis/goose/v3"
)

func init() {
	goose.AddMigrationNoTx(nil, nil)
}
