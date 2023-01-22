package req

import (
	"net/http"
)

// ------------------------------------------------------------------------- //

type request struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the config of HTTP(s) request
	*/

	Host string
	Path string
	Meth string

	DParams map[string][]string
	QParams map[string][]string
	HParams map[string][]string
}

// ------------------------------------------------------------------------- //

func setUParams(req *request, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing url params
				of the HTTP(s) request
	*/

	req.Host = r.Host
	req.Meth = r.Method
	req.Path = r.URL.Path

}

func setHParams(req *request, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing header params
				of the HTTP(s) request
	*/

	req.HParams = r.Header

}

func setQParams(req *request, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing query params
				of the HTTP(s) request
	*/

	req.QParams = r.URL.Query()

}

func setDParams(req *request, r *http.Request) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing data body
				of the HTTP(s) request
	*/

	err := r.ParseMultipartForm(0)
	if err == nil {
		req.DParams = r.PostForm
	}

}

// ------------------------------------------------------------------------- //

func Parse(r *http.Request) (*request, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing HTTP(S) request
	*/

	req := new(request)

	setUParams(req, r)
	setHParams(req, r)
	setQParams(req, r)
	setDParams(req, r)

	return req, nil

}

// ------------------------------------------------------------------------- //
