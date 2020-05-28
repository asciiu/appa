package gopg_test

import (
	"log"
	"testing"

	adb "github.com/asciiu/appa/lib/db"
	"github.com/asciiu/appa/lib/user/db/gopg"
	"github.com/go-pg/pg/v10"
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

	sqldb, err := adb.NewDB(testDB)
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

	options, err := pg.ParseURL(testDB)
	checkError(err)

	db := pg.Connect(options)
	defer db.Close()

	t.Run("Find User by email", func(t *testing.T) {
		//newUser := users.NewUser("jester", "admin@flo", "password")
		//err = db.Insert(newUser)
		//if err != nil {
		//	panic(err)
		//}

		foundUser, err := gopg.FindUserByEmail(db, "test@email")
		assert.Nil(t, err, "this should be nil")

		assert.Equal(t, "test@email", foundUser.Email, "email incorrect")
		assert.Equal(t, "tester", foundUser.Username, "email incorrect")
	})
}
