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

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

The user was retrieved.
*/
type GetUserOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.User
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserBadRequest creates a GetUserBadRequest with default headers values
func NewGetUserBadRequest() *GetUserBadRequest {
	return &GetUserBadRequest{}
}

/*GetUserBadRequest handles this case with default header values.

Bad Request
*/
type GetUserBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] getUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*GetUserForbidden handles this case with default header values.

Forbidden
*/
type GetUserForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*GetUserNotFound handles this case with default header values.

Not Found
*/
type GetUserNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] getUserNotFound  %+v", 404, o.Payload)
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserInternalServerError creates a GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {
	return &GetUserInternalServerError{}
}

/*GetUserInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetUserInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] getUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDefault creates a GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	return &GetUserDefault{
		_statusCode: code,
	}
}

/*GetUserDefault handles this case with default header values.

An error has occurred.

*/
type GetUserDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get user default response
func (o *GetUserDefault) Code() int {
	return o._statusCode
}

func (o *GetUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userId}][%d] GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
