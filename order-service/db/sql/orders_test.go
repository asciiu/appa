package sql_test

import (
	"log"
	"testing"

	"github.com/asciiu/oldiez/common/db"
	"github.com/asciiu/oldiez/user-service/db/sql"
	"github.com/asciiu/oldiez/user-service/models"
	//"github.com/asciiu/oldiez/order-service/db/sql"
	//"github.com/asciiu/oldiez/order-service/models"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func TestInsertOrder(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/oldiez_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("first", "last", "test@email", "hash")
	_, error := sql.InsertUser(db, user)
	if error != nil {
		t.Errorf("%s", error)
	}
	//fmt.Printf("%#v", *user)
}
