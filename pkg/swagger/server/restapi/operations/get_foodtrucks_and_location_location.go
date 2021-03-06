// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFoodtrucksAndLocationLocationHandlerFunc turns a function with the right signature into a get foodtrucks and location location handler
type GetFoodtrucksAndLocationLocationHandlerFunc func(GetFoodtrucksAndLocationLocationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFoodtrucksAndLocationLocationHandlerFunc) Handle(params GetFoodtrucksAndLocationLocationParams) middleware.Responder {
	return fn(params)
}

// GetFoodtrucksAndLocationLocationHandler interface for that can handle valid get foodtrucks and location location params
type GetFoodtrucksAndLocationLocationHandler interface {
	Handle(GetFoodtrucksAndLocationLocationParams) middleware.Responder
}

// NewGetFoodtrucksAndLocationLocation creates a new http.Handler for the get foodtrucks and location location operation
func NewGetFoodtrucksAndLocationLocation(ctx *middleware.Context, handler GetFoodtrucksAndLocationLocationHandler) *GetFoodtrucksAndLocationLocation {
	return &GetFoodtrucksAndLocationLocation{Context: ctx, Handler: handler}
}

/* GetFoodtrucksAndLocationLocation swagger:route GET /foodtrucks&location={location} getFoodtrucksAndLocationLocation

Returns a list of food Trucks.

Optional extended description in Markdown.

*/
type GetFoodtrucksAndLocationLocation struct {
	Context *middleware.Context
	Handler GetFoodtrucksAndLocationLocationHandler
}

func (o *GetFoodtrucksAndLocationLocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFoodtrucksAndLocationLocationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
