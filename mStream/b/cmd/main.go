package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/textproto"
	"net/url"
)

// URL to fetch
var webUrl string = "https://api.ipify.org"

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

func main() {

	text, err := textproto.Dial("tcp", torProxy)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(text)
	// Parse Tor proxy URL string to a URL type
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport}

	// Make request
	resp, err := client.Get(webUrl)
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	log.Println(string(body))
	log.Println("Return status code:", resp.StatusCode)

}
