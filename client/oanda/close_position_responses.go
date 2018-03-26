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

// ClosePositionReader is a Reader for the ClosePosition structure.
type ClosePositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClosePositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClosePositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewClosePositionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewClosePositionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewClosePositionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewClosePositionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClosePositionOK creates a ClosePositionOK with default headers values
func NewClosePositionOK() *ClosePositionOK {
	return &ClosePositionOK{}
}

/*ClosePositionOK handles this case with default header values.

The Position closeout request has been successfully processed.
*/
type ClosePositionOK struct {
	/*A link to the Position that was just closed out
	 */
	Location string
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ClosePositionOKBody
}

func (o *ClosePositionOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionOK  %+v", 200, o.Payload)
}

func (o *ClosePositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ClosePositionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionBadRequest creates a ClosePositionBadRequest with default headers values
func NewClosePositionBadRequest() *ClosePositionBadRequest {
	return &ClosePositionBadRequest{}
}

/*ClosePositionBadRequest handles this case with default header values.

The Parameters provided that describe the Position closeout are invalid.
*/
type ClosePositionBadRequest struct {
	Payload *models.ClosePositionBadRequestBody
}

func (o *ClosePositionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionBadRequest  %+v", 400, o.Payload)
}

func (o *ClosePositionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClosePositionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionUnauthorized creates a ClosePositionUnauthorized with default headers values
func NewClosePositionUnauthorized() *ClosePositionUnauthorized {
	return &ClosePositionUnauthorized{}
}

/*ClosePositionUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ClosePositionUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ClosePositionUnauthorizedBody
}

func (o *ClosePositionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionUnauthorized  %+v", 401, o.Payload)
}

func (o *ClosePositionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ClosePositionUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionNotFound creates a ClosePositionNotFound with default headers values
func NewClosePositionNotFound() *ClosePositionNotFound {
	return &ClosePositionNotFound{}
}

/*ClosePositionNotFound handles this case with default header values.

The Account or one or more of the Positions specified does not exist.
*/
type ClosePositionNotFound struct {
	Payload *models.ClosePositionNotFoundBody
}

func (o *ClosePositionNotFound) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionNotFound  %+v", 404, o.Payload)
}

func (o *ClosePositionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClosePositionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClosePositionMethodNotAllowed creates a ClosePositionMethodNotAllowed with default headers values
func NewClosePositionMethodNotAllowed() *ClosePositionMethodNotAllowed {
	return &ClosePositionMethodNotAllowed{}
}

/*ClosePositionMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ClosePositionMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ClosePositionMethodNotAllowedBody
}

func (o *ClosePositionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /accounts/{accountID}/positions/{instrument}/close][%d] closePositionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ClosePositionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ClosePositionMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
