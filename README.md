# brimstone 🔥

----------

A wrapper for default HTTP-client.

> Sometimes different problems require different solutions.

## Getting started

   1. With Go module support, simply add the following import:
       ```go
       import "github.com/Kieraya/brimstone"
       ```
   2. To use the package, we need to apply using the following code:
      ```go
      package test
      
      import (
         "github.com/Kieraya/brimstone"
         "net/http"
      )
      
      func foo() {
         var params = brimstone.SplintParameters{
            ClientParams: brimstone.ClientParams{
               BaseURL:                        "https://reqres.in",
               Username:                       "",
               Password:                       "",
               RequestTimeoutInSeconds:        5,
               Name:                           "REQRES",
               InsecureSkipVerify:             false,
               ShouldHaveAuthenticationHeader: false,
            },
            RequestParams: brimstone.RequestParams{
               URLParams: brimstone.URLParams{
                  Path:       "/api/register",
                  MethodType: http.MethodPost,
               },
               HTTPVariables: brimstone.HTTPVariables{
                  QueryParams: nil,
                  UriParams:   nil,
                  Headers:     nil,
                  Payload: map[string]interface{}{
                     "email":    "eve.holt@reqres.in",
                     "password": "pistol",
                  },
               },
            },
         }
         bytes, response, err := params.Strike(nil)
         if err != nil {
         // handle error
         }
      }
      ```



