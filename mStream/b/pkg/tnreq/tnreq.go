package tnreq

import (
	"b/pkg/cslog"
	"log"

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

func Get(u string, p map[string]string) ([]byte, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Making HTTP/GET request via tor network
	*/

	c, err := client()
	if err != nil {
		cslog.Fail(
			"dfd",
			"DF",
		)
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:108.0) Gecko/20100101 Firefox/108.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Alt-Used", "www.youtube.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "VISITOR_INFO1_LIVE=W6M6X2JqgEY; PREF=tz=Europe.Moscow&f4=4000000&f6=40000000&f7=100; GPS=1; YSC=4ZZiorEUrhM")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("TE", "trailers")

	q := req.URL.Query()

	for i, v := range p {
		q.Add(i, v)
	}
	req.URL.RawQuery = q.Encode()

	r, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

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
