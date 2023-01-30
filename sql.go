package Infrastructure

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type MasterDbInstance struct {
	dbx *sqlx.DB
}

func (m *MasterDbInstance) DBX() *sqlx.DB {
	return m.dbx
}

func NewMasterDbInstance(driver string, databaseURL string) *MasterDbInstance {
	return &MasterDbInstance{
		dbx: NewSqlXInstance(driver, databaseURL, 10),
	}
}

// NewDatabase will create new database instance
func NewSqlXInstance(driver string, databaseURL string, maxConnections int) *sqlx.DB {
	dbx, err := sqlx.Open(driver, databaseURL)
	if err != nil {
		log.Panic(err)
	}

	if err := dbx.Ping(); err != nil {
		log.Fatal(err)
	}

	dbx.SetConnMaxLifetime(time.Minute * 5)
	dbx.SetMaxIdleConns(0)
	dbx.SetMaxOpenConns(maxConnections)
	return dbx
}
