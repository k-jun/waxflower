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




type Ticket struct {

	Id string `json:"id"`

	Price int64 `json:"price"`

	UserId string `json:"userId,omitempty"`

	GameId string `json:"gameId"`

	SeatId string `json:"seatId"`
}

// AssertTicketRequired checks if the required fields are not zero-ed
func AssertTicketRequired(obj Ticket) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"price": obj.Price,
		"gameId": obj.GameId,
		"seatId": obj.SeatId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTicketConstraints checks if the values respects the defined constraints
func AssertTicketConstraints(obj Ticket) error {
	return nil
}
