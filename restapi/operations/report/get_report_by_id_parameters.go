// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReportByIDParams creates a new GetReportByIDParams object
// no default values defined in spec.
func NewGetReportByIDParams() GetReportByIDParams {

	return GetReportByIDParams{}
}

// GetReportByIDParams contains all the bound params for the get report by ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters getReportByID
type GetReportByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attack ID
	  Required: true
	  In: path
	*/
	AttackID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetReportByIDParams() beforehand.
func (o *GetReportByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAttackID, rhkAttackID, _ := route.Params.GetOK("attackID")
	if err := o.bindAttackID(rAttackID, rhkAttackID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAttackID binds and validates parameter AttackID from path.
func (o *GetReportByIDParams) bindAttackID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AttackID = raw

	return nil
}
