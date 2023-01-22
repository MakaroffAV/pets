package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// ------------------------------------------------------------------------- //

type request struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the config of HTTP(s) request
	*/

	host string
	path string

	qParams []qParams
	hParams []hParams
}

type qParams struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the
	*/

	key string
	val []string
}

type hParams struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the
	*/
}

// ------------------------------------------------------------------------- //

func (req *request) parseUParams(r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing url params
	*/

	req.host = r.URL.Host
	req.path = r.URL.Path

}
func (r *request) parseQParams(uv url.Values) error {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing query params
	*/

	for key, v := range uv {
		qArgs := make([]string, len(v))
		for _, e := range v {
			qArgs = append(qArgs, e)
		}
		r.qParams = append(r.qParams, qParams{key: key, val: qArgs})
	}

}

func (r *request) parseHParams() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing header params
	*/

}

// ------------------------------------------------------------------------- //

func Parse(r *http.Request) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing HTTP request
	*/

	req := new(request)

	req.parseUParams(r)
	req.parseQParams(r)
	req.parseHParams(r)

	return json.Marshal(req), errors.New("sdfs")

}

// ------------------------------------------------------------------------- //
