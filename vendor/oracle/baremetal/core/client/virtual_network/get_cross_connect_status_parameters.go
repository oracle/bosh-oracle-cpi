package virtual_network

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

// NewGetCrossConnectStatusParams creates a new GetCrossConnectStatusParams object
// with the default values initialized.
func NewGetCrossConnectStatusParams() *GetCrossConnectStatusParams {
	var ()
	return &GetCrossConnectStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCrossConnectStatusParamsWithTimeout creates a new GetCrossConnectStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCrossConnectStatusParamsWithTimeout(timeout time.Duration) *GetCrossConnectStatusParams {
	var ()
	return &GetCrossConnectStatusParams{

		timeout: timeout,
	}
}

// NewGetCrossConnectStatusParamsWithContext creates a new GetCrossConnectStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCrossConnectStatusParamsWithContext(ctx context.Context) *GetCrossConnectStatusParams {
	var ()
	return &GetCrossConnectStatusParams{

		Context: ctx,
	}
}

// NewGetCrossConnectStatusParamsWithHTTPClient creates a new GetCrossConnectStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCrossConnectStatusParamsWithHTTPClient(client *http.Client) *GetCrossConnectStatusParams {
	var ()
	return &GetCrossConnectStatusParams{
		HTTPClient: client,
	}
}

/*GetCrossConnectStatusParams contains all the parameters to send to the API endpoint
for the get cross connect status operation typically these are written to a http.Request
*/
type GetCrossConnectStatusParams struct {

	/*CrossConnectID
	  The OCID of the cross-connect.

	*/
	CrossConnectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cross connect status params
func (o *GetCrossConnectStatusParams) WithTimeout(timeout time.Duration) *GetCrossConnectStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cross connect status params
func (o *GetCrossConnectStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cross connect status params
func (o *GetCrossConnectStatusParams) WithContext(ctx context.Context) *GetCrossConnectStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cross connect status params
func (o *GetCrossConnectStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cross connect status params
func (o *GetCrossConnectStatusParams) WithHTTPClient(client *http.Client) *GetCrossConnectStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cross connect status params
func (o *GetCrossConnectStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrossConnectID adds the crossConnectID to the get cross connect status params
func (o *GetCrossConnectStatusParams) WithCrossConnectID(crossConnectID string) *GetCrossConnectStatusParams {
	o.SetCrossConnectID(crossConnectID)
	return o
}

// SetCrossConnectID adds the crossConnectId to the get cross connect status params
func (o *GetCrossConnectStatusParams) SetCrossConnectID(crossConnectID string) {
	o.CrossConnectID = crossConnectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCrossConnectStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crossConnectId
	if err := r.SetPathParam("crossConnectId", o.CrossConnectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
