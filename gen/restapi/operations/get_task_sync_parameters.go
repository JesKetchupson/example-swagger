// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTaskSyncParams creates a new GetTaskSyncParams object
// no default values defined in spec.
func NewGetTaskSyncParams() GetTaskSyncParams {

	return GetTaskSyncParams{}
}

// GetTaskSyncParams contains all the bound params for the get task sync operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTaskSync
type GetTaskSyncParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID Задачи
	  Required: true
	  In: path
	*/
	TaskID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTaskSyncParams() beforehand.
func (o *GetTaskSyncParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTaskID, rhkTaskID, _ := route.Params.GetOK("taskId")
	if err := o.bindTaskID(rTaskID, rhkTaskID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTaskID binds and validates parameter TaskID from path.
func (o *GetTaskSyncParams) bindTaskID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("taskId", "path", "strfmt.UUID", raw)
	}
	o.TaskID = *(value.(*strfmt.UUID))

	if err := o.validateTaskID(formats); err != nil {
		return err
	}

	return nil
}

// validateTaskID carries on validations for parameter TaskID
func (o *GetTaskSyncParams) validateTaskID(formats strfmt.Registry) error {

	if err := validate.FormatOf("taskId", "path", "uuid", o.TaskID.String(), formats); err != nil {
		return err
	}
	return nil
}
