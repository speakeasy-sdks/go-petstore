// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"PB/pkg/utils"
	"errors"
)

type ComplexObjectDataBirds2 struct {
}

type ComplexObjectDataBirdsType string

const (
	ComplexObjectDataBirdsTypeArrayOfany              ComplexObjectDataBirdsType = "arrayOfany"
	ComplexObjectDataBirdsTypeComplexObjectDataBirds2 ComplexObjectDataBirdsType = "ComplexObject_data_birds_2"
)

type ComplexObjectDataBirds struct {
	ArrayOfany              []interface{}
	ComplexObjectDataBirds2 *ComplexObjectDataBirds2

	Type ComplexObjectDataBirdsType
}

func CreateComplexObjectDataBirdsArrayOfany(arrayOfany []interface{}) ComplexObjectDataBirds {
	typ := ComplexObjectDataBirdsTypeArrayOfany

	return ComplexObjectDataBirds{
		ArrayOfany: arrayOfany,
		Type:       typ,
	}
}

func CreateComplexObjectDataBirdsComplexObjectDataBirds2(complexObjectDataBirds2 ComplexObjectDataBirds2) ComplexObjectDataBirds {
	typ := ComplexObjectDataBirdsTypeComplexObjectDataBirds2

	return ComplexObjectDataBirds{
		ComplexObjectDataBirds2: &complexObjectDataBirds2,
		Type:                    typ,
	}
}

