package blockstorage

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

// NewUpdateVolumeParams creates a new UpdateVolumeParams object
// with the default values initialized.
func NewUpdateVolumeParams() *UpdateVolumeParams {
	var ()
	return &UpdateVolumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVolumeParamsWithTimeout creates a new UpdateVolumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVolumeParamsWithTimeout(timeout time.Duration) *UpdateVolumeParams {
	var ()
	return &UpdateVolumeParams{

		timeout: timeout,
	}
}

// NewUpdateVolumeParamsWithContext creates a new UpdateVolumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVolumeParamsWithContext(ctx context.Context) *UpdateVolumeParams {
	var ()
	return &UpdateVolumeParams{

		Context: ctx,
	}
}

// NewUpdateVolumeParamsWithHTTPClient creates a new UpdateVolumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVolumeParamsWithHTTPClient(client *http.Client) *UpdateVolumeParams {
	var ()
	return &UpdateVolumeParams{
		HTTPClient: client,
	}
}

/*UpdateVolumeParams contains all the parameters to send to the API endpoint
for the update volume operation typically these are written to a http.Request
*/
type UpdateVolumeParams struct {

	/*UpdateVolumeDetails
	  Update volume's display name.

	*/
	UpdateVolumeDetails *models.UpdateVolumeDetails
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*VolumeID
	  The OCID of the volume.

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update volume params
func (o *UpdateVolumeParams) WithTimeout(timeout time.Duration) *UpdateVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update volume params
func (o *UpdateVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update volume params
func (o *UpdateVolumeParams) WithContext(ctx context.Context) *UpdateVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update volume params
func (o *UpdateVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update volume params
func (o *UpdateVolumeParams) WithHTTPClient(client *http.Client) *UpdateVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update volume params
func (o *UpdateVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateVolumeDetails adds the updateVolumeDetails to the update volume params
func (o *UpdateVolumeParams) WithUpdateVolumeDetails(updateVolumeDetails *models.UpdateVolumeDetails) *UpdateVolumeParams {
	o.SetUpdateVolumeDetails(updateVolumeDetails)
	return o
}

// SetUpdateVolumeDetails adds the updateVolumeDetails to the update volume params
func (o *UpdateVolumeParams) SetUpdateVolumeDetails(updateVolumeDetails *models.UpdateVolumeDetails) {
	o.UpdateVolumeDetails = updateVolumeDetails
}

// WithIfMatch adds the ifMatch to the update volume params
func (o *UpdateVolumeParams) WithIfMatch(ifMatch *string) *UpdateVolumeParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update volume params
func (o *UpdateVolumeParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithVolumeID adds the volumeID to the update volume params
func (o *UpdateVolumeParams) WithVolumeID(volumeID string) *UpdateVolumeParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the update volume params
func (o *UpdateVolumeParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateVolumeDetails == nil {
		o.UpdateVolumeDetails = new(models.UpdateVolumeDetails)
	}

	if err := r.SetBodyParam(o.UpdateVolumeDetails); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param volumeId
	if err := r.SetPathParam("volumeId", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
