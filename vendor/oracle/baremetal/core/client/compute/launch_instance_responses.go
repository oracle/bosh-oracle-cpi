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

// LaunchInstanceReader is a Reader for the LaunchInstance structure.
type LaunchInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LaunchInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLaunchInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLaunchInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewLaunchInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLaunchInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewLaunchInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLaunchInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewLaunchInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLaunchInstanceOK creates a LaunchInstanceOK with default headers values
func NewLaunchInstanceOK() *LaunchInstanceOK {
	return &LaunchInstanceOK{}
}

/*LaunchInstanceOK handles this case with default header values.

The instance is being launched.
*/
type LaunchInstanceOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Instance
}

func (o *LaunchInstanceOK) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceOK  %+v", 200, o.Payload)
}

func (o *LaunchInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLaunchInstanceBadRequest creates a LaunchInstanceBadRequest with default headers values
func NewLaunchInstanceBadRequest() *LaunchInstanceBadRequest {
	return &LaunchInstanceBadRequest{}
}

/*LaunchInstanceBadRequest handles this case with default header values.

Bad Request
*/
type LaunchInstanceBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *LaunchInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *LaunchInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInstanceUnauthorized creates a LaunchInstanceUnauthorized with default headers values
func NewLaunchInstanceUnauthorized() *LaunchInstanceUnauthorized {
	return &LaunchInstanceUnauthorized{}
}

/*LaunchInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type LaunchInstanceUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *LaunchInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *LaunchInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInstanceNotFound creates a LaunchInstanceNotFound with default headers values
func NewLaunchInstanceNotFound() *LaunchInstanceNotFound {
	return &LaunchInstanceNotFound{}
}

/*LaunchInstanceNotFound handles this case with default header values.

Not Found
*/
type LaunchInstanceNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *LaunchInstanceNotFound) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceNotFound  %+v", 404, o.Payload)
}

func (o *LaunchInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInstanceConflict creates a LaunchInstanceConflict with default headers values
func NewLaunchInstanceConflict() *LaunchInstanceConflict {
	return &LaunchInstanceConflict{}
}

/*LaunchInstanceConflict handles this case with default header values.

Conflict
*/
type LaunchInstanceConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *LaunchInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceConflict  %+v", 409, o.Payload)
}

func (o *LaunchInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInstanceInternalServerError creates a LaunchInstanceInternalServerError with default headers values
func NewLaunchInstanceInternalServerError() *LaunchInstanceInternalServerError {
	return &LaunchInstanceInternalServerError{}
}

/*LaunchInstanceInternalServerError handles this case with default header values.

Internal Server Error
*/
type LaunchInstanceInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *LaunchInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] launchInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *LaunchInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchInstanceDefault creates a LaunchInstanceDefault with default headers values
func NewLaunchInstanceDefault(code int) *LaunchInstanceDefault {
	return &LaunchInstanceDefault{
		_statusCode: code,
	}
}

/*LaunchInstanceDefault handles this case with default header values.

An error has occurred.
*/
type LaunchInstanceDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the launch instance default response
func (o *LaunchInstanceDefault) Code() int {
	return o._statusCode
}

func (o *LaunchInstanceDefault) Error() string {
	return fmt.Sprintf("[POST /instances/][%d] LaunchInstance default  %+v", o._statusCode, o.Payload)
}

func (o *LaunchInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
