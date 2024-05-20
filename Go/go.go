package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	proxyURL, _ := url.Parse("http://USERNAME:PASSWORD@resi-global-gateways.databay.com:7676")
	resp, err := (&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}).Get("https://databay.com/cdn-cgi/trace")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(string(body))
	} else {
		fmt.Println("Error:", err)
	}
}