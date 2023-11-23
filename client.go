package brimstone

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"time"
)

func (p ClientParams) createNewHTTPClient(shouldHaveAuthHeader bool) clientConstruct {
	var c *http.Client
	var basicAuthKey *string = nil

	c = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: p.InsecureSkipVerify,
			},
		},
		Timeout: time.Duration(p.RequestTimeoutInSeconds) * time.Second,
	}

	if shouldHaveAuthHeader {
		authKey := base64.StdEncoding.EncodeToString([]byte(p.Username + ":" + p.Password))
		basicAuthKey = &authKey
	}
	return clientConstruct{
		client:       c,
		basicAuthKey: basicAuthKey,
		baseURL:      p.BaseURL,
	}
}
