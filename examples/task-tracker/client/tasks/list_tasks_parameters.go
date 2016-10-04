package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListTasksParams creates a new ListTasksParams object
// with the default values initialized.
func NewListTasksParams() *ListTasksParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &ListTasksParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksParamsWithTimeout creates a new ListTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTasksParamsWithTimeout(timeout time.Duration) *ListTasksParams {
	var (
		pageSizeDefault int32 = int32(20)
	)
	return &ListTasksParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListTasksParamsWithContext creates a new ListTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTasksParamsWithContext(ctx context.Context) *ListTasksParams {
	var (
		pageSizeDefault int32 = int32(20)
	)
	return &ListTasksParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

/*ListTasksParams contains all the parameters to send to the API endpoint
for the list tasks operation typically these are written to a http.Request
*/
type ListTasksParams struct {

	/*PageSize
	  Amount of items to return in a single page

	*/
	PageSize *int32
	/*SinceID
	  The last id that was seen.

	*/
	SinceID *int64
	/*Status
	  the status to filter by

	*/
	Status []string
	/*Tags
	  the tags to filter by

	*/
	Tags []string

	timeout time.Duration
	Context context.Context
}

func (o *ListTasksParams) WithTimeout(timeout time.Duration) *ListTasksParams {
	o.SetTimeout(timeout)
	return o
}

func (o *ListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

func (o *ListTasksParams) WithContext(ctx context.Context) *ListTasksParams {
	o.SetContext(ctx)
	return o
}

func (o *ListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithPageSize adds the pageSize to the list tasks params
func (o *ListTasksParams) WithPageSize(pageSize *int32) *ListTasksParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list tasks params
func (o *ListTasksParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSinceID adds the sinceID to the list tasks params
func (o *ListTasksParams) WithSinceID(sinceID *int64) *ListTasksParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the list tasks params
func (o *ListTasksParams) SetSinceID(sinceID *int64) {
	o.SinceID = sinceID
}

// WithStatus adds the status to the list tasks params
func (o *ListTasksParams) WithStatus(status []string) *ListTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list tasks params
func (o *ListTasksParams) SetStatus(status []string) {
	o.Status = status
}

// WithTags adds the tags to the list tasks params
func (o *ListTasksParams) WithTags(tags []string) *ListTasksParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the list tasks params
func (o *ListTasksParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.SinceID != nil {

		// query param sinceId
		var qrSinceID int64
		if o.SinceID != nil {
			qrSinceID = *o.SinceID
		}
		qSinceID := swag.FormatInt64(qrSinceID)
		if qSinceID != "" {
			if err := r.SetQueryParam("sinceId", qSinceID); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "pipes")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
