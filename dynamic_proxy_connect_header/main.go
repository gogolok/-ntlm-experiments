package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	header := make(http.Header)
	header.Add("Proxy-Authorization", "gogolok-0")

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		ProxyConnectHeader:  header,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	url := "http://www.focus.de"

	response, err := netClient.Get(url)
	if err != nil {
		log.Printf("GET 0: err = %v\n", err)
		return
	}
	log.Printf("GET 0 response = %v\n", response)

	// second round
	url = "http://www.spiegel.de"
	header.Set("Proxy-Authorization", "gogolok-3")
	netTransport.ProxyConnectHeader = header

	response, err = netClient.Get(url)
	if err != nil {
		log.Printf("GET 1: err = %v\n", err)
		return
	}
	log.Printf("GET 1 response = %v\n", response)

	log.Printf("Done.\n")
}
