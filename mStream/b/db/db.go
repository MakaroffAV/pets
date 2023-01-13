package db

import (
	"b/cnf"
	"b/pkg/cslog"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DelCon(c *sql.DB) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Delete connection with database
	*/

	c.Close()

}

func GetCon() (*sql.DB, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Create connection with database
	*/

	c, err := sql.Open(`sqlite3`, cnf.DbPath)
	if err != nil {
		cslog.Fail(
			`F0000001`,
		)
		return nil, err
	}

	return c, nil

}
