package brimstone

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/aditya109/shazam"
	"golang.org/x/exp/maps"
)

func interpolateURLWithURLParams(url *string, urlParams map[string]string) {
	for key, value := range urlParams {
		*url = strings.Replace(*url, key, value, 1)
	}
}

func addHeaders(h *http.Header, authKey *string, extraHeaders http.Header) {
	if authKey != nil {
		(*h).Add("Authorization", "Basic "+*authKey)
	}
	maps.Copy(*h, extraHeaders)
}

func getCurlForRequest(request *http.Request) (string, error) {
	curl, err := shazam.Boom(request)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return curl, nil
}

func processResponse(response *http.Response, incErr error, curl string, defaultTimeout time.Duration, name string) ([]byte, *http.Response, error) {
	var body []byte
	var err error
	if incErr != nil {
		var apiError net.Error
		if errors.As(incErr, &apiError) && apiError.Timeout() {
			return nil, response, fmt.Errorf("API timed-out(%d secs timeout), api error: %v, cURL: %s", defaultTimeout, apiError, curl)
		}
		log.Printf("found error in API, api error: %s, cURL: %s", apiError.Error(), curl)
		return nil, response, apiError
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			err = fmt.Errorf("found error while closing response body, error: %v, cURL: %s", err, curl)
		}
	}(response.Body)

	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, response, fmt.Errorf("found error while reading response body, error: %v, cURL: %s", err, curl)
		}
		body, err = io.ReadAll(reader)
		if err != nil {
			return nil, response, fmt.Errorf("found error while reading response body, error: %v, cURL: %s", err, curl)
		}
	} else {
		body, err = io.ReadAll(response.Body)
		if err != nil {
			return nil, response, fmt.Errorf("found error while reading response body, error: %v, cURL: %s", err, curl)
		}
	}
	if response.StatusCode == http.StatusBadGateway {
		log.Printf("found status code : 502 (STATUS_BAD_GATEWAY), cURL: %s", curl)
		return nil, response, fmt.Errorf("found status code : 502 (STATUS_BAD_GATEWAY), service_name: %s", name)
	} else if response.StatusCode == http.StatusServiceUnavailable {
		log.Printf("found status code : 503 (STATUS_SERVICE_UNAVAILABLE), cURL: %s", curl)
		return nil, response, fmt.Errorf("found status code : 503 (STATUS_SERVICE_UNAVAILABLE), service_name: %s", name)
	} else if response.StatusCode > 399 {
		log.Printf("found status code > 399, cURL: %s", curl)
	}

	return body, response, err
}
