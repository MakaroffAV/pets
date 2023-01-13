package db

import (
	"testing"
)

func TestGetConn(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Create connection with database"
	*/

	_, err := GetCon()
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Create connection with database"
			`,
		)
	}

}

func TestDelCon(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Delete connection with database"
	*/
}
