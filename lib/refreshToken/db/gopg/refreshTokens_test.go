package gopg

import (
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

	tokenRepo := NewTokenRepo(db)

	t.Run("Find token", func(t *testing.T) {
		token, err := tokenRepo.FindRefreshToken("QuiQqp+CPmeKuhR7Sg6diQ==")
		assert.Nil(t, err, "find should be nil")
		assert.Equal(t, "QuiQqp+CPmeKuhR7Sg6diQ==", token.Selector, "selector incorrect")
		assert.Equal(t, "fci0UvzlZVcFJvvAqonddYK4TSs6JUCT83ptIy6mvcY=", token.TokenHash, "token hash incorrect")
	})
}
