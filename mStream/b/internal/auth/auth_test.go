package auth

import "testing"

// ------------------------------------------------------------------------- //

type caseTestCheckToken struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
	*/

	a string
	e bool
}

// ------------------------------------------------------------------------- //

var casesTestCheckToken = []caseTestCheckToken{
	{
		a: "a",
		e: false,
	},
}

func TestCheckToken(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Check if user token is valid"
	*/

	for _, c := range casesTestCheckToken {

		r, err := CheckToken(&c.a)
		if err != nil {
			t.Fatalf(err.Error())
			t.Fatalf(
				`
					Test failed:	Test "Check if user token is valid"
									(in test occurred undefined error)	
				`,
			)
		}

		if r != c.e {
			t.Fatalf(
				`
					Test failed:	Test "Check if user token is valid"
									(returned and expected values do not match)
				`,
			)
		}

	}

}

// ------------------------------------------------------------------------- //
