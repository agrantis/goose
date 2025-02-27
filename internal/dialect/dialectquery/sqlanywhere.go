package dialectquery

import "fmt"

type SQLAnywhere struct{}

func (s *SQLAnywhere) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		id INTEGER DEFAULT AUTOINCREMENT,
		version_id bigint NOT NULL,
		is_applied bit NOT NULL,
		tstamp timestamp NULL default now(),
		PRIMARY KEY(id)
	)`
	return fmt.Sprintf(q, tableName)
}

func (s *SQLAnywhere) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (version_id, is_applied) VALUES (?, ?)`
	return fmt.Sprintf(q, tableName)
}

func (s *SQLAnywhere) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, tableName)
}

func (s *SQLAnywhere) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

func (s *SQLAnywhere) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied from %s ORDER BY id DESC`
	return fmt.Sprintf(q, tableName)
}

var _ Querier = (*SQLAnywhere)(nil)
