package req

import "testing"

func TestSetUpClient(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Setting up request client"
	*/

	_, err := setUpClient()
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Setting up request client"
								(error while setting up the request client)
			`,
		)
	}

}
