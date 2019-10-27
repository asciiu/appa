// Fluid elixir API
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/asciiu/appa/lib/db"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		panic("true/false command line argument required for http/https")
	}

	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	fluidDB, err := db.NewDB(dbURL)
	if err != nil {
		log.Printf("ERROR: %s", err)
	} else {
		defer fluidDB.Close()
		router := NewRouter(fluidDB)

		if isHTTPS, _ := strconv.ParseBool(os.Args[1]); isHTTPS {
			// HTTPs server
			router.Logger.Fatal(router.StartAutoTLS(":443"))
		} else {
			// HTTP server
			port := fmt.Sprintf("%s", os.Getenv("API_PORT"))
			router.Logger.Fatal(router.Start(":" + port))
		}
	}
}
