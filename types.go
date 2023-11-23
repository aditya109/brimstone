package brimstone

import (
	"net/http"
)

const (
	name string = "NAME"
)

type clientConstruct struct {
	client       *http.Client
	basicAuthKey *string
	baseURL      string
}

// Pyre provides a construct which is formal parameter of Strike(), containing exportable and mellable fields, like Meta, Request, Curl.
type Pyre struct {
	cltCtrct clientConstruct
	Meta     map[string]interface{}
	Request  *http.Request
	Curl     string
}

// ClientParams are a part of SplintParameters which serve as the input parameters of HTTP request, using host details of target.
type ClientParams struct {
	BaseURL                        string
	Username                       string
	Password                       string
	RequestTimeoutInSeconds        int
	Name                           string
	InsecureSkipVerify             bool
	ShouldHaveAuthenticationHeader bool
}

// RequestParams are a part of SplintParameters which serve as the counterpart of ClientParams, contains other necessary parameters of HTTP request.
type RequestParams struct {
	URLParams     URLParams
	HTTPVariables HTTPVariables
}

// URLParams contains path and methodtype values within itself.
type URLParams struct {
	Path       string
	MethodType string
}

// HTTPVariables contains various informatives like query parameters, uri parameters, headers and payload.
type HTTPVariables struct {
	QueryParams map[string][]string
	UriParams   map[string]string //nolint:all
	Headers     http.Header
	Payload     map[string]interface{}
}

// SplintParameters defines the structure of the callee object.
type SplintParameters struct {
	ClientParams  ClientParams
	RequestParams RequestParams
	Meta          map[string]interface{}
}
