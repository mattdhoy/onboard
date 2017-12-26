package crud

import (
	"database/sql"
	"fmt"
)

// Connect returns a database connection
func Connect() *sql.DB {
	c := GetConf()
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.DbUser, c.DbPass, c.DbName)
	db, err := sql.Open("postgres", dbInfo)
	CheckErr(err)
	return db
}

// CheckErr panics if an error exists
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
