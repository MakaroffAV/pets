package req

import (
	"net/http"
)

type Request struct {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Data structure for
				describing the HTTP request`s params
	*/

	url string

	Href *string
	Meth *string

	ParamsH *map[string]string
	ParamsQ *map[string]string
}

func (r *Request) setUpReq() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Setting up the request config
	*/

	u := Url{}

	q, err := http.NewRequest(*r.Meth, r.url, nil)
	if err != nil {
		return
	}

}

func (r *Request) fetchReq() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Fetching HTTP request
	*/

}

func (r *Request) parseReq() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Parsing HTTP request
	*/
}

func (r Request) Do() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Execute the HTTP request
	*/

	q, err := r.setUpReq()
	if err != nil {
		return nil, err
	}

	r, err := fetchReq()
	if err != nil {
		return nil, err
	}

	b, err := parseRes()
	if err != nil {
		return nil, err
	}

	return b, nil

}
