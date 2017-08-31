package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// UpdateInstanceReader is a Reader for the UpdateInstance structure.
type UpdateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateInstancePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateInstanceOK creates a UpdateInstanceOK with default headers values
func NewUpdateInstanceOK() *UpdateInstanceOK {
	return &UpdateInstanceOK{}
}

/*UpdateInstanceOK handles this case with default header values.

The instance was updated.
*/
type UpdateInstanceOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Instance
}

func (o *UpdateInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstanceOK  %+v", 200, o.Payload)
}

func (o *UpdateInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceUnauthorized creates a UpdateInstanceUnauthorized with default headers values
func NewUpdateInstanceUnauthorized() *UpdateInstanceUnauthorized {
	return &UpdateInstanceUnauthorized{}
}

/*UpdateInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateInstanceUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceNotFound creates a UpdateInstanceNotFound with default headers values
func NewUpdateInstanceNotFound() *UpdateInstanceNotFound {
	return &UpdateInstanceNotFound{}
}

/*UpdateInstanceNotFound handles this case with default header values.

Not Found
*/
type UpdateInstanceNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateInstanceNotFound) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstanceNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceConflict creates a UpdateInstanceConflict with default headers values
func NewUpdateInstanceConflict() *UpdateInstanceConflict {
	return &UpdateInstanceConflict{}
}

/*UpdateInstanceConflict handles this case with default header values.

Conflict
*/
type UpdateInstanceConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateInstanceConflict) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstanceConflict  %+v", 409, o.Payload)
}

func (o *UpdateInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstancePreconditionFailed creates a UpdateInstancePreconditionFailed with default headers values
func NewUpdateInstancePreconditionFailed() *UpdateInstancePreconditionFailed {
	return &UpdateInstancePreconditionFailed{}
}

/*UpdateInstancePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateInstancePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateInstancePreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstancePreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateInstancePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceInternalServerError creates a UpdateInstanceInternalServerError with default headers values
func NewUpdateInstanceInternalServerError() *UpdateInstanceInternalServerError {
	return &UpdateInstanceInternalServerError{}
}

/*UpdateInstanceInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateInstanceInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] updateInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstanceDefault creates a UpdateInstanceDefault with default headers values
func NewUpdateInstanceDefault(code int) *UpdateInstanceDefault {
	return &UpdateInstanceDefault{
		_statusCode: code,
	}
}

/*UpdateInstanceDefault handles this case with default header values.

An error has occurred.
*/
type UpdateInstanceDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update instance default response
func (o *UpdateInstanceDefault) Code() int {
	return o._statusCode
}

func (o *UpdateInstanceDefault) Error() string {
	return fmt.Sprintf("[PUT /instances/{instanceId}][%d] UpdateInstance default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
