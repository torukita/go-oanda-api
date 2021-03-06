// Code generated by go-swagger; DO NOT EDIT.

package oanda

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/torukita/go-oanda-api/models"
)

// ReplaceOrderReader is a Reader for the ReplaceOrder structure.
type ReplaceOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewReplaceOrderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewReplaceOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewReplaceOrderMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceOrderCreated creates a ReplaceOrderCreated with default headers values
func NewReplaceOrderCreated() *ReplaceOrderCreated {
	return &ReplaceOrderCreated{}
}

/*ReplaceOrderCreated handles this case with default header values.

The Order was successfully cancelled and replaced
*/
type ReplaceOrderCreated struct {
	/*A link to the replacing Order
	 */
	Location string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ReplaceOrderCreatedBody
}

func (o *ReplaceOrderCreated) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}][%d] replaceOrderCreated  %+v", 201, o.Payload)
}

func (o *ReplaceOrderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ReplaceOrderCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceOrderBadRequest creates a ReplaceOrderBadRequest with default headers values
func NewReplaceOrderBadRequest() *ReplaceOrderBadRequest {
	return &ReplaceOrderBadRequest{}
}

/*ReplaceOrderBadRequest handles this case with default header values.

The Order specification was invalid
*/
type ReplaceOrderBadRequest struct {
	Payload *models.ReplaceOrderBadRequestBody
}

func (o *ReplaceOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}][%d] replaceOrderBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplaceOrderBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceOrderUnauthorized creates a ReplaceOrderUnauthorized with default headers values
func NewReplaceOrderUnauthorized() *ReplaceOrderUnauthorized {
	return &ReplaceOrderUnauthorized{}
}

/*ReplaceOrderUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ReplaceOrderUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ReplaceOrderUnauthorizedBody
}

func (o *ReplaceOrderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}][%d] replaceOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *ReplaceOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ReplaceOrderUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceOrderNotFound creates a ReplaceOrderNotFound with default headers values
func NewReplaceOrderNotFound() *ReplaceOrderNotFound {
	return &ReplaceOrderNotFound{}
}

/*ReplaceOrderNotFound handles this case with default header values.

The Account or Order specified does not exist.
*/
type ReplaceOrderNotFound struct {
	Payload *models.ReplaceOrderNotFoundBody
}

func (o *ReplaceOrderNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}][%d] replaceOrderNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplaceOrderNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceOrderMethodNotAllowed creates a ReplaceOrderMethodNotAllowed with default headers values
func NewReplaceOrderMethodNotAllowed() *ReplaceOrderMethodNotAllowed {
	return &ReplaceOrderMethodNotAllowed{}
}

/*ReplaceOrderMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ReplaceOrderMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ReplaceOrderMethodNotAllowedBody
}

func (o *ReplaceOrderMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/orders/{orderSpecifier}][%d] replaceOrderMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ReplaceOrderMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ReplaceOrderMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
