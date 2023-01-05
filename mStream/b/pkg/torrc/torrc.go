package torrc

import (
	"b/pkg/cslog"
	"net/textproto"
)

func Signal(c *textproto.Conn, s string) (int, string, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Sending signal to tor node
	*/

	r, err := c.Cmd(s)
	if err != nil {
		cslog.Fail(
			"F001",
		)
		return 0, "", err
	}

	c.StartResponse(r)

	defer c.EndResponse(r)

	return c.ReadResponse(250)

}
