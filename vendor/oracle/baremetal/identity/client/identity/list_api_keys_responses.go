package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/identity/models"
)

// ListAPIKeysReader is a Reader for the ListAPIKeys structure.
type ListAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAPIKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListAPIKeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAPIKeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAPIKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAPIKeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListAPIKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAPIKeysOK creates a ListAPIKeysOK with default headers values
func NewListAPIKeysOK() *ListAPIKeysOK {
	return &ListAPIKeysOK{}
}

/*ListAPIKeysOK handles this case with default header values.

The list is being retrieved.
*/
type ListAPIKeysOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.APIKey
}

func (o *ListAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] listApiKeysOK  %+v", 200, o.Payload)
}

func (o *ListAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeysBadRequest creates a ListAPIKeysBadRequest with default headers values
func NewListAPIKeysBadRequest() *ListAPIKeysBadRequest {
	return &ListAPIKeysBadRequest{}
}

/*ListAPIKeysBadRequest handles this case with default header values.

Bad Request
*/
type ListAPIKeysBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAPIKeysBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] listApiKeysBadRequest  %+v", 400, o.Payload)
}

func (o *ListAPIKeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeysForbidden creates a ListAPIKeysForbidden with default headers values
func NewListAPIKeysForbidden() *ListAPIKeysForbidden {
	return &ListAPIKeysForbidden{}
}

/*ListAPIKeysForbidden handles this case with default header values.

Forbidden
*/
type ListAPIKeysForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAPIKeysForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] listApiKeysForbidden  %+v", 403, o.Payload)
}

func (o *ListAPIKeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeysNotFound creates a ListAPIKeysNotFound with default headers values
func NewListAPIKeysNotFound() *ListAPIKeysNotFound {
	return &ListAPIKeysNotFound{}
}

/*ListAPIKeysNotFound handles this case with default header values.

Not Found
*/
type ListAPIKeysNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAPIKeysNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] listApiKeysNotFound  %+v", 404, o.Payload)
}

func (o *ListAPIKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeysInternalServerError creates a ListAPIKeysInternalServerError with default headers values
func NewListAPIKeysInternalServerError() *ListAPIKeysInternalServerError {
	return &ListAPIKeysInternalServerError{}
}

/*ListAPIKeysInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAPIKeysInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAPIKeysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] listApiKeysInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAPIKeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeysDefault creates a ListAPIKeysDefault with default headers values
func NewListAPIKeysDefault(code int) *ListAPIKeysDefault {
	return &ListAPIKeysDefault{
		_statusCode: code,
	}
}

/*ListAPIKeysDefault handles this case with default header values.

An error has occurred.

*/
type ListAPIKeysDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list Api keys default response
func (o *ListAPIKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListAPIKeysDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/apiKeys/][%d] ListApiKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListAPIKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
