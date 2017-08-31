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

// CreateCompartmentReader is a Reader for the CreateCompartment structure.
type CreateCompartmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCompartmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCompartmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateCompartmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCompartmentOK creates a CreateCompartmentOK with default headers values
func NewCreateCompartmentOK() *CreateCompartmentOK {
	return &CreateCompartmentOK{}
}

/*CreateCompartmentOK handles this case with default header values.

The compartment is being created.
*/
type CreateCompartmentOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Compartment
}

func (o *CreateCompartmentOK) Error() string {
	return fmt.Sprintf("[POST /compartments/][%d] createCompartmentOK  %+v", 200, o.Payload)
}

func (o *CreateCompartmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Compartment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCompartmentDefault creates a CreateCompartmentDefault with default headers values
func NewCreateCompartmentDefault(code int) *CreateCompartmentDefault {
	return &CreateCompartmentDefault{
		_statusCode: code,
	}
}

/*CreateCompartmentDefault handles this case with default header values.

An error has occurred.

*/
type CreateCompartmentDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create compartment default response
func (o *CreateCompartmentDefault) Code() int {
	return o._statusCode
}

func (o *CreateCompartmentDefault) Error() string {
	return fmt.Sprintf("[POST /compartments/][%d] CreateCompartment default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCompartmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
