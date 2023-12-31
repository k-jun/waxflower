/*
 * WaxFlower
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 * Contact: keijun091221@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"errors"
)



type Seat struct {

	Id string `json:"id"`

	Sec int32 `json:"sec"`

	Col int32 `json:"col"`

	Row string `json:"row"`
}

// AssertSeatRequired checks if the required fields are not zero-ed
func AssertSeatRequired(obj Seat) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"sec": obj.Sec,
		"col": obj.Col,
		"row": obj.Row,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSeatConstraints checks if the values respects the defined constraints
func AssertSeatConstraints(obj Seat) error {
	if obj.Sec < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Sec > 9 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Col < 1 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Col > 99 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
