package search

import (
	"b/cnf"
	"b/pkg/cslog"
	"b/pkg/tnreq"
	"fmt"
)

// ------------------------------------------------------------------------- //

func parseRes(r []byte) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing search results
	*/
}

func fetchRes(q *string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Fetching search results
	*/

	r, err := tnreq.Get(fmt.Sprintf("%s/results?search_query%s", cnf.Host, *q))
	if err != nil {
		cslog.Fail(
			"F00000001",
			err.Error(),
		)
		return nil, err
	}

	return r, nil

}

// ------------------------------------------------------------------------- //

func Do(q string, t string) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:
	*/
}

// ------------------------------------------------------------------------- //
