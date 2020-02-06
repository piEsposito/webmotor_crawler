package query_handler

import (
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

type QueryClient struct {
	/*
		struct for header on our webmotors queries
		it is our way through capcha and into the API

		we can also set a proxy url if we want to disguise our IP
	*/
	Accept                    string
	Accept_encoding           string
	Accept_language           string
	Connection                string
	Host                      string
	TE                        string
	Upgrade_insecure_requests string
	User_agent                string
	ProxyUrl                  string
}

func (client_info QueryClient) GenerateClient() *http.Client {
	/*
		This method allows our QueryClient to create a http client make the requests
		This one proxies our connection
	*/
	torProxy := client_info.createProxy()
	tbTransport := &http.Transport{Dial: torProxy.Dial}
	client := &http.Client{Transport: tbTransport}
	return client

}

func (client_info QueryClient) createProxy() proxy.Dialer {
	/*
		If we select to generate a proxied client, this method creates the client proxy
	*/
	proxyUrl, err := url.Parse(client_info.ProxyUrl)
	if err != nil {
		panic(err)
	}
	_ = err

	tbDialer, err := proxy.FromURL(proxyUrl, proxy.Direct)

	if err != nil {
		panic(err)
	}

	return tbDialer

}

func (client_info QueryClient) GenerateNoProxiedClient() *http.Client {
	/*
		If we want to use no proxy, just generate a vanilla client
	*/
	client := &http.Client{}
	return client

}

func (client_info QueryClient) CreateRequest(url string) *http.Request {
	/*
		This method creates the request from the URL and QueryClient header info
	*/
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//setting our masked headers
	req.Header.Set("Accept", client_info.Accept)
	req.Header.Set("Accept-Encoding", client_info.Accept_encoding)
	req.Header.Set("Accept-Language", client_info.Accept_language)
	req.Header.Set("Connection", client_info.Connection)

	req.Header.Set("Host", client_info.Host)
	req.Header.Set("TE", client_info.TE)
	req.Header.Set("Upgrade-Insecure-Requests", client_info.Upgrade_insecure_requests)
	req.Header.Set("User-Agent", client_info.Upgrade_insecure_requests)

	return req
}
