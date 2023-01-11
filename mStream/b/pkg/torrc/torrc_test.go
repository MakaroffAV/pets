package torrc

import (
	"fmt"
	"testing"
)

func TestCtrl(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Get tor node controller"
	*/

	_, err := Ctrl()
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Get tor node controller"
								(getting tor node controller error)
			`,
		)
	}

}

func TestSendSignal(t *testing.T) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Test "Send signal to tor node controller"
	*/

	c, err := Ctrl()
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Send signal to tor node controller"
								(getting tor node controller error)
			`,
		)
	}

	_, l, err := SendSignal(c, "NEWNYM")
	if err != nil {
		t.Fatalf(
			`
				Test failed:	Test "Send signal to tor node controller"
								(sending )	
			`,
		)
	}

	fmt.Println(l)

}
