// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"PB/v2/pkg/models/shared"
	"PB/v2/pkg/utils"
	"errors"
	"net/http"
)

type GetAllLivingThingsRequest struct {
	// include all filters
	Filter []interface{} `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetAllLivingThingsRequest) GetFilter() []interface{} {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetAllLivingThingsPagination struct {
	PageNumber *int64 `json:"pageNumber,omitempty"`
}

func (o *GetAllLivingThingsPagination) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type GetAllLivingThings2 struct {
	Pagination *GetAllLivingThingsPagination `json:"pagination,omitempty"`
}

func (o *GetAllLivingThings2) GetPagination() *GetAllLivingThingsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type GetAllLivingThings1 struct {
	Pagination *shared.Pagination `json:"pagination,omitempty"`
}

func (o *GetAllLivingThings1) GetPagination() *shared.Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type GetAllLivingThingsMetaType string

const (
	GetAllLivingThingsMetaTypeGetAllLivingThings1 GetAllLivingThingsMetaType = "get-all-living-things_1"
	GetAllLivingThingsMetaTypeGetAllLivingThings2 GetAllLivingThingsMetaType = "get-all-living-things_2"
)

type GetAllLivingThingsMeta struct {
	GetAllLivingThings1 *GetAllLivingThings1
	GetAllLivingThings2 *GetAllLivingThings2

	Type GetAllLivingThingsMetaType
}

func CreateGetAllLivingThingsMetaGetAllLivingThings1(getAllLivingThings1 GetAllLivingThings1) GetAllLivingThingsMeta {
	typ := GetAllLivingThingsMetaTypeGetAllLivingThings1

	return GetAllLivingThingsMeta{
		GetAllLivingThings1: &getAllLivingThings1,
		Type:                typ,
	}
}

func CreateGetAllLivingThingsMetaGetAllLivingThings2(getAllLivingThings2 GetAllLivingThings2) GetAllLivingThingsMeta {
	typ := GetAllLivingThingsMetaTypeGetAllLivingThings2

	return GetAllLivingThingsMeta{
		GetAllLivingThings2: &getAllLivingThings2,
		Type:                typ,
	}
}

func (u *GetAllLivingThingsMeta) UnmarshalJSON(data []byte) error {

	getAllLivingThings1 := GetAllLivingThings1{}
	if err := utils.UnmarshalJSON(data, &getAllLivingThings1, "", true, true); err == nil {
		u.GetAllLivingThings1 = &getAllLivingThings1
		u.Type = GetAllLivingThingsMetaTypeGetAllLivingThings1
		return nil
	}

	getAllLivingThings2 := GetAllLivingThings2{}
	if err := utils.UnmarshalJSON(data, &getAllLivingThings2, "", true, true); err == nil {
		u.GetAllLivingThings2 = &getAllLivingThings2
		u.Type = GetAllLivingThingsMetaTypeGetAllLivingThings2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetAllLivingThingsMeta) MarshalJSON() ([]byte, error) {
	if u.GetAllLivingThings1 != nil {
		return utils.MarshalJSON(u.GetAllLivingThings1, "", true)
	}

	if u.GetAllLivingThings2 != nil {
		return utils.MarshalJSON(u.GetAllLivingThings2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Two struct {
	Animals []shared.Animals        `json:"animals,omitempty"`
	Meta    *GetAllLivingThingsMeta `json:"meta,omitempty"`
}

func (o *Two) GetAnimals() []shared.Animals {
	if o == nil {
		return nil
	}
	return o.Animals
}

func (o *Two) GetMeta() *GetAllLivingThingsMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type Pagination struct {
	PageNumber *int64 `json:"pageNumber,omitempty"`
}

func (o *Pagination) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type GetAllLivingThingsAnimalsMeta struct {
	Pagination *Pagination `json:"pagination,omitempty"`
}

func (o *GetAllLivingThingsAnimalsMeta) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type One struct {
	Birds []shared.Birds                 `json:"birds,omitempty"`
	Meta  *GetAllLivingThingsAnimalsMeta `json:"meta,omitempty"`
}

func (o *One) GetBirds() []shared.Birds {
	if o == nil {
		return nil
	}
	return o.Birds
}

func (o *One) GetMeta() *GetAllLivingThingsAnimalsMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type GetAllLivingThingsResponseBodyType string

const (
	GetAllLivingThingsResponseBodyTypeOne GetAllLivingThingsResponseBodyType = "1"
	GetAllLivingThingsResponseBodyTypeTwo GetAllLivingThingsResponseBodyType = "2"
)

// GetAllLivingThingsResponseBody - OK
type GetAllLivingThingsResponseBody struct {
	One *One
	Two *Two

	Type GetAllLivingThingsResponseBodyType
}

func CreateGetAllLivingThingsResponseBodyOne(one One) GetAllLivingThingsResponseBody {
	typ := GetAllLivingThingsResponseBodyTypeOne

	return GetAllLivingThingsResponseBody{
		One:  &one,
		Type: typ,
	}
}

func CreateGetAllLivingThingsResponseBodyTwo(two Two) GetAllLivingThingsResponseBody {
	typ := GetAllLivingThingsResponseBodyTypeTwo

	return GetAllLivingThingsResponseBody{
		Two:  &two,
		Type: typ,
	}
}

func (u *GetAllLivingThingsResponseBody) UnmarshalJSON(data []byte) error {

	one := One{}
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = &one
		u.Type = GetAllLivingThingsResponseBodyTypeOne
		return nil
	}

	two := Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = GetAllLivingThingsResponseBodyTypeTwo
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetAllLivingThingsResponseBody) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetAllLivingThingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	OneOf *GetAllLivingThingsResponseBody
}

func (o *GetAllLivingThingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllLivingThingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllLivingThingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllLivingThingsResponse) GetOneOf() *GetAllLivingThingsResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
