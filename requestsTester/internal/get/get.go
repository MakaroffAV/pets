package get

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

	reqConfig := req.Request{ReqB: r}

	reqConfig.Parse()

	j, err := json.Marshal(reqConfig)
	if err != nil {
		fmt.Println(err)
	}

	w.Write(j)

}

// ------------------------------------------------------------------------- //
