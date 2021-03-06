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

// GetPositionReader is a Reader for the GetPosition structure.
type GetPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPositionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPositionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetPositionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPositionOK creates a GetPositionOK with default headers values
func NewGetPositionOK() *GetPositionOK {
	return &GetPositionOK{}
}

/*GetPositionOK handles this case with default header values.

The Position is provided.
*/
type GetPositionOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPositionOKBody
}

func (o *GetPositionOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions/{instrument}][%d] getPositionOK  %+v", 200, o.Payload)
}

func (o *GetPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPositionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPositionUnauthorized creates a GetPositionUnauthorized with default headers values
func NewGetPositionUnauthorized() *GetPositionUnauthorized {
	return &GetPositionUnauthorized{}
}

/*GetPositionUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetPositionUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPositionUnauthorizedBody
}

func (o *GetPositionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions/{instrument}][%d] getPositionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPositionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPositionUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPositionNotFound creates a GetPositionNotFound with default headers values
func NewGetPositionNotFound() *GetPositionNotFound {
	return &GetPositionNotFound{}
}

/*GetPositionNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetPositionNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPositionNotFoundBody
}

func (o *GetPositionNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions/{instrument}][%d] getPositionNotFound  %+v", 404, o.Payload)
}

func (o *GetPositionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPositionNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPositionMethodNotAllowed creates a GetPositionMethodNotAllowed with default headers values
func NewGetPositionMethodNotAllowed() *GetPositionMethodNotAllowed {
	return &GetPositionMethodNotAllowed{}
}

/*GetPositionMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetPositionMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPositionMethodNotAllowedBody
}

func (o *GetPositionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/positions/{instrument}][%d] getPositionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetPositionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPositionMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
