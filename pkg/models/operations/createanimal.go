// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"PB/v3/pkg/models/shared"
	"net/http"
)

type CreateAnimalRequestBody struct {
	Age   *int64 `json:"age,omitempty"`
	Color string `json:"color"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

func (o *CreateAnimalRequestBody) GetAge() *int64 {
	if o == nil {
		return nil
	}
	return o.Age
}

func (o *CreateAnimalRequestBody) GetColor() string {
	if o == nil {
		return ""
	}
	return o.Color
}

func (o *CreateAnimalRequestBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateAnimalRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type CreateAnimalResponse struct {
	// OK
	Animals *shared.Animals
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAnimalResponse) GetAnimals() *shared.Animals {
	if o == nil {
		return nil
	}
	return o.Animals
}

func (o *CreateAnimalResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAnimalResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAnimalResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
