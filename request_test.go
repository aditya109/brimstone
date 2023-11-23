package brimstone

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSplintParameters_Strike(t *testing.T) {
	type fields struct {
		ClientParams  ClientParams
		RequestParams RequestParams
		Meta          map[string]interface{}
	}
	type args struct {
		py *Pyre
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		want1   *http.Response
		wantErr bool
	}{
		{
			name: "sanity test case for splint parameter working with strike",
			fields: fields{
				ClientParams: ClientParams{
					BaseURL:                        "https://reqres.in",
					RequestTimeoutInSeconds:        20,
					Name:                           "REQRES",
					InsecureSkipVerify:             false,
					ShouldHaveAuthenticationHeader: false,
				},
				RequestParams: RequestParams{
					URLParams: URLParams{
						Path:       "/api/users/:userId",
						MethodType: http.MethodGet,
					},
					HTTPVariables: HTTPVariables{
						UriParams: map[string]string{
							":userId": "2",
						},
						Headers: map[string][]string{
							"x-gatekeeper-token": {"NdUrqpKWbgbToUfjZR3Gv2C5"},
						},
					},
				},
				Meta: map[string]interface{}{
					"trace_id": "082d37e5-6dad-4396-b76c-f399e7af859c",
				},
			},
			args: args{
				py: nil,
			},
			want:    []byte(`{"data":{"id":2,"email":"janet.weaver@reqres.in","first_name":"Janet","last_name":"Weaver","avatar":"https://reqres.in/img/faces/2-image.jpg"},"support":{"url":"https://reqres.in/#support-heading","text":"To keep ReqRes free, contributions towards server costs are appreciated!"}}`),
			wantErr: false,
		},
		{
			name: "negative sanity test case for splint parameter not working with strike",
			fields: fields{
				ClientParams: ClientParams{
					BaseURL:                        "https://rqres.in",
					RequestTimeoutInSeconds:        20,
					Name:                           "REQRES",
					InsecureSkipVerify:             false,
					ShouldHaveAuthenticationHeader: false,
				},
				RequestParams: RequestParams{
					URLParams: URLParams{
						Path:       "/api/users/:userId",
						MethodType: http.MethodGet,
					},
					HTTPVariables: HTTPVariables{
						UriParams: map[string]string{
							":userId": "2",
						},
						Headers: map[string][]string{
							"x-gatekeeper-token": {"NdUrqpKWbgbToUfjZR3Gv2C5"},
						},
					},
				},
				Meta: map[string]interface{}{
					"trace_id": "082d37e5-6dad-4396-b76c-f399e7af859c",
				},
			},
			args: args{
				py: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SplintParameters{
				ClientParams:  tt.fields.ClientParams,
				RequestParams: tt.fields.RequestParams,
				Meta:          tt.fields.Meta,
			}
			got, _, err := p.Strike(tt.args.py)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplintParameters.Strike() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplintParameters.Strike() got = %v, want %v", got, tt.want)
			}
		})
	}
}
