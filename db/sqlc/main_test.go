package arfcasesqlc

import (
	"database/sql"
	"log"
	"os"
	"serhatbxld/arf-case/util"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbSource := `postgresql://` + config.DBUser + `:` + config.DBPassword + `@postgres:` + config.DBPort + `/` + config.DBName + `?sslmode=disable`

	testDB, err = sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
