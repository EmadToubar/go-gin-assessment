package datab

import "github.com/jmoiron/sqlx"

func DbCtxProvider() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
