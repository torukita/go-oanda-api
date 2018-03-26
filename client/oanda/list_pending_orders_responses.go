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

// ListPendingOrdersReader is a Reader for the ListPendingOrders structure.
type ListPendingOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPendingOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPendingOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListPendingOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListPendingOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListPendingOrdersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPendingOrdersOK creates a ListPendingOrdersOK with default headers values
func NewListPendingOrdersOK() *ListPendingOrdersOK {
	return &ListPendingOrdersOK{}
}

/*ListPendingOrdersOK handles this case with default header values.

List of pending Orders for the Account
*/
type ListPendingOrdersOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ListPendingOrdersOKBody
}

func (o *ListPendingOrdersOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pendingOrders][%d] listPendingOrdersOK  %+v", 200, o.Payload)
}

func (o *ListPendingOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ListPendingOrdersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPendingOrdersUnauthorized creates a ListPendingOrdersUnauthorized with default headers values
func NewListPendingOrdersUnauthorized() *ListPendingOrdersUnauthorized {
	return &ListPendingOrdersUnauthorized{}
}

/*ListPendingOrdersUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type ListPendingOrdersUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ListPendingOrdersUnauthorizedBody
}

func (o *ListPendingOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pendingOrders][%d] listPendingOrdersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPendingOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ListPendingOrdersUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPendingOrdersNotFound creates a ListPendingOrdersNotFound with default headers values
func NewListPendingOrdersNotFound() *ListPendingOrdersNotFound {
	return &ListPendingOrdersNotFound{}
}

/*ListPendingOrdersNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type ListPendingOrdersNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ListPendingOrdersNotFoundBody
}

func (o *ListPendingOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pendingOrders][%d] listPendingOrdersNotFound  %+v", 404, o.Payload)
}

func (o *ListPendingOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ListPendingOrdersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPendingOrdersMethodNotAllowed creates a ListPendingOrdersMethodNotAllowed with default headers values
func NewListPendingOrdersMethodNotAllowed() *ListPendingOrdersMethodNotAllowed {
	return &ListPendingOrdersMethodNotAllowed{}
}

/*ListPendingOrdersMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type ListPendingOrdersMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.ListPendingOrdersMethodNotAllowedBody
}

func (o *ListPendingOrdersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pendingOrders][%d] listPendingOrdersMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListPendingOrdersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.ListPendingOrdersMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
