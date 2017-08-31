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

// NewRemoveUserFromGroupParams creates a new RemoveUserFromGroupParams object
// with the default values initialized.
func NewRemoveUserFromGroupParams() *RemoveUserFromGroupParams {
	var ()
	return &RemoveUserFromGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveUserFromGroupParamsWithTimeout creates a new RemoveUserFromGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveUserFromGroupParamsWithTimeout(timeout time.Duration) *RemoveUserFromGroupParams {
	var ()
	return &RemoveUserFromGroupParams{

		timeout: timeout,
	}
}

// NewRemoveUserFromGroupParamsWithContext creates a new RemoveUserFromGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveUserFromGroupParamsWithContext(ctx context.Context) *RemoveUserFromGroupParams {
	var ()
	return &RemoveUserFromGroupParams{

		Context: ctx,
	}
}

// NewRemoveUserFromGroupParamsWithHTTPClient creates a new RemoveUserFromGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveUserFromGroupParamsWithHTTPClient(client *http.Client) *RemoveUserFromGroupParams {
	var ()
	return &RemoveUserFromGroupParams{
		HTTPClient: client,
	}
}

/*RemoveUserFromGroupParams contains all the parameters to send to the API endpoint
for the remove user from group operation typically these are written to a http.Request
*/
type RemoveUserFromGroupParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*UserGroupMembershipID
	  The OCID of the userGroupMembership.

	*/
	UserGroupMembershipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove user from group params
func (o *RemoveUserFromGroupParams) WithTimeout(timeout time.Duration) *RemoveUserFromGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove user from group params
func (o *RemoveUserFromGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove user from group params
func (o *RemoveUserFromGroupParams) WithContext(ctx context.Context) *RemoveUserFromGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove user from group params
func (o *RemoveUserFromGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove user from group params
func (o *RemoveUserFromGroupParams) WithHTTPClient(client *http.Client) *RemoveUserFromGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove user from group params
func (o *RemoveUserFromGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the remove user from group params
func (o *RemoveUserFromGroupParams) WithIfMatch(ifMatch *string) *RemoveUserFromGroupParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the remove user from group params
func (o *RemoveUserFromGroupParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithUserGroupMembershipID adds the userGroupMembershipID to the remove user from group params
func (o *RemoveUserFromGroupParams) WithUserGroupMembershipID(userGroupMembershipID string) *RemoveUserFromGroupParams {
	o.SetUserGroupMembershipID(userGroupMembershipID)
	return o
}

// SetUserGroupMembershipID adds the userGroupMembershipId to the remove user from group params
func (o *RemoveUserFromGroupParams) SetUserGroupMembershipID(userGroupMembershipID string) {
	o.UserGroupMembershipID = userGroupMembershipID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveUserFromGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userGroupMembershipId
	if err := r.SetPathParam("userGroupMembershipId", o.UserGroupMembershipID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
