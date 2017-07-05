package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/anynines/go-ntlm-auth/ntlm"
)

var (
	proxy  = "https://10.0.0.45:3128"
	target = "http://google.com"
)

func main() {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		fmt.Printf("Failed to create new request object: %v\n", err.Error())
		return
	}

	proxyUrl, err := url.Parse(proxy)
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	res, err := ntlm.DoNTLMRequest(myClient, req)
	if err != nil {
		fmt.Printf("NTLM request failed: %v\n", err.Error())
		return
	}
	fmt.Printf("NTLM seemed to work\n")
	fmt.Printf("res = %v\n", res)
}
