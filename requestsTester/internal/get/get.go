package get

import (
	"fmt"
	"net/http"
	"requestsTester/pkg/request"
)

// ------------------------------------------------------------------------- //

func Handler(w http.ResponseWriter, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Handler for GET request
	*/

	fmt.Println("GET params were:", r.URL.Query())

	req, err := request.Parse(r)
	if err != nil {

	}

	w.Write(req)

}

// ------------------------------------------------------------------------- //
