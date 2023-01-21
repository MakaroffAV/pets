package torrc

import (
	"testing"
)

func TestGetCtrl(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Create the new tor controller"
	*/

	_, err := GetCtrl()
	if err != nil {
		t.Fatalf(
			`
				Test failed:	"Creating the new tor controller"
								(occurred error %s)
			`,
			err.Error(),
		)
	}
}

func TestSendSignal(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Send signal to tor controller"
	*/

	c, err := GetCtrl()
	if err != nil {
		t.Fatal()
	}

	_, er := sendSignal(c, "SIGNAL NEWNYM")
	if er != nil {
		t.Fatal()
	}
}
