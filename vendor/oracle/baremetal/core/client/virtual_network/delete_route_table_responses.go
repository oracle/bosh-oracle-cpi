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

// DeleteRouteTableReader is a Reader for the DeleteRouteTable structure.
type DeleteRouteTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRouteTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRouteTableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteRouteTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteRouteTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteRouteTablePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteRouteTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteRouteTableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRouteTableNoContent creates a DeleteRouteTableNoContent with default headers values
func NewDeleteRouteTableNoContent() *DeleteRouteTableNoContent {
	return &DeleteRouteTableNoContent{}
}

/*DeleteRouteTableNoContent handles this case with default header values.

The route table is being deleted.
*/
type DeleteRouteTableNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeleteRouteTableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] deleteRouteTableNoContent ", 204)
}

func (o *DeleteRouteTableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeleteRouteTableUnauthorized creates a DeleteRouteTableUnauthorized with default headers values
func NewDeleteRouteTableUnauthorized() *DeleteRouteTableUnauthorized {
	return &DeleteRouteTableUnauthorized{}
}

/*DeleteRouteTableUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRouteTableUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteRouteTableUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] deleteRouteTableUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRouteTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteTableNotFound creates a DeleteRouteTableNotFound with default headers values
func NewDeleteRouteTableNotFound() *DeleteRouteTableNotFound {
	return &DeleteRouteTableNotFound{}
}

/*DeleteRouteTableNotFound handles this case with default header values.

Not Found
*/
type DeleteRouteTableNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteRouteTableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] deleteRouteTableNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRouteTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteTablePreconditionFailed creates a DeleteRouteTablePreconditionFailed with default headers values
func NewDeleteRouteTablePreconditionFailed() *DeleteRouteTablePreconditionFailed {
	return &DeleteRouteTablePreconditionFailed{}
}

/*DeleteRouteTablePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeleteRouteTablePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteRouteTablePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] deleteRouteTablePreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRouteTablePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteTableInternalServerError creates a DeleteRouteTableInternalServerError with default headers values
func NewDeleteRouteTableInternalServerError() *DeleteRouteTableInternalServerError {
	return &DeleteRouteTableInternalServerError{}
}

/*DeleteRouteTableInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteRouteTableInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeleteRouteTableInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] deleteRouteTableInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRouteTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRouteTableDefault creates a DeleteRouteTableDefault with default headers values
func NewDeleteRouteTableDefault(code int) *DeleteRouteTableDefault {
	return &DeleteRouteTableDefault{
		_statusCode: code,
	}
}

/*DeleteRouteTableDefault handles this case with default header values.

An error has occurred.
*/
type DeleteRouteTableDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete route table default response
func (o *DeleteRouteTableDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRouteTableDefault) Error() string {
	return fmt.Sprintf("[DELETE /routeTables/{rtId}][%d] DeleteRouteTable default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRouteTableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
