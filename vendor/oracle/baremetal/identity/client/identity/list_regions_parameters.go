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

// NewListRegionsParams creates a new ListRegionsParams object
// with the default values initialized.
func NewListRegionsParams() *ListRegionsParams {

	return &ListRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRegionsParamsWithTimeout creates a new ListRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRegionsParamsWithTimeout(timeout time.Duration) *ListRegionsParams {

	return &ListRegionsParams{

		timeout: timeout,
	}
}

// NewListRegionsParamsWithContext creates a new ListRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRegionsParamsWithContext(ctx context.Context) *ListRegionsParams {

	return &ListRegionsParams{

		Context: ctx,
	}
}

// NewListRegionsParamsWithHTTPClient creates a new ListRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRegionsParamsWithHTTPClient(client *http.Client) *ListRegionsParams {

	return &ListRegionsParams{
		HTTPClient: client,
	}
}

/*ListRegionsParams contains all the parameters to send to the API endpoint
for the list regions operation typically these are written to a http.Request
*/
type ListRegionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list regions params
func (o *ListRegionsParams) WithTimeout(timeout time.Duration) *ListRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list regions params
func (o *ListRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list regions params
func (o *ListRegionsParams) WithContext(ctx context.Context) *ListRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list regions params
func (o *ListRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list regions params
func (o *ListRegionsParams) WithHTTPClient(client *http.Client) *ListRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list regions params
func (o *ListRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
