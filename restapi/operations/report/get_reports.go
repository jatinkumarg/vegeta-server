// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReportsHandlerFunc turns a function with the right signature into a get reports handler
type GetReportsHandlerFunc func(GetReportsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReportsHandlerFunc) Handle(params GetReportsParams) middleware.Responder {
	return fn(params)
}

// GetReportsHandler interface for that can handle valid get reports params
type GetReportsHandler interface {
	Handle(GetReportsParams) middleware.Responder
}

// NewGetReports creates a new http.Handler for the get reports operation
func NewGetReports(ctx *middleware.Context, handler GetReportsHandler) *GetReports {
	return &GetReports{Context: ctx, Handler: handler}
}

/*GetReports swagger:route GET /report report getReports

Get a list of all reports

*/
type GetReports struct {
	Context *middleware.Context
	Handler GetReportsHandler
}

func (o *GetReports) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReportsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
