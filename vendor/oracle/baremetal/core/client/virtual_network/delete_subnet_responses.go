package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/baremetal/core/models"
)

// DeleteSubnetReader is a Reader for the DeleteSubnet structure.
type DeleteSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSubnetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteSubnetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteSubnetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSubnetNoContent creates a DeleteSubnetNoContent with default headers values
func NewDeleteSubnetNoContent() *DeleteSubnetNoContent {
	return &DeleteSubnetNoContent{}
}

/*DeleteSubnetNoContent handles this case with default header values.

The subnet is being deleted.
*/
type DeleteSubnetNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteSubnetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] deleteSubnetNoContent ", 204)
}

func (o *DeleteSubnetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteSubnetUnauthorized creates a DeleteSubnetUnauthorized with default headers values
func NewDeleteSubnetUnauthorized() *DeleteSubnetUnauthorized {
	return &DeleteSubnetUnauthorized{}
}

/*DeleteSubnetUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteSubnetUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteSubnetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] deleteSubnetUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSubnetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetNotFound creates a DeleteSubnetNotFound with default headers values
func NewDeleteSubnetNotFound() *DeleteSubnetNotFound {
	return &DeleteSubnetNotFound{}
}

/*DeleteSubnetNotFound handles this case with default header values.

Not Found
*/
type DeleteSubnetNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteSubnetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] deleteSubnetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetPreconditionFailed creates a DeleteSubnetPreconditionFailed with default headers values
func NewDeleteSubnetPreconditionFailed() *DeleteSubnetPreconditionFailed {
	return &DeleteSubnetPreconditionFailed{}
}

/*DeleteSubnetPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteSubnetPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteSubnetPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] deleteSubnetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteSubnetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetInternalServerError creates a DeleteSubnetInternalServerError with default headers values
func NewDeleteSubnetInternalServerError() *DeleteSubnetInternalServerError {
	return &DeleteSubnetInternalServerError{}
}

/*DeleteSubnetInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteSubnetInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] deleteSubnetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSubnetDefault creates a DeleteSubnetDefault with default headers values
func NewDeleteSubnetDefault(code int) *DeleteSubnetDefault {
	return &DeleteSubnetDefault{
		_statusCode: code,
	}
}

/*DeleteSubnetDefault handles this case with default header values.

An error has occurred.
*/
type DeleteSubnetDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete subnet default response
func (o *DeleteSubnetDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSubnetDefault) Error() string {
	return fmt.Sprintf("[DELETE /subnets/{subnetId}][%d] DeleteSubnet default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
