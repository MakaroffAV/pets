package auth

import (
	"b/db"
	"fmt"
)

// ------------------------------------------------------------------------- //

func CheckToken(t *string) (bool, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Check if user token is valid
	*/

	c, err := db.GetCon()
	if err != nil {
		return false, err
	}

	r, err := c.Query(
		"SELECT * FROM `users_tokens` WHERE `token` = ?",
		*t,
	)
	if err != nil {
		return false, err
	}

	defer r.Close()

	fmt.Println(r)

	return true, nil

}

// ------------------------------------------------------------------------- //
