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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListUsersParams creates a new ListUsersParams object
// with the default values initialized.
func NewListUsersParams() *ListUsersParams {
	var ()
	return &ListUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUsersParamsWithTimeout creates a new ListUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUsersParamsWithTimeout(timeout time.Duration) *ListUsersParams {
	var ()
	return &ListUsersParams{

		timeout: timeout,
	}
}

// NewListUsersParamsWithContext creates a new ListUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUsersParamsWithContext(ctx context.Context) *ListUsersParams {
	var ()
	return &ListUsersParams{

		Context: ctx,
	}
}

// NewListUsersParamsWithHTTPClient creates a new ListUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUsersParamsWithHTTPClient(client *http.Client) *ListUsersParams {
	var ()
	return &ListUsersParams{
		HTTPClient: client,
	}
}

/*ListUsersParams contains all the parameters to send to the API endpoint
for the list users operation typically these are written to a http.Request
*/
type ListUsersParams struct {

	/*CompartmentID
	  The OCID of the compartment (remember that the tenancy is simply the root compartment).


	*/
	CompartmentID string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list users params
func (o *ListUsersParams) WithTimeout(timeout time.Duration) *ListUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list users params
func (o *ListUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list users params
func (o *ListUsersParams) WithContext(ctx context.Context) *ListUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list users params
func (o *ListUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list users params
func (o *ListUsersParams) WithHTTPClient(client *http.Client) *ListUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list users params
func (o *ListUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list users params
func (o *ListUsersParams) WithCompartmentID(compartmentID string) *ListUsersParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list users params
func (o *ListUsersParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithLimit adds the limit to the list users params
func (o *ListUsersParams) WithLimit(limit *int64) *ListUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list users params
func (o *ListUsersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list users params
func (o *ListUsersParams) WithPage(page *string) *ListUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list users params
func (o *ListUsersParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
