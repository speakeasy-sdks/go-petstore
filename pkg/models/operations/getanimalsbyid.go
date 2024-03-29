// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"PB/v3/pkg/models/shared"
	"net/http"
)

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
	Animals *shared.Animals
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
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
