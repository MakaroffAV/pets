package search

import (
	"testing"
)

// ------------------------------------------------------------------------- //

type caseTestFetchRes struct {
	q string
}

type caseTestParseRes struct {
	q string
	e SearchRes
}

// ------------------------------------------------------------------------- //

var casesTestFetchRes = []caseTestFetchRes{
	{
		q: "NWA - Fuk Da Police",
	},
	{
		q: "G-Eazy - The Beauti",
	},
	{
		q: "Eminem - Lose Yours",
	},
}

var casesTestParseRes = []caseTestParseRes{
	{
		q: "G-Eazy - The Beauti",
		e: SearchRes{
			Id: "6PUf1Rj4kpc",
			Ch: "G-Eazy",
			Nm: "G-Eazy - The Beautiful \\u0026 Damned (Audio) ft. Zoe Nash",
			Dr: "3:11",
			Vw: "34,303,705",
		},
	},
}

// ------------------------------------------------------------------------- //

func TestFetchRes(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Fetching search results"
	*/

	for _, c := range casesTestFetchRes {
		_, err := fetchRes(&c.q)
		if err != nil {
			t.Fatalf(
				`
					Test failed:	Test "Fetching search results"
									(fetching search results error)
				`,
			)
		}

	}

}

func TestParseRes(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Parsing search results"
	*/

	for _, c := range casesTestParseRes {

		b, err := fetchRes(&c.q)
		if err != nil {
			t.Fatalf(
				`
					Test failed: 	Test "Parsing search results"
									(fetching search results error)
				`,
			)
		}

		r, err := parseRes(b)
		if err != nil {
			t.Fatalf(
				`
					Test failed:	Test "Parsing search results"
									(parsing search results error) 	
				`,
			)
		}

		if r[0].Id != c.e.Id {
			t.Fatalf(
				`
					Test failed:	Test "Parsing search results"
									(expected and actual values match error)
				`,
			)
		}

	}

}

// ------------------------------------------------------------------------- //
