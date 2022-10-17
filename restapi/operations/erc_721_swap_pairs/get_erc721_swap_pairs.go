// Code generated by go-swagger; DO NOT EDIT.

package erc_721_swap_pairs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetErc721SwapPairsHandlerFunc turns a function with the right signature into a get erc721 swap pairs handler
type GetErc721SwapPairsHandlerFunc func(GetErc721SwapPairsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetErc721SwapPairsHandlerFunc) Handle(params GetErc721SwapPairsParams) middleware.Responder {
	return fn(params)
}

// GetErc721SwapPairsHandler interface for that can handle valid get erc721 swap pairs params
type GetErc721SwapPairsHandler interface {
	Handle(GetErc721SwapPairsParams) middleware.Responder
}

// NewGetErc721SwapPairs creates a new http.Handler for the get erc721 swap pairs operation
func NewGetErc721SwapPairs(ctx *middleware.Context, handler GetErc721SwapPairsHandler) *GetErc721SwapPairs {
	return &GetErc721SwapPairs{Context: ctx, Handler: handler}
}

/* GetErc721SwapPairs swagger:route GET /v1/erc-721-swap-pairs erc_721_swap_pairs getErc721SwapPairs

Gets a list of available ERC721 swap pairs.

*/
type GetErc721SwapPairs struct {
	Context *middleware.Context
	Handler GetErc721SwapPairsHandler
}

func (o *GetErc721SwapPairs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetErc721SwapPairsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
