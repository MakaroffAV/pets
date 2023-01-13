package search

import (
	"b/cnf"
	"b/pkg/cslog"
	"b/pkg/tnreq"
)

// ------------------------------------------------------------------------- //

type SearchRes struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for describing
				the found element from the search
	*/

	Id string
	Ch string
	Nm string
	Dr string
	Vw string
}

// ------------------------------------------------------------------------- //

func parseResponse(r []byte) (sr []*SearchRes, err error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing search results
	*/

	for _, m := range cnf.RegSome.FindAllStringSubmatch(string(r), -1) {
		sr = append(
			sr,
			&SearchRes{
				Id: m[1],
				Ch: m[3],
				Nm: m[2],
				Dr: m[4],
				Vw: m[5],
			},
		)
	}

	return sr, nil

}

func fetchResponse(p *string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Fetching search results
	*/

	g := map[string]string{
		"search_query": *p,
	}

	r, err := tnreq.Get(cnf.SHost, g)
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
