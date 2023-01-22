package h

import (
	"encoding/json"
	"fmt"
	"net/http"
	"requestsTester/pkg/req"
)

// ------------------------------------------------------------------------- //

func Handler(w http.ResponseWriter, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Handler for GET request
	*/

	reqCnf, reqCnfErr := req.Parse(r)
	if reqCnfErr != nil {
		fmt.Println(reqCnfErr)
	}

	reqSer, reqSerErr := json.Marshal(reqCnf)
	if reqSerErr != nil {
		fmt.Println(reqSerErr)
	}

	w.Write(reqSer)

}

// ------------------------------------------------------------------------- //
