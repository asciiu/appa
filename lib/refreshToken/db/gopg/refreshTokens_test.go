package gopg

import (
	"fmt"
	"log"
	"testing"

	"github.com/asciiu/appa/lib/db/gopg"
	"github.com/asciiu/appa/lib/db/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestInsertUser(t *testing.T) {
	testDB := "postgres://postgres@localhost/appa_test?&sslmode=disable"

	sqldb, err := sql.NewDB(testDB)
	checkError(err)
	defer sqldb.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqldb),     // You database connection
		testfixtures.Dialect("postgres"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Files(
			"../../fixtures/users.yaml",
			"../../fixtures/refresh_tokens.yaml",
		),
	)
	checkError(err)

	err = fixtures.Load()
	checkError(err)

	db, err := gopg.NewDB(testDB)
	checkError(err)

	defer db.Close()

	t.Run("Find token", func(t *testing.T) {
		token, err := FindRefreshToken(db, "QuiQqp+CPmeKuhR7Sg6diQ==")
		assert.Nil(t, err, "find should be nil")
		fmt.Println(token)
	})
}
