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

// UpdateDhcpOptionsReader is a Reader for the UpdateDhcpOptions structure.
type UpdateDhcpOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDhcpOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDhcpOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateDhcpOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateDhcpOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateDhcpOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateDhcpOptionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateDhcpOptionsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateDhcpOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateDhcpOptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDhcpOptionsOK creates a UpdateDhcpOptionsOK with default headers values
func NewUpdateDhcpOptionsOK() *UpdateDhcpOptionsOK {
	return &UpdateDhcpOptionsOK{}
}

/*UpdateDhcpOptionsOK handles this case with default header values.

The set of DHCP options was updated.
*/
type UpdateDhcpOptionsOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.DhcpOptions
}

func (o *UpdateDhcpOptionsOK) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsOK  %+v", 200, o.Payload)
}

func (o *UpdateDhcpOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.DhcpOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsBadRequest creates a UpdateDhcpOptionsBadRequest with default headers values
func NewUpdateDhcpOptionsBadRequest() *UpdateDhcpOptionsBadRequest {
	return &UpdateDhcpOptionsBadRequest{}
}

/*UpdateDhcpOptionsBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDhcpOptionsBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDhcpOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsUnauthorized creates a UpdateDhcpOptionsUnauthorized with default headers values
func NewUpdateDhcpOptionsUnauthorized() *UpdateDhcpOptionsUnauthorized {
	return &UpdateDhcpOptionsUnauthorized{}
}

/*UpdateDhcpOptionsUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDhcpOptionsUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDhcpOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsNotFound creates a UpdateDhcpOptionsNotFound with default headers values
func NewUpdateDhcpOptionsNotFound() *UpdateDhcpOptionsNotFound {
	return &UpdateDhcpOptionsNotFound{}
}

/*UpdateDhcpOptionsNotFound handles this case with default header values.

Not Found
*/
type UpdateDhcpOptionsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDhcpOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsConflict creates a UpdateDhcpOptionsConflict with default headers values
func NewUpdateDhcpOptionsConflict() *UpdateDhcpOptionsConflict {
	return &UpdateDhcpOptionsConflict{}
}

/*UpdateDhcpOptionsConflict handles this case with default header values.

Conflict
*/
type UpdateDhcpOptionsConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsConflict) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsConflict  %+v", 409, o.Payload)
}

func (o *UpdateDhcpOptionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsPreconditionFailed creates a UpdateDhcpOptionsPreconditionFailed with default headers values
func NewUpdateDhcpOptionsPreconditionFailed() *UpdateDhcpOptionsPreconditionFailed {
	return &UpdateDhcpOptionsPreconditionFailed{}
}

/*UpdateDhcpOptionsPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateDhcpOptionsPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateDhcpOptionsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsInternalServerError creates a UpdateDhcpOptionsInternalServerError with default headers values
func NewUpdateDhcpOptionsInternalServerError() *UpdateDhcpOptionsInternalServerError {
	return &UpdateDhcpOptionsInternalServerError{}
}

/*UpdateDhcpOptionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateDhcpOptionsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateDhcpOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] updateDhcpOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDhcpOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDhcpOptionsDefault creates a UpdateDhcpOptionsDefault with default headers values
func NewUpdateDhcpOptionsDefault(code int) *UpdateDhcpOptionsDefault {
	return &UpdateDhcpOptionsDefault{
		_statusCode: code,
	}
}

/*UpdateDhcpOptionsDefault handles this case with default header values.

An error has occurred.
*/
type UpdateDhcpOptionsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update dhcp options default response
func (o *UpdateDhcpOptionsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDhcpOptionsDefault) Error() string {
	return fmt.Sprintf("[PUT /dhcps/{dhcpId}][%d] UpdateDhcpOptions default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDhcpOptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
