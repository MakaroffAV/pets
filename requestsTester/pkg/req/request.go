package req

import (
	"net/http"
)

// ------------------------------------------------------------------------- //

type Request struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the config of HTTP(s) request
	*/

	ReqB *http.Request

	host string
	path string

	qParams map[string][]string
	hParams map[string][]string
}

// ------------------------------------------------------------------------- //

func (req *Request) parseUParams() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing url params
	*/

	req.host = req.ReqB.URL.Host
	req.path = req.ReqB.URL.Path

}

func (req *Request) parseHParams() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing header params
	*/

	req.hParams = req.ReqB.Header

}

func (req *Request) parseQParams() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing query params
	*/

	req.qParams = req.ReqB.URL.Query()

}

// ------------------------------------------------------------------------- //

func (req Request) Parse() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing HTTP request
	*/

	req.parseHParams()
	req.parseQParams()
	req.parseUParams()

}

// ------------------------------------------------------------------------- //
