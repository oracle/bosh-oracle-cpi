package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// UpdateVnicReader is a Reader for the UpdateVnic structure.
type UpdateVnicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVnicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateVnicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateVnicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewUpdateVnicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateVnicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateVnicConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateVnicPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateVnicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateVnicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVnicOK creates a UpdateVnicOK with default headers values
func NewUpdateVnicOK() *UpdateVnicOK {
	return &UpdateVnicOK{}
}

/*UpdateVnicOK handles this case with default header values.

The VNIC was updated.
*/
type UpdateVnicOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Vnic
}

func (o *UpdateVnicOK) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicOK  %+v", 200, o.Payload)
}

func (o *UpdateVnicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Vnic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicBadRequest creates a UpdateVnicBadRequest with default headers values
func NewUpdateVnicBadRequest() *UpdateVnicBadRequest {
	return &UpdateVnicBadRequest{}
}

/*UpdateVnicBadRequest handles this case with default header values.

Bad Request
*/
type UpdateVnicBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicBadRequest) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVnicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicUnauthorized creates a UpdateVnicUnauthorized with default headers values
func NewUpdateVnicUnauthorized() *UpdateVnicUnauthorized {
	return &UpdateVnicUnauthorized{}
}

/*UpdateVnicUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateVnicUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateVnicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicNotFound creates a UpdateVnicNotFound with default headers values
func NewUpdateVnicNotFound() *UpdateVnicNotFound {
	return &UpdateVnicNotFound{}
}

/*UpdateVnicNotFound handles this case with default header values.

Not Found
*/
type UpdateVnicNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicNotFound) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVnicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicConflict creates a UpdateVnicConflict with default headers values
func NewUpdateVnicConflict() *UpdateVnicConflict {
	return &UpdateVnicConflict{}
}

/*UpdateVnicConflict handles this case with default header values.

Conflict
*/
type UpdateVnicConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicConflict) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicConflict  %+v", 409, o.Payload)
}

func (o *UpdateVnicConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicPreconditionFailed creates a UpdateVnicPreconditionFailed with default headers values
func NewUpdateVnicPreconditionFailed() *UpdateVnicPreconditionFailed {
	return &UpdateVnicPreconditionFailed{}
}

/*UpdateVnicPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type UpdateVnicPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateVnicPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicInternalServerError creates a UpdateVnicInternalServerError with default headers values
func NewUpdateVnicInternalServerError() *UpdateVnicInternalServerError {
	return &UpdateVnicInternalServerError{}
}

/*UpdateVnicInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateVnicInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *UpdateVnicInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] updateVnicInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateVnicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVnicDefault creates a UpdateVnicDefault with default headers values
func NewUpdateVnicDefault(code int) *UpdateVnicDefault {
	return &UpdateVnicDefault{
		_statusCode: code,
	}
}

/*UpdateVnicDefault handles this case with default header values.

An error has occurred.
*/
type UpdateVnicDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the update vnic default response
func (o *UpdateVnicDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVnicDefault) Error() string {
	return fmt.Sprintf("[PUT /vnics/{vnicId}][%d] UpdateVnic default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVnicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}