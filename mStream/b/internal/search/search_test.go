package search

import "testing"

// ------------------------------------------------------------------------- //

type caseTestFetchRes struct {
	q string
}

type caseTestParseRes struct {
	q string
	r string
}

// ------------------------------------------------------------------------- //

var casesTestFetchRes = []caseTestFetchRes{
	{q: "NWA - Fuk Da Police"},
	{q: "G-Eazy - The Beauti"},
	{q: "Eminem - Lose Yours"},
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
}

// ------------------------------------------------------------------------- //
