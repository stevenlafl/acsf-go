package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/proxy"
)

type Response struct {
	Body string
	Code int
}

func APICall(APIEnvironment string, APIMethod string, APIPayload string) Response {

	// @todo don't hardcode this
	var APIUrls map[string]string = GetAPIUrls()

	if _, ok := APIUrls[APIEnvironment]; !ok {
		return Response{
			Body: fmt.Sprintf("Unknown Environment:  %s", APIEnvironment),
			Code: 0,
		}
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	// create a socks5 dialer
	if USE_PROXY {
		dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
		if err != nil {
			return Response{
				Body: err.Error(),
				Code: -1,
			}
		}
		// set our socks5 as the dialer
		httpTransport.Dial = dialer.Dial
	}

	httpTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// create a request

	var HTTP_METHOD string
	if len(APIPayload) > 0 {
		HTTP_METHOD = "POST"
	} else {
		HTTP_METHOD = "GET"
	}

	req, err := http.NewRequest(HTTP_METHOD, fmt.Sprintf("%s/%s", APIUrls[APIEnvironment], APIMethod), bytes.NewBufferString(APIPayload))
	if len(APIPayload) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return Response{
			Body: err.Error(),
			Code: -2,
		}
	}

	req.SetBasicAuth(API_USER, API_PASS)
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		return Response{
			Body: err.Error(),
			Code: -3,
		}
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{
			Body: err.Error(),
			Code: -4,
		}
	}

	return Response{
		Body: string(b),
		Code: resp.StatusCode,
	}
}
