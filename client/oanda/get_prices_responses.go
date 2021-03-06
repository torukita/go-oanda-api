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

// GetPricesReader is a Reader for the GetPrices structure.
type GetPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPricesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPricesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPricesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGetPricesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPricesOK creates a GetPricesOK with default headers values
func NewGetPricesOK() *GetPricesOK {
	return &GetPricesOK{}
}

/*GetPricesOK handles this case with default header values.

Pricing information has been successfully provided.
*/
type GetPricesOK struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPricesOKBody
}

func (o *GetPricesOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing][%d] getPricesOK  %+v", 200, o.Payload)
}

func (o *GetPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPricesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricesBadRequest creates a GetPricesBadRequest with default headers values
func NewGetPricesBadRequest() *GetPricesBadRequest {
	return &GetPricesBadRequest{}
}

/*GetPricesBadRequest handles this case with default header values.

Bad Request. The client has provided invalid data to be processed by the server.
*/
type GetPricesBadRequest struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPricesBadRequestBody
}

func (o *GetPricesBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing][%d] getPricesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPricesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPricesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricesUnauthorized creates a GetPricesUnauthorized with default headers values
func NewGetPricesUnauthorized() *GetPricesUnauthorized {
	return &GetPricesUnauthorized{}
}

/*GetPricesUnauthorized handles this case with default header values.

Unauthorized. The endpoint being access required the client to authenticated, however the the authentication token is invalid or has not been provided.
*/
type GetPricesUnauthorized struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPricesUnauthorizedBody
}

func (o *GetPricesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing][%d] getPricesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPricesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPricesUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricesNotFound creates a GetPricesNotFound with default headers values
func NewGetPricesNotFound() *GetPricesNotFound {
	return &GetPricesNotFound{}
}

/*GetPricesNotFound handles this case with default header values.

Not Found. The client has attempted to access an entity that does not exist.
*/
type GetPricesNotFound struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPricesNotFoundBody
}

func (o *GetPricesNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing][%d] getPricesNotFound  %+v", 404, o.Payload)
}

func (o *GetPricesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPricesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPricesMethodNotAllowed creates a GetPricesMethodNotAllowed with default headers values
func NewGetPricesMethodNotAllowed() *GetPricesMethodNotAllowed {
	return &GetPricesMethodNotAllowed{}
}

/*GetPricesMethodNotAllowed handles this case with default header values.

Method Not Allowed. The client has attempted to access an endpoint using an HTTP method that is not supported.
*/
type GetPricesMethodNotAllowed struct {
	/*The unique identifier generated for the request
	 */
	RequestID string

	Payload *models.GetPricesMethodNotAllowedBody
}

func (o *GetPricesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountID}/pricing][%d] getPricesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetPricesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RequestID
	o.RequestID = response.GetHeader("RequestID")

	o.Payload = new(models.GetPricesMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
