package query_handler

import (
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

/*
struct for header on our webmotors queries
it is our way through capcha and into the API
*/

type QueryClient struct {
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

//generate the http client, with tor proxy if necessary

func (client_info QueryClient) GenerateClient() *http.Client {

	torProxy := client_info.createProxy()
	tbTransport := &http.Transport{Dial: torProxy.Dial}
	client := &http.Client{Transport: tbTransport}
	return client

}

//Tor Proxy if user wants to
func (client_info QueryClient) createProxy() proxy.Dialer {
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

	client := &http.Client{}
	return client

}

func (client_info QueryClient) CreateRequest(url string) *http.Request {
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
