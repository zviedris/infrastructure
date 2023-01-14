package Infrastructure

import (
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type MasterDbInstance struct {
	dbx *sqlx.DB
}

func (m *MasterDbInstance) DBX() *sqlx.DB {
	return m.dbx
}

func NewMasterDbInstance(databaseURL string) *MasterDbInstance {
	return &MasterDbInstance{
		dbx: NewSqlXInstance(databaseURL),
	}
}

// NewDatabase will create new database instance
func NewSqlXInstance(databaseURL string) *sqlx.DB {
	dbx, err := sqlx.Open("pgx", databaseURL)
	if err != nil {
		log.Panic(err)
	}

	if err := dbx.Ping(); err != nil {
		log.Fatal(err)
	}

	dbx.SetConnMaxLifetime(time.Minute * 5)
	dbx.SetMaxIdleConns(0)
	dbx.SetMaxOpenConns(5)
	return dbx
}
