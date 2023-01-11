package tnreq

import (
	"testing"
)

func TestGet(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Making HTTP/GET request via tor network"
	*/

	_, err := Get("https://ya.ru")
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Making HTTP/GET request via tor network"
								(fetching HTTP/GET tor request results failed)
			`,
		)
	}

}
