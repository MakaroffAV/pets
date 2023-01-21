package req

import (
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

	req    *http.Request
	reqErr error

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

	u := rUrl{
		href:    *r.href,
		qParams: *r.qParams,
	}

	r.req, r.reqErr = http.NewRequest(*r.meth, u.Build(), nil)
	if r.reqErr != nil {
		return
	}

	for k, v := range *r.hParams {
		r.req.Header.Add(k, v)
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
		r.resErr = err
	}

	r.res, r.resErr = c.Do(r.req)
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
		return
	}

}

// ------------------------------------------------------------------------- //

func (r request) Do() ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Execute the HTTP request
	*/

	r.setUpReq()
	if r.reqErr != nil {
		return nil, r.reqErr
	}

	r.fetchReq()
	if r.resErr != nil {
		return nil, r.resErr
	}

	r.parseRes()
	if r.cntErr != nil {
		return nil, r.cntErr
	}

	return r.cnt, nil

}

// ------------------------------------------------------------------------- //
