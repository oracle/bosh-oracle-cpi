package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/identity/models"
)

// ListSwiftPasswordsReader is a Reader for the ListSwiftPasswords structure.
type ListSwiftPasswordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSwiftPasswordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSwiftPasswordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSwiftPasswordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSwiftPasswordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListSwiftPasswordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSwiftPasswordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListSwiftPasswordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSwiftPasswordsOK creates a ListSwiftPasswordsOK with default headers values
func NewListSwiftPasswordsOK() *ListSwiftPasswordsOK {
	return &ListSwiftPasswordsOK{}
}

/*ListSwiftPasswordsOK handles this case with default header values.

The list is being retrieved.
*/
type ListSwiftPasswordsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.SwiftPassword
}

func (o *ListSwiftPasswordsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] listSwiftPasswordsOK  %+v", 200, o.Payload)
}

func (o *ListSwiftPasswordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListSwiftPasswordsBadRequest creates a ListSwiftPasswordsBadRequest with default headers values
func NewListSwiftPasswordsBadRequest() *ListSwiftPasswordsBadRequest {
	return &ListSwiftPasswordsBadRequest{}
}

/*ListSwiftPasswordsBadRequest handles this case with default header values.

Bad Request
*/
type ListSwiftPasswordsBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListSwiftPasswordsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] listSwiftPasswordsBadRequest  %+v", 400, o.Payload)
}

func (o *ListSwiftPasswordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSwiftPasswordsForbidden creates a ListSwiftPasswordsForbidden with default headers values
func NewListSwiftPasswordsForbidden() *ListSwiftPasswordsForbidden {
	return &ListSwiftPasswordsForbidden{}
}

/*ListSwiftPasswordsForbidden handles this case with default header values.

Forbidden
*/
type ListSwiftPasswordsForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListSwiftPasswordsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] listSwiftPasswordsForbidden  %+v", 403, o.Payload)
}

func (o *ListSwiftPasswordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSwiftPasswordsNotFound creates a ListSwiftPasswordsNotFound with default headers values
func NewListSwiftPasswordsNotFound() *ListSwiftPasswordsNotFound {
	return &ListSwiftPasswordsNotFound{}
}

/*ListSwiftPasswordsNotFound handles this case with default header values.

Not Found
*/
type ListSwiftPasswordsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListSwiftPasswordsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] listSwiftPasswordsNotFound  %+v", 404, o.Payload)
}

func (o *ListSwiftPasswordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSwiftPasswordsInternalServerError creates a ListSwiftPasswordsInternalServerError with default headers values
func NewListSwiftPasswordsInternalServerError() *ListSwiftPasswordsInternalServerError {
	return &ListSwiftPasswordsInternalServerError{}
}

/*ListSwiftPasswordsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListSwiftPasswordsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListSwiftPasswordsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] listSwiftPasswordsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSwiftPasswordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSwiftPasswordsDefault creates a ListSwiftPasswordsDefault with default headers values
func NewListSwiftPasswordsDefault(code int) *ListSwiftPasswordsDefault {
	return &ListSwiftPasswordsDefault{
		_statusCode: code,
	}
}

/*ListSwiftPasswordsDefault handles this case with default header values.

An error has occurred.

*/
type ListSwiftPasswordsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list swift passwords default response
func (o *ListSwiftPasswordsDefault) Code() int {
	return o._statusCode
}

func (o *ListSwiftPasswordsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/swiftPasswords/][%d] ListSwiftPasswords default  %+v", o._statusCode, o.Payload)
}

func (o *ListSwiftPasswordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
