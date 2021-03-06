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

// GetAccountChangesReader is a Reader for the GetAccountChanges structure.
type GetAccountChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountChangesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountChangesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetAccountChangesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 416:
		result := NewGetAccountChangesRequestRangeNotSatisfiable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountChangesOK creates a GetAccountChangesOK with default headers values
func NewGetAccountChangesOK() *GetAccountChangesOK {
	return &GetAccountChangesOK{}
}

/*GetAccountChangesOK handles this case with default header values.

The Account state and changes are provided.
*/
type GetAccountChangesOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetAccountChangesOKBody
}

func (o *GetAccountChangesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/changes][%d] getAccountChangesOK  %+v", 200, o.Payload)
}

func (o *GetAccountChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetAccountChangesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountChangesUnauthorized creates a GetAccountChangesUnauthorized with default headers values
func NewGetAccountChangesUnauthorized() *GetAccountChangesUnauthorized {
	return &GetAccountChangesUnauthorized{}
}

/*GetAccountChangesUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetAccountChangesUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetAccountChangesUnauthorizedBody
}

func (o *GetAccountChangesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/changes][%d] getAccountChangesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountChangesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetAccountChangesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountChangesNotFound creates a GetAccountChangesNotFound with default headers values
func NewGetAccountChangesNotFound() *GetAccountChangesNotFound {
	return &GetAccountChangesNotFound{}
}

/*GetAccountChangesNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetAccountChangesNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetAccountChangesNotFoundBody
}

func (o *GetAccountChangesNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/changes][%d] getAccountChangesNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountChangesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetAccountChangesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountChangesMethodNotAllowed creates a GetAccountChangesMethodNotAllowed with default headers values
func NewGetAccountChangesMethodNotAllowed() *GetAccountChangesMethodNotAllowed {
	return &GetAccountChangesMethodNotAllowed{}
}

/*GetAccountChangesMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetAccountChangesMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetAccountChangesMethodNotAllowedBody
}

func (o *GetAccountChangesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/changes][%d] getAccountChangesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetAccountChangesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetAccountChangesMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountChangesRequestRangeNotSatisfiable creates a GetAccountChangesRequestRangeNotSatisfiable with default headers values
func NewGetAccountChangesRequestRangeNotSatisfiable() *GetAccountChangesRequestRangeNotSatisfiable {
	return &GetAccountChangesRequestRangeNotSatisfiable{}
}

/*GetAccountChangesRequestRangeNotSatisfiable handles this case with default header values.

Range Not Satisfiable. The client has specified a range that is invalid or cannot be processed.
*/
type GetAccountChangesRequestRangeNotSatisfiable struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetAccountChangesRequestedRangeNotSatisfiableBody
}

func (o *GetAccountChangesRequestRangeNotSatisfiable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/changes][%d] getAccountChangesRequestRangeNotSatisfiable  %+v", 416, o.Payload)
}

func (o *GetAccountChangesRequestRangeNotSatisfiable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetAccountChangesRequestedRangeNotSatisfiableBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
