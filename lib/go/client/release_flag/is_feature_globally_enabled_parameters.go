// Code generated by go-swagger; DO NOT EDIT.

package release_flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/toggler-io/toggler/lib/go/models"
)

// NewIsFeatureGloballyEnabledParams creates a new IsFeatureGloballyEnabledParams object
// with the default values initialized.
func NewIsFeatureGloballyEnabledParams() *IsFeatureGloballyEnabledParams {
	var ()
	return &IsFeatureGloballyEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIsFeatureGloballyEnabledParamsWithTimeout creates a new IsFeatureGloballyEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIsFeatureGloballyEnabledParamsWithTimeout(timeout time.Duration) *IsFeatureGloballyEnabledParams {
	var ()
	return &IsFeatureGloballyEnabledParams{

		timeout: timeout,
	}
}

// NewIsFeatureGloballyEnabledParamsWithContext creates a new IsFeatureGloballyEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewIsFeatureGloballyEnabledParamsWithContext(ctx context.Context) *IsFeatureGloballyEnabledParams {
	var ()
	return &IsFeatureGloballyEnabledParams{

		Context: ctx,
	}
}

// NewIsFeatureGloballyEnabledParamsWithHTTPClient creates a new IsFeatureGloballyEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIsFeatureGloballyEnabledParamsWithHTTPClient(client *http.Client) *IsFeatureGloballyEnabledParams {
	var ()
	return &IsFeatureGloballyEnabledParams{
		HTTPClient: client,
	}
}

/*IsFeatureGloballyEnabledParams contains all the parameters to send to the API endpoint
for the is feature globally enabled operation typically these are written to a http.Request
*/
type IsFeatureGloballyEnabledParams struct {

	/*Body*/
	Body *models.IsFeatureGloballyEnabledRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) WithTimeout(timeout time.Duration) *IsFeatureGloballyEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) WithContext(ctx context.Context) *IsFeatureGloballyEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) WithHTTPClient(client *http.Client) *IsFeatureGloballyEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) WithBody(body *models.IsFeatureGloballyEnabledRequestBody) *IsFeatureGloballyEnabledParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the is feature globally enabled params
func (o *IsFeatureGloballyEnabledParams) SetBody(body *models.IsFeatureGloballyEnabledRequestBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IsFeatureGloballyEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
