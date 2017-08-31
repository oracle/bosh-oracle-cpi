package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSwiftPasswordParams creates a new DeleteSwiftPasswordParams object
// with the default values initialized.
func NewDeleteSwiftPasswordParams() *DeleteSwiftPasswordParams {
	var ()
	return &DeleteSwiftPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSwiftPasswordParamsWithTimeout creates a new DeleteSwiftPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSwiftPasswordParamsWithTimeout(timeout time.Duration) *DeleteSwiftPasswordParams {
	var ()
	return &DeleteSwiftPasswordParams{

		timeout: timeout,
	}
}

// NewDeleteSwiftPasswordParamsWithContext creates a new DeleteSwiftPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSwiftPasswordParamsWithContext(ctx context.Context) *DeleteSwiftPasswordParams {
	var ()
	return &DeleteSwiftPasswordParams{

		Context: ctx,
	}
}

// NewDeleteSwiftPasswordParamsWithHTTPClient creates a new DeleteSwiftPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSwiftPasswordParamsWithHTTPClient(client *http.Client) *DeleteSwiftPasswordParams {
	var ()
	return &DeleteSwiftPasswordParams{
		HTTPClient: client,
	}
}

/*DeleteSwiftPasswordParams contains all the parameters to send to the API endpoint
for the delete swift password operation typically these are written to a http.Request
*/
type DeleteSwiftPasswordParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*SwiftPasswordID
	  The OCID of the Swift password.

	*/
	SwiftPasswordID string
	/*UserID
	  The OCID of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithTimeout(timeout time.Duration) *DeleteSwiftPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithContext(ctx context.Context) *DeleteSwiftPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithHTTPClient(client *http.Client) *DeleteSwiftPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithIfMatch(ifMatch *string) *DeleteSwiftPasswordParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithSwiftPasswordID adds the swiftPasswordID to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithSwiftPasswordID(swiftPasswordID string) *DeleteSwiftPasswordParams {
	o.SetSwiftPasswordID(swiftPasswordID)
	return o
}

// SetSwiftPasswordID adds the swiftPasswordId to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetSwiftPasswordID(swiftPasswordID string) {
	o.SwiftPasswordID = swiftPasswordID
}

// WithUserID adds the userID to the delete swift password params
func (o *DeleteSwiftPasswordParams) WithUserID(userID string) *DeleteSwiftPasswordParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete swift password params
func (o *DeleteSwiftPasswordParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSwiftPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param swiftPasswordId
	if err := r.SetPathParam("swiftPasswordId", o.SwiftPasswordID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
