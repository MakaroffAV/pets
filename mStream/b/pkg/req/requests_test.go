package req

import "testing"

// ------------------------------------------------------------------------- //

type caseTestGet struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the test case for "GET" function
	*/

	exp string
	arg caseTestGetArg
}

type caseTestGetArg struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the test case for ....
	*/

	u       string
	qParams map[string]string
	hParams map[string]string
}

// ------------------------------------------------------------------------- //

var (
	casesTestGetTemp = []caseTestGet{
		{
			exp: "sdfs",
			arg: caseTestGetArg{
				u: "https://ya.ru",
				qParams: map[string]string{
					"alex": "alex",
				},
				hParams: map[string]string{
					"alex": "alex",
				},
			},
		},
	}
)

// ------------------------------------------------------------------------- //

func TestGet(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Do HTTP/GET request"
	*/

	for _, c := range casesTestGetTemp {

		resp, respErr := Get(c.arg.u, c.arg.qParams, c.arg.hParams)
		if respErr != nil {
			t.Fatalf(
				`
					Test failed:	Test "Do HTTP/GET request"
									(occurred unhandled error, debug it) 	
				`,
			)
		}

		if string(resp) != c.exp {
			t.Fatalf(
				`
					Test failed:	Test "Do HTTP/GET request"
									(expected and returned values do not match)
				`,
			)
		}

	}
}

// ------------------------------------------------------------------------- //
