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

// CloseTradeReader is a Reader for the CloseTrade structure.
type CloseTradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseTradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloseTradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCloseTradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCloseTradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCloseTradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCloseTradeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloseTradeOK creates a CloseTradeOK with default headers values
func NewCloseTradeOK() *CloseTradeOK {
	return &CloseTradeOK{}
}

/*CloseTradeOK handles this case with default header values.

The Trade has been closed as requested
*/
type CloseTradeOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.CloseTradeOKBody
}

func (o *CloseTradeOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/trades/{tradeSpecifier}/close][%d] closeTradeOK  %+v", 200, o.Payload)
}

func (o *CloseTradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.CloseTradeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeBadRequest creates a CloseTradeBadRequest with default headers values
func NewCloseTradeBadRequest() *CloseTradeBadRequest {
	return &CloseTradeBadRequest{}
}

/*CloseTradeBadRequest handles this case with default header values.

The Trade cannot be closed as requested.
*/
type CloseTradeBadRequest struct {
	Payload *models.CloseTradeBadRequestBody
}

func (o *CloseTradeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/trades/{tradeSpecifier}/close][%d] closeTradeBadRequest  %+v", 400, o.Payload)
}

func (o *CloseTradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloseTradeBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeUnauthorized creates a CloseTradeUnauthorized with default headers values
func NewCloseTradeUnauthorized() *CloseTradeUnauthorized {
	return &CloseTradeUnauthorized{}
}

/*CloseTradeUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type CloseTradeUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.CloseTradeUnauthorizedBody
}

func (o *CloseTradeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/trades/{tradeSpecifier}/close][%d] closeTradeUnauthorized  %+v", 401, o.Payload)
}

func (o *CloseTradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.CloseTradeUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeNotFound creates a CloseTradeNotFound with default headers values
func NewCloseTradeNotFound() *CloseTradeNotFound {
	return &CloseTradeNotFound{}
}

/*CloseTradeNotFound handles this case with default header values.

The Account or Trade specified does not exist.
*/
type CloseTradeNotFound struct {
	Payload *models.CloseTradeNotFoundBody
}

func (o *CloseTradeNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/trades/{tradeSpecifier}/close][%d] closeTradeNotFound  %+v", 404, o.Payload)
}

func (o *CloseTradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloseTradeNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeMethodNotAllowed creates a CloseTradeMethodNotAllowed with default headers values
func NewCloseTradeMethodNotAllowed() *CloseTradeMethodNotAllowed {
	return &CloseTradeMethodNotAllowed{}
}

/*CloseTradeMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type CloseTradeMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.CloseTradeMethodNotAllowedBody
}

func (o *CloseTradeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/trades/{tradeSpecifier}/close][%d] closeTradeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CloseTradeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.CloseTradeMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}