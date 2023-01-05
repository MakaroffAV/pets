package tnreq

import (
	"b/pkg/cslog"
	"io/ioutil"
	"net/http"
)

// ------------------------------------------------------------------------- //

func getClient() (*http.Client, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Get HTTP request client
	*/

	t = &http.Transport{Proxy: http.ProxyURL("socks5://127.0.0.1:9050")}

	return &http.Client{Transport: t}, nil

}

// ------------------------------------------------------------------------- //

func Get(u string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Making HTTP/GET request
	*/

	c, err := getClient()
	if err != nil {
		cslog.Fail(
			`F002`,
		)
		return nil, err
	}

	r, err := c.Get(u)
	if err != nil {
		cslog.Fail(
			`F003`,
		)
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		cslog.Fail(
			`F004`,
		)
		return nil, err
	}

	return b, nil

}

// ------------------------------------------------------------------------- //
