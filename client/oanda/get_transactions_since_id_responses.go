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

// GetTransactionsSinceIDReader is a Reader for the GetTransactionsSinceID structure.
type GetTransactionsSinceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionsSinceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionsSinceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionsSinceIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTransactionsSinceIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTransactionsSinceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetTransactionsSinceIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 416:
		result := NewGetTransactionsSinceIDRequestRangeNotSatisfiable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionsSinceIDOK creates a GetTransactionsSinceIDOK with default headers values
func NewGetTransactionsSinceIDOK() *GetTransactionsSinceIDOK {
	return &GetTransactionsSinceIDOK{}
}

/*GetTransactionsSinceIDOK handles this case with default header values.

The requested time range of Transactions are provided.
*/
type GetTransactionsSinceIDOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDOKBody
}

func (o *GetTransactionsSinceIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdOK  %+v", 200, o.Payload)
}

func (o *GetTransactionsSinceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsSinceIDBadRequest creates a GetTransactionsSinceIDBadRequest with default headers values
func NewGetTransactionsSinceIDBadRequest() *GetTransactionsSinceIDBadRequest {
	return &GetTransactionsSinceIDBadRequest{}
}

/*GetTransactionsSinceIDBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetTransactionsSinceIDBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDBadRequestBody
}

func (o *GetTransactionsSinceIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionsSinceIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsSinceIDUnauthorized creates a GetTransactionsSinceIDUnauthorized with default headers values
func NewGetTransactionsSinceIDUnauthorized() *GetTransactionsSinceIDUnauthorized {
	return &GetTransactionsSinceIDUnauthorized{}
}

/*GetTransactionsSinceIDUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetTransactionsSinceIDUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDUnauthorizedBody
}

func (o *GetTransactionsSinceIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionsSinceIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsSinceIDNotFound creates a GetTransactionsSinceIDNotFound with default headers values
func NewGetTransactionsSinceIDNotFound() *GetTransactionsSinceIDNotFound {
	return &GetTransactionsSinceIDNotFound{}
}

/*GetTransactionsSinceIDNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetTransactionsSinceIDNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDNotFoundBody
}

func (o *GetTransactionsSinceIDNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionsSinceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsSinceIDMethodNotAllowed creates a GetTransactionsSinceIDMethodNotAllowed with default headers values
func NewGetTransactionsSinceIDMethodNotAllowed() *GetTransactionsSinceIDMethodNotAllowed {
	return &GetTransactionsSinceIDMethodNotAllowed{}
}

/*GetTransactionsSinceIDMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetTransactionsSinceIDMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDMethodNotAllowedBody
}

func (o *GetTransactionsSinceIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetTransactionsSinceIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsSinceIDRequestRangeNotSatisfiable creates a GetTransactionsSinceIDRequestRangeNotSatisfiable with default headers values
func NewGetTransactionsSinceIDRequestRangeNotSatisfiable() *GetTransactionsSinceIDRequestRangeNotSatisfiable {
	return &GetTransactionsSinceIDRequestRangeNotSatisfiable{}
}

/*GetTransactionsSinceIDRequestRangeNotSatisfiable handles this case with default header values.

Range Not Satisfiable. The client has specified a range that is invalid or cannot be processed.
*/
type GetTransactionsSinceIDRequestRangeNotSatisfiable struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetTransactionsSinceIDRequestedRangeNotSatisfiableBody
}

func (o *GetTransactionsSinceIDRequestRangeNotSatisfiable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/transactions/sinceid][%d] getTransactionsSinceIdRequestRangeNotSatisfiable  %+v", 416, o.Payload)
}

func (o *GetTransactionsSinceIDRequestRangeNotSatisfiable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetTransactionsSinceIDRequestedRangeNotSatisfiableBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}