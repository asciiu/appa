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
		),
	)
	checkError(err)

	err = fixtures.Load()
	checkError(err)

	db, err := gopg.NewDB(testDB)
	checkError(err)

	defer db.Close()

	userRepo := NewUserRepo(db)

	t.Run("Find User by email", func(t *testing.T) {
		foundUser, err := userRepo.FindUserByEmail("test@email")
		assert.Nil(t, err, "this should be nil")

		assert.Equal(t, "test@email", foundUser.Email, "email incorrect")
		assert.Equal(t, "tester", foundUser.Username, "email incorrect")
	})
}
