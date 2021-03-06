// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/keptn/keptn/api/models"
)

// PostConfigureBridgeExposeHandlerFunc turns a function with the right signature into a post configure bridge expose handler
type PostConfigureBridgeExposeHandlerFunc func(PostConfigureBridgeExposeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostConfigureBridgeExposeHandlerFunc) Handle(params PostConfigureBridgeExposeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostConfigureBridgeExposeHandler interface for that can handle valid post configure bridge expose params
type PostConfigureBridgeExposeHandler interface {
	Handle(PostConfigureBridgeExposeParams, *models.Principal) middleware.Responder
}

// NewPostConfigureBridgeExpose creates a new http.Handler for the post configure bridge expose operation
func NewPostConfigureBridgeExpose(ctx *middleware.Context, handler PostConfigureBridgeExposeHandler) *PostConfigureBridgeExpose {
	return &PostConfigureBridgeExpose{Context: ctx, Handler: handler}
}

/*PostConfigureBridgeExpose swagger:route POST /configure/bridge/expose configure postConfigureBridgeExpose

Exposes the bridge

*/
type PostConfigureBridgeExpose struct {
	Context *middleware.Context
	Handler PostConfigureBridgeExposeHandler
}

func (o *PostConfigureBridgeExpose) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostConfigureBridgeExposeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
