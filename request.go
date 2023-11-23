package brimstone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (p SplintParameters) initPyre() Pyre {
	// create a pyre object
	return Pyre{
		cltCtrct: p.ClientParams.createNewHTTPClient(p.ClientParams.ShouldHaveAuthenticationHeader),
		Request:  nil,
		Curl:     "",
		Meta: map[string]interface{}{
			name: p.ClientParams.Name,
		},
	}
}

// Kindle creates a Pyre object which ought to be used for calling Strike
func (p SplintParameters) Kindle() (Pyre, error) {
	// base path
	var pyre = p.initPyre()
	var body *bytes.Buffer
	var request *http.Request
	var err error

	URL := fmt.Sprintf("%s%s", pyre.cltCtrct.baseURL, p.RequestParams.URLParams.Path)
	var httpHeaders = http.Header{}
	if p.RequestParams.HTTPVariables.UriParams != nil {
		interpolateURLWithURLParams(&URL, p.RequestParams.HTTPVariables.UriParams)
	}
	addHeaders(&httpHeaders, pyre.cltCtrct.basicAuthKey, p.RequestParams.HTTPVariables.Headers)
	if p.RequestParams.HTTPVariables.Payload != nil {
		payloadJSON, err := json.Marshal(p.RequestParams.HTTPVariables.Payload)
		if err != nil {
			return Pyre{}, fmt.Errorf("error while marshalling JSON payload, %v", err)
		}
		body = bytes.NewBuffer(payloadJSON)
		request, err = http.NewRequest(p.RequestParams.URLParams.MethodType, URL, body)
		if err != nil {
			log.Print(err.Error())
			return Pyre{}, err
		}
	} else {
		req, err := http.NewRequest(p.RequestParams.URLParams.MethodType, URL, nil)
		if err != nil {
			log.Print(err.Error())
			return Pyre{}, err
		}
		request = req
	}

	request.URL.RawQuery = url.Values(p.RequestParams.HTTPVariables.QueryParams).Encode()
	request.Header = httpHeaders
	pyre.Request = request

	if pyre.Curl, err = getCurlForRequest(request); err != nil {
		return Pyre{}, fmt.Errorf("error while getting a curl, err: %v", err)
	}

	return pyre, nil
}

// Strike can be calling with SplintParameters (callee object) with/without Pyre to execute the HTTP request
func (p SplintParameters) Strike(py *Pyre) ([]byte, *http.Response, error) {
	var pyre Pyre
	var err error
	if py != nil {
		pyre = *py
	} else {
		pyre, err = p.Kindle()
		if err != nil {
			return nil, nil, err
		}
	}
	startTime := time.Now()
	response, err := pyre.cltCtrct.client.Do(pyre.Request)
	fmt.Print("\n")
	for k, v := range p.Meta {
		fmt.Printf("%s: %s; ", k, v)
	}
	fmt.Printf("latency: %d ms; outgoing cURL: %s\n", time.Since(startTime).Milliseconds(), pyre.Curl)

	return processResponse(response, err, pyre.Curl, time.Duration(p.ClientParams.RequestTimeoutInSeconds), p.ClientParams.Name)
}
