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

// GetPolicyReader is a Reader for the GetPolicy structure.
type GetPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPolicyOK creates a GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {
	return &GetPolicyOK{}
}

/*GetPolicyOK handles this case with default header values.

The policy was retrieved.
*/
type GetPolicyOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Policy
}

func (o *GetPolicyOK) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] getPolicyOK  %+v", 200, o.Payload)
}

func (o *GetPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyBadRequest creates a GetPolicyBadRequest with default headers values
func NewGetPolicyBadRequest() *GetPolicyBadRequest {
	return &GetPolicyBadRequest{}
}

/*GetPolicyBadRequest handles this case with default header values.

Bad Request
*/
type GetPolicyBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] getPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyForbidden creates a GetPolicyForbidden with default headers values
func NewGetPolicyForbidden() *GetPolicyForbidden {
	return &GetPolicyForbidden{}
}

/*GetPolicyForbidden handles this case with default header values.

Forbidden
*/
type GetPolicyForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] getPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyNotFound creates a GetPolicyNotFound with default headers values
func NewGetPolicyNotFound() *GetPolicyNotFound {
	return &GetPolicyNotFound{}
}

/*GetPolicyNotFound handles this case with default header values.

Not Found
*/
type GetPolicyNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] getPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyInternalServerError creates a GetPolicyInternalServerError with default headers values
func NewGetPolicyInternalServerError() *GetPolicyInternalServerError {
	return &GetPolicyInternalServerError{}
}

/*GetPolicyInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPolicyInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] getPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyDefault creates a GetPolicyDefault with default headers values
func NewGetPolicyDefault(code int) *GetPolicyDefault {
	return &GetPolicyDefault{
		_statusCode: code,
	}
}

/*GetPolicyDefault handles this case with default header values.

An error has occurred.

*/
type GetPolicyDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get policy default response
func (o *GetPolicyDefault) Code() int {
	return o._statusCode
}

func (o *GetPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /policies/{policyId}][%d] GetPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
