package migrations

import (
	"database/sql"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

//Migrate run migrations
func Migrate(uri string, migrationDir string, migrationDirection int) {
	db, err := sql.Open("postgres", uri)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}

	_, err = migrate.Exec(db, "postgres", migrations, migrate.MigrationDirection(migrationDirection))
	if err != nil {
		panic(err)
	}
}
