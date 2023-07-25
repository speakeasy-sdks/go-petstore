// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"PB/pkg/models/shared"
	"net/http"
)

type GetAnimalsByIDSecurity struct {
	Key1 string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *GetAnimalsByIDSecurity) GetKey1() string {
	if o == nil {
		return ""
	}
	return o.Key1
}

type GetAnimalsByIDRequest struct {
	Animals *shared.Animals `request:"mediaType=application/json"`
	ID      string          `pathParam:"style=simple,explode=false,name=id"`
	PerPage *int64          `queryParam:"style=form,explode=true,name=perPage"`
}

func (o *GetAnimalsByIDRequest) GetAnimals() *shared.Animals {
	if o == nil {
		return nil
	}
	return o.Animals
}

func (o *GetAnimalsByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetAnimalsByIDRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetAnimalsByIDResponse struct {
	// OK
	Animals     *shared.Animals
	ContentType string
	// Internal Server Error
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetAnimalsByIDResponse) GetAnimals() *shared.Animals {
	if o == nil {
		return nil
	}
	return o.Animals
}

func (o *GetAnimalsByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAnimalsByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetAnimalsByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAnimalsByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
