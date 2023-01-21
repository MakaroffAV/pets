package torrc

import "net/textproto"

func sendSignal(c *textproto.Conn, s string) (int, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Send signal to tor controller
	*/

	i, err := c.Cmd(s)
	if err != nil {
		return 0, err
	}

	c.StartResponse(i)
	defer c.EndResponse(i)

	return 1, nil

}

func GetCtrl() (*textproto.Conn, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Create the new tor controller
	*/

	c, err := textproto.Dial("tcp", "127.0.0.1:9050")
	if err != nil {
		return nil, err
	}

	return c, nil

}
