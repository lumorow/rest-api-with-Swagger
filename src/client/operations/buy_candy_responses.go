// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuyCandyReader is a Reader for the BuyCandy structure.
type BuyCandyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuyCandyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewBuyCandyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBuyCandyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewBuyCandyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBuyCandyCreated creates a BuyCandyCreated with default headers values
func NewBuyCandyCreated() *BuyCandyCreated {
	return &BuyCandyCreated{}
}

/*
BuyCandyCreated describes a response with status code 201, with default header values.

purchase successful
*/
type BuyCandyCreated struct {
	Payload *BuyCandyCreatedBody
}

// IsSuccess returns true when this buy candy created response has a 2xx status code
func (o *BuyCandyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this buy candy created response has a 3xx status code
func (o *BuyCandyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this buy candy created response has a 4xx status code
func (o *BuyCandyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this buy candy created response has a 5xx status code
func (o *BuyCandyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this buy candy created response a status code equal to that given
func (o *BuyCandyCreated) IsCode(code int) bool {
	return code == 201
}

func (o *BuyCandyCreated) Error() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyCreated  %+v", 201, o.Payload)
}

func (o *BuyCandyCreated) String() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyCreated  %+v", 201, o.Payload)
}

func (o *BuyCandyCreated) GetPayload() *BuyCandyCreatedBody {
	return o.Payload
}

func (o *BuyCandyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(BuyCandyCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuyCandyBadRequest creates a BuyCandyBadRequest with default headers values
func NewBuyCandyBadRequest() *BuyCandyBadRequest {
	return &BuyCandyBadRequest{}
}

/*
BuyCandyBadRequest describes a response with status code 400, with default header values.

some error in input data
*/
type BuyCandyBadRequest struct {
	Payload *BuyCandyBadRequestBody
}

// IsSuccess returns true when this buy candy bad request response has a 2xx status code
func (o *BuyCandyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this buy candy bad request response has a 3xx status code
func (o *BuyCandyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this buy candy bad request response has a 4xx status code
func (o *BuyCandyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this buy candy bad request response has a 5xx status code
func (o *BuyCandyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this buy candy bad request response a status code equal to that given
func (o *BuyCandyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BuyCandyBadRequest) Error() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyBadRequest  %+v", 400, o.Payload)
}

func (o *BuyCandyBadRequest) String() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyBadRequest  %+v", 400, o.Payload)
}

func (o *BuyCandyBadRequest) GetPayload() *BuyCandyBadRequestBody {
	return o.Payload
}

func (o *BuyCandyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(BuyCandyBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuyCandyPaymentRequired creates a BuyCandyPaymentRequired with default headers values
func NewBuyCandyPaymentRequired() *BuyCandyPaymentRequired {
	return &BuyCandyPaymentRequired{}
}

/*
BuyCandyPaymentRequired describes a response with status code 402, with default header values.

not enough money
*/
type BuyCandyPaymentRequired struct {
	Payload *BuyCandyPaymentRequiredBody
}

// IsSuccess returns true when this buy candy payment required response has a 2xx status code
func (o *BuyCandyPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this buy candy payment required response has a 3xx status code
func (o *BuyCandyPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this buy candy payment required response has a 4xx status code
func (o *BuyCandyPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this buy candy payment required response has a 5xx status code
func (o *BuyCandyPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this buy candy payment required response a status code equal to that given
func (o *BuyCandyPaymentRequired) IsCode(code int) bool {
	return code == 402
}

func (o *BuyCandyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyPaymentRequired  %+v", 402, o.Payload)
}

func (o *BuyCandyPaymentRequired) String() string {
	return fmt.Sprintf("[POST /buy_candy][%d] buyCandyPaymentRequired  %+v", 402, o.Payload)
}

func (o *BuyCandyPaymentRequired) GetPayload() *BuyCandyPaymentRequiredBody {
	return o.Payload
}

func (o *BuyCandyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(BuyCandyPaymentRequiredBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
BuyCandyBadRequestBody buy candy bad request body
swagger:model BuyCandyBadRequestBody
*/
type BuyCandyBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this buy candy bad request body
func (o *BuyCandyBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy bad request body based on context it is used
func (o *BuyCandyBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyBadRequestBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BuyCandyBody buy candy body
swagger:model BuyCandyBody
*/
type BuyCandyBody struct {

	// number of candy
	// Required: true
	CandyCount *int64 `json:"candyCount"`

	// kind of candy
	// Required: true
	CandyType *string `json:"candyType"`

	// amount of money put into vending machine
	// Required: true
	Money *int64 `json:"money"`
}

// Validate validates this buy candy body
func (o *BuyCandyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCandyCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCandyType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMoney(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BuyCandyBody) validateCandyCount(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"candyCount", "body", o.CandyCount); err != nil {
		return err
	}

	return nil
}

func (o *BuyCandyBody) validateCandyType(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"candyType", "body", o.CandyType); err != nil {
		return err
	}

	return nil
}

func (o *BuyCandyBody) validateMoney(formats strfmt.Registry) error {

	if err := validate.Required("order"+"."+"money", "body", o.Money); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this buy candy body based on context it is used
func (o *BuyCandyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BuyCandyCreatedBody buy candy created body
swagger:model BuyCandyCreatedBody
*/
type BuyCandyCreatedBody struct {

	// change
	Change int64 `json:"change,omitempty"`

	// thanks
	Thanks string `json:"thanks,omitempty"`
}

// Validate validates this buy candy created body
func (o *BuyCandyCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy created body based on context it is used
func (o *BuyCandyCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyCreatedBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
BuyCandyPaymentRequiredBody buy candy payment required body
swagger:model BuyCandyPaymentRequiredBody
*/
type BuyCandyPaymentRequiredBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this buy candy payment required body
func (o *BuyCandyPaymentRequiredBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this buy candy payment required body based on context it is used
func (o *BuyCandyPaymentRequiredBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BuyCandyPaymentRequiredBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BuyCandyPaymentRequiredBody) UnmarshalBinary(b []byte) error {
	var res BuyCandyPaymentRequiredBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
