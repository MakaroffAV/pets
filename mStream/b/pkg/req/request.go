package req

import (
	"b/pkg/cslog"
	"io/ioutil"
	"net/http"
)

// ------------------------------------------------------------------------- //

type request struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the HTTP request`s params
	*/

	href *string
	meth *string

	hParams *map[string]string
	qParams *map[string]string

	enq    *http.Request
	enqErr error

	res    *http.Response
	resErr error

	cnt    []byte
	cntErr error
}

// ------------------------------------------------------------------------- //

func (r *request) setUpReq() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Setting up the request config
	*/

	u := Url{
		Href: *r.Href,
		Qprm: *r.ParamsQ,
	}

	q, err := http.NewRequest(*r.Meth, u.Build(), nil)
	if err != nil {
		return
	}

	for k, v := range *r.ParamsH {
		q.Header.Set(k, v)
	}

}

func (r *request) fetchReq() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Fetching HTTP request
	*/

	c, err := setUpClient()
	if err != nil {
		return
	}

	r.res, r.resErr = c.Do(r.enq)
	if r.resErr != nil {
		return
	}

}

func (r *request) parseRes() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing HTTP request
	*/

	r.cnt, r.cntErr = ioutil.ReadAll(r.res.Body)
	if r.cntErr != nil {
		cslog.Fail("r0", r.cntErr)
		return
	}

}

// ------------------------------------------------------------------------- //

func (r request) Do() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Execute the HTTP request
	*/

	r.setUpReq()
	r.fetchReq()
	r.fetchReq()

}

// ------------------------------------------------------------------------- //
