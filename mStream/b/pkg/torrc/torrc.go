package torrc

import (
	"b/pkg/cslog"
	"net/textproto"
)

func Ctrl() (*textproto.Conn, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Get tor node controller
	*/

	c, err := textproto.Dial("tcp", "127.0.0.1:9050")
	if err != nil {
		cslog.Fail(
			`F001`,
		)
		return nil, err
	}

	return c, nil
}

func SendSignal(c *textproto.Conn, s string) (int, string, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Send signal to tor node controller
	*/

	i, err := c.Cmd("SIGNAL " + s)
	if err != nil {
		cslog.Fail(
			`F001`,
		)
		return 0, "", err
	}

	c.StartResponse(i)
	defer c.EndResponse(i)

	z, v, t := c.ReadResponse(20)

	return z, v, t

}
