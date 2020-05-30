package gopg

import (
	"github.com/go-pg/pg/v10"
)

func NewDB(dataSourceName string) (*pg.DB, error) {
	options, err := pg.ParseURL(dataSourceName)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(options)
	return db, nil
}
