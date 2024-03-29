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




type TicketReserve struct {

	UserId string `json:"userId"`

	TicketId string `json:"ticketId"`
}

// AssertTicketReserveRequired checks if the required fields are not zero-ed
func AssertTicketReserveRequired(obj TicketReserve) error {
	elements := map[string]interface{}{
		"userId": obj.UserId,
		"ticketId": obj.TicketId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTicketReserveConstraints checks if the values respects the defined constraints
func AssertTicketReserveConstraints(obj TicketReserve) error {
	return nil
}
