package tnreq

import (
	"b/pkg/cslog"

	"io/ioutil"
	"net/http"
	"net/url"
)

// ------------------------------------------------------------------------- //

func client() (*http.Client, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Get HTTP request client
	*/

	u, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		cslog.Fail(
			`F00000001`,
			err.Error(),
		)
		return nil, err
	}

	return &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(u)},
	}, nil

}

// ------------------------------------------------------------------------- //

func Get(u string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Making HTTP/GET request via tor network
	*/

	c, err := client()
	if err != nil {
		cslog.Fail(
			`F00000001`,
			err.Error(),
		)
		return nil, err
	}

	r, err := c.Get(u)
	if err != nil {
		cslog.Fail(
			`F00000001`,
			err.Error(),
		)
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		cslog.Fail(
			`F00000001`,
			err.Error(),
		)
		return nil, err
	}

	return b, nil

}

// ------------------------------------------------------------------------- //
