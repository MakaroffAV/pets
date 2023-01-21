package req

import (
	"net/http"
	"net/url"
	"time"
)

func setUpClient() (*http.Client, error) {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Setting up request client
	*/

	u, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Timeout:   time.Second * 5,
		Transport: &http.Transport{Proxy: http.ProxyURL(u)},
	}, nil

}
