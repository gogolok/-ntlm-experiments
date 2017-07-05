package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/anynines/go-ntlm-auth/ntlm"
)

var (
	proxy  = flag.String("proxy", "https://10.0.0.45:3128", "the NTLM proxy")
	target = flag.String("target", "http://google.com", "the web page to get")
)

func main() {
	flag.Parse()

	fmt.Printf("proxy = %v\n", *proxy)
	fmt.Printf("target = %v\n", *target)

	req, err := http.NewRequest("GET", *target, nil)
	if err != nil {
		fmt.Printf("Failed to create new request object: %v\n", err.Error())
		return
	}

	proxyUrl, err := url.Parse(*proxy)
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	resp, err := ntlm.DoNTLMRequest(myClient, req)
	if err != nil {
		fmt.Printf("NTLM request failed: %v\n", err.Error())
		return
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("Failed to dump response: %v", err)
		return
	}

	fmt.Printf("%q\n", dump)
}
