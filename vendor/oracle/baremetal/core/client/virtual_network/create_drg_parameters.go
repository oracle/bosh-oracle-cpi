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

	"oracle/baremetal/core/models"
)

// NewCreateDrgParams creates a new CreateDrgParams object
// with the default values initialized.
func NewCreateDrgParams() *CreateDrgParams {
	var ()
	return &CreateDrgParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDrgParamsWithTimeout creates a new CreateDrgParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDrgParamsWithTimeout(timeout time.Duration) *CreateDrgParams {
	var ()
	return &CreateDrgParams{

		timeout: timeout,
	}
}

// NewCreateDrgParamsWithContext creates a new CreateDrgParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDrgParamsWithContext(ctx context.Context) *CreateDrgParams {
	var ()
	return &CreateDrgParams{

		Context: ctx,
	}
}

// NewCreateDrgParamsWithHTTPClient creates a new CreateDrgParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDrgParamsWithHTTPClient(client *http.Client) *CreateDrgParams {
	var ()
	return &CreateDrgParams{
		HTTPClient: client,
	}
}

/*CreateDrgParams contains all the parameters to send to the API endpoint
for the create drg operation typically these are written to a http.Request
*/
type CreateDrgParams struct {

	/*CreateDrgDetails
	  Details for creating a DRG.

	*/
	CreateDrgDetails *models.CreateDrgDetails
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create drg params
func (o *CreateDrgParams) WithTimeout(timeout time.Duration) *CreateDrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create drg params
func (o *CreateDrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create drg params
func (o *CreateDrgParams) WithContext(ctx context.Context) *CreateDrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create drg params
func (o *CreateDrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create drg params
func (o *CreateDrgParams) WithHTTPClient(client *http.Client) *CreateDrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create drg params
func (o *CreateDrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateDrgDetails adds the createDrgDetails to the create drg params
func (o *CreateDrgParams) WithCreateDrgDetails(createDrgDetails *models.CreateDrgDetails) *CreateDrgParams {
	o.SetCreateDrgDetails(createDrgDetails)
	return o
}

// SetCreateDrgDetails adds the createDrgDetails to the create drg params
func (o *CreateDrgParams) SetCreateDrgDetails(createDrgDetails *models.CreateDrgDetails) {
	o.CreateDrgDetails = createDrgDetails
}

// WithOpcRetryToken adds the opcRetryToken to the create drg params
func (o *CreateDrgParams) WithOpcRetryToken(opcRetryToken *string) *CreateDrgParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the create drg params
func (o *CreateDrgParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateDrgDetails == nil {
		o.CreateDrgDetails = new(models.CreateDrgDetails)
	}

	if err := r.SetBodyParam(o.CreateDrgDetails); err != nil {
		return err
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
