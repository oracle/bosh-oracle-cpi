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

// ListGroupsReader is a Reader for the ListGroups structure.
type ListGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGroupsOK creates a ListGroupsOK with default headers values
func NewListGroupsOK() *ListGroupsOK {
	return &ListGroupsOK{}
}

/*ListGroupsOK handles this case with default header values.

The list is being retrieved.
*/
type ListGroupsOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.Group
}

func (o *ListGroupsOK) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] listGroupsOK  %+v", 200, o.Payload)
}

func (o *ListGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListGroupsBadRequest creates a ListGroupsBadRequest with default headers values
func NewListGroupsBadRequest() *ListGroupsBadRequest {
	return &ListGroupsBadRequest{}
}

/*ListGroupsBadRequest handles this case with default header values.

Bad Request
*/
type ListGroupsBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] listGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *ListGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupsForbidden creates a ListGroupsForbidden with default headers values
func NewListGroupsForbidden() *ListGroupsForbidden {
	return &ListGroupsForbidden{}
}

/*ListGroupsForbidden handles this case with default header values.

Forbidden
*/
type ListGroupsForbidden struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] listGroupsForbidden  %+v", 403, o.Payload)
}

func (o *ListGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupsNotFound creates a ListGroupsNotFound with default headers values
func NewListGroupsNotFound() *ListGroupsNotFound {
	return &ListGroupsNotFound{}
}

/*ListGroupsNotFound handles this case with default header values.

Not Found
*/
type ListGroupsNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] listGroupsNotFound  %+v", 404, o.Payload)
}

func (o *ListGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupsInternalServerError creates a ListGroupsInternalServerError with default headers values
func NewListGroupsInternalServerError() *ListGroupsInternalServerError {
	return &ListGroupsInternalServerError{}
}

/*ListGroupsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListGroupsInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] listGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupsDefault creates a ListGroupsDefault with default headers values
func NewListGroupsDefault(code int) *ListGroupsDefault {
	return &ListGroupsDefault{
		_statusCode: code,
	}
}

/*ListGroupsDefault handles this case with default header values.

An error has occurred.

*/
type ListGroupsDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list groups default response
func (o *ListGroupsDefault) Code() int {
	return o._statusCode
}

func (o *ListGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /groups/][%d] ListGroups default  %+v", o._statusCode, o.Payload)
}

func (o *ListGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
