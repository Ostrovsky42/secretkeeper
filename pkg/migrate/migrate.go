package migrate

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // pg driver
	"github.com/pressly/goose/v3"
)

const DriverPostgres = "postgres"

type Migration interface {
	Execute() error
}

type migration struct {
	migrationPath string
	dsn           string
}

func New(migrationPath string, dsn string) Migration {
	return &migration{
		migrationPath: migrationPath,
		dsn:           dsn,
	}
}

func (m *migration) Execute() error {
	command, filename := parseFlag()
	if len(command) == 0 {
		return fmt.Errorf("use -c flag with commands: %s", helpCommand)
	}

	db, err := sql.Open(DriverPostgres, m.dsn)
	if err != nil {
		return err
	}

	switch command {
	case "up":
		err = goose.Up(db, m.migrationPath)
	case "upbyone":
		err = goose.UpByOne(db, m.migrationPath)
	case "down":
		err = goose.Down(db, m.migrationPath)
	case "reset":
		err = goose.Reset(db, m.migrationPath)
	case "status":
		err = goose.Status(db, m.migrationPath)
	case "create":
		err = goose.Create(db, m.migrationPath, filename, "sql")
	default:
		err = fmt.Errorf("expected commands: %s", helpCommand)
	}

	return err
}

const helpCommand = `
	up        Migrate the DB to the most recent version available
	upbyone   Migrate the DB up by 1
	down      Rollback the last database migration
	reset     Roll back all migrations
	create    Create migration with -n flag
	status    Dump the migration status for the current DB
`
