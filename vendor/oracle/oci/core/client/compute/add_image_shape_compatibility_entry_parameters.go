package compute

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

// NewAddImageShapeCompatibilityEntryParams creates a new AddImageShapeCompatibilityEntryParams object
// with the default values initialized.
func NewAddImageShapeCompatibilityEntryParams() *AddImageShapeCompatibilityEntryParams {
	var ()
	return &AddImageShapeCompatibilityEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddImageShapeCompatibilityEntryParamsWithTimeout creates a new AddImageShapeCompatibilityEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddImageShapeCompatibilityEntryParamsWithTimeout(timeout time.Duration) *AddImageShapeCompatibilityEntryParams {
	var ()
	return &AddImageShapeCompatibilityEntryParams{

		timeout: timeout,
	}
}

// NewAddImageShapeCompatibilityEntryParamsWithContext creates a new AddImageShapeCompatibilityEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddImageShapeCompatibilityEntryParamsWithContext(ctx context.Context) *AddImageShapeCompatibilityEntryParams {
	var ()
	return &AddImageShapeCompatibilityEntryParams{

		Context: ctx,
	}
}

// NewAddImageShapeCompatibilityEntryParamsWithHTTPClient creates a new AddImageShapeCompatibilityEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddImageShapeCompatibilityEntryParamsWithHTTPClient(client *http.Client) *AddImageShapeCompatibilityEntryParams {
	var ()
	return &AddImageShapeCompatibilityEntryParams{
		HTTPClient: client,
	}
}

/*AddImageShapeCompatibilityEntryParams contains all the parameters to send to the API endpoint
for the add image shape compatibility entry operation typically these are written to a http.Request
*/
type AddImageShapeCompatibilityEntryParams struct {

	/*ImageID
	  The OCID of the image.

	*/
	ImageID string
	/*ShapeName
	  Shape name.

	*/
	ShapeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) WithTimeout(timeout time.Duration) *AddImageShapeCompatibilityEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) WithContext(ctx context.Context) *AddImageShapeCompatibilityEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) WithHTTPClient(client *http.Client) *AddImageShapeCompatibilityEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) WithImageID(imageID string) *AddImageShapeCompatibilityEntryParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithShapeName adds the shapeName to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) WithShapeName(shapeName string) *AddImageShapeCompatibilityEntryParams {
	o.SetShapeName(shapeName)
	return o
}

// SetShapeName adds the shapeName to the add image shape compatibility entry params
func (o *AddImageShapeCompatibilityEntryParams) SetShapeName(shapeName string) {
	o.ShapeName = shapeName
}

// WriteToRequest writes these params to a swagger request
func (o *AddImageShapeCompatibilityEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
		return err
	}

	// path param shapeName
	if err := r.SetPathParam("shapeName", o.ShapeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