func (u *ComplexObjectDataBirds) UnmarshalJSON(data []byte) error {

	complexObjectDataBirds2 := ComplexObjectDataBirds2{}
	if err := utils.UnmarshalJSON(data, &complexObjectDataBirds2, "", true, true); err == nil {
		u.ComplexObjectDataBirds2 = &complexObjectDataBirds2
		u.Type = ComplexObjectDataBirdsTypeComplexObjectDataBirds2
		return nil
	}

	arrayOfany := []interface{}{}
	if err := utils.UnmarshalJSON(data, &arrayOfany, "", true, true); err == nil {
		u.ArrayOfany = arrayOfany
		u.Type = ComplexObjectDataBirdsTypeArrayOfany
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ComplexObjectDataBirds) MarshalJSON() ([]byte, error) {
	if u.ArrayOfany != nil {
		return utils.MarshalJSON(u.ArrayOfany, "", true)
	}

	if u.ComplexObjectDataBirds2 != nil {
		return utils.MarshalJSON(u.ComplexObjectDataBirds2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ComplexObjectDataCreatedDateType string

const (
	ComplexObjectDataCreatedDateTypeInteger ComplexObjectDataCreatedDateType = "integer"
	ComplexObjectDataCreatedDateTypeStr     ComplexObjectDataCreatedDateType = "str"
)

type ComplexObjectDataCreatedDate struct {
	Integer *int64
	Str     *string

	Type ComplexObjectDataCreatedDateType
}

func CreateComplexObjectDataCreatedDateInteger(integer int64) ComplexObjectDataCreatedDate {
	typ := ComplexObjectDataCreatedDateTypeInteger

	return ComplexObjectDataCreatedDate{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateComplexObjectDataCreatedDateStr(str string) ComplexObjectDataCreatedDate {
	typ := ComplexObjectDataCreatedDateTypeStr

	return ComplexObjectDataCreatedDate{
		Str:  &str,
		Type: typ,
	}
}

func (u *ComplexObjectDataCreatedDate) UnmarshalJSON(data []byte) error {

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ComplexObjectDataCreatedDateTypeInteger
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ComplexObjectDataCreatedDateTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ComplexObjectDataCreatedDate) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ComplexObjectDataUpdatedDateType string

const (
	ComplexObjectDataUpdatedDateTypeInteger ComplexObjectDataUpdatedDateType = "integer"
	ComplexObjectDataUpdatedDateTypeNumber  ComplexObjectDataUpdatedDateType = "number"
)

type ComplexObjectDataUpdatedDate struct {
	Integer *int64
	Number  *float64

	Type ComplexObjectDataUpdatedDateType
}

func CreateComplexObjectDataUpdatedDateInteger(integer int64) ComplexObjectDataUpdatedDate {
	typ := ComplexObjectDataUpdatedDateTypeInteger

	return ComplexObjectDataUpdatedDate{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateComplexObjectDataUpdatedDateNumber(number float64) ComplexObjectDataUpdatedDate {
	typ := ComplexObjectDataUpdatedDateTypeNumber

	return ComplexObjectDataUpdatedDate{
		Number: &number,
		Type:   typ,
	}
}

func (u *ComplexObjectDataUpdatedDate) UnmarshalJSON(data []byte) error {

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ComplexObjectDataUpdatedDateTypeInteger
		return nil
	}

	number := float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = ComplexObjectDataUpdatedDateTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ComplexObjectDataUpdatedDate) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ComplexObjectData struct {
	Animal      []Animals                     `json:"animal,omitempty"`
	Birds       *ComplexObjectDataBirds       `json:"birds,omitempty"`
	CreatedDate *ComplexObjectDataCreatedDate `json:"createdDate,omitempty"`
	UpdatedDate *ComplexObjectDataUpdatedDate `json:"updatedDate,omitempty"`
}

func (o *ComplexObjectData) GetAnimal() []Animals {
	if o == nil {
		return nil
	}
	return o.Animal
}

func (o *ComplexObjectData) GetBirds() *ComplexObjectDataBirds {
	if o == nil {
		return nil
	}
	return o.Birds
}

func (o *ComplexObjectData) GetCreatedDate() *ComplexObjectDataCreatedDate {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *ComplexObjectData) GetUpdatedDate() *ComplexObjectDataUpdatedDate {
	if o == nil {
		return nil
	}
	return o.UpdatedDate
}

type ComplexObjectMeta2 struct {
	PageNumber *string `json:"pageNumber,omitempty"`
}

func (o *ComplexObjectMeta2) GetPageNumber() *string {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

type ComplexObjectMetaType string

const (
	ComplexObjectMetaTypePagination         ComplexObjectMetaType = "Pagination"
	ComplexObjectMetaTypeComplexObjectMeta2 ComplexObjectMetaType = "ComplexObject_meta_2"
)

type ComplexObjectMeta struct {
	Pagination         *Pagination
	ComplexObjectMeta2 *ComplexObjectMeta2

	Type ComplexObjectMetaType
}

func CreateComplexObjectMetaPagination(pagination Pagination) ComplexObjectMeta {
	typ := ComplexObjectMetaTypePagination

	return ComplexObjectMeta{
		Pagination: &pagination,
		Type:       typ,
	}
}

func CreateComplexObjectMetaComplexObjectMeta2(complexObjectMeta2 ComplexObjectMeta2) ComplexObjectMeta {
	typ := ComplexObjectMetaTypeComplexObjectMeta2

	return ComplexObjectMeta{
		ComplexObjectMeta2: &complexObjectMeta2,
		Type:               typ,
	}
}

func (u *ComplexObjectMeta) UnmarshalJSON(data []byte) error {

	complexObjectMeta2 := ComplexObjectMeta2{}
	if err := utils.UnmarshalJSON(data, &complexObjectMeta2, "", true, true); err == nil {
		u.ComplexObjectMeta2 = &complexObjectMeta2
		u.Type = ComplexObjectMetaTypeComplexObjectMeta2
		return nil
	}

	pagination := Pagination{}
	if err := utils.UnmarshalJSON(data, &pagination, "", true, true); err == nil {
		u.Pagination = &pagination
		u.Type = ComplexObjectMetaTypePagination
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ComplexObjectMeta) MarshalJSON() ([]byte, error) {
	if u.Pagination != nil {
		return utils.MarshalJSON(u.Pagination, "", true)
	}

	if u.ComplexObjectMeta2 != nil {
		return utils.MarshalJSON(u.ComplexObjectMeta2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ComplexObject struct {
	Data *ComplexObjectData `json:"data,omitempty"`
	Meta *ComplexObjectMeta `json:"meta,omitempty"`
	Name *string            `json:"name,omitempty"`
}

func (o *ComplexObject) GetData() *ComplexObjectData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ComplexObject) GetMeta() *ComplexObjectMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *ComplexObject) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
