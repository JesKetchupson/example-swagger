// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTaskSyncHandlerFunc turns a function with the right signature into a get task sync handler
type GetTaskSyncHandlerFunc func(GetTaskSyncParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskSyncHandlerFunc) Handle(params GetTaskSyncParams) middleware.Responder {
	return fn(params)
}

// GetTaskSyncHandler interface for that can handle valid get task sync params
type GetTaskSyncHandler interface {
	Handle(GetTaskSyncParams) middleware.Responder
}

// NewGetTaskSync creates a new http.Handler for the get task sync operation
func NewGetTaskSync(ctx *middleware.Context, handler GetTaskSyncHandler) *GetTaskSync {
	return &GetTaskSync{Context: ctx, Handler: handler}
}

/*GetTaskSync swagger:route GET /task/{taskId} getTaskSync

getTaskSync

1. Ищет в БД текущее состояние задачи.
2. Возвращает текущее состояние задачи

*/
type GetTaskSync struct {
	Context *middleware.Context
	Handler GetTaskSyncHandler
}

func (o *GetTaskSync) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTaskSyncParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}