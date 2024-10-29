package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Method = string
type Pattern = string
type GinHandler = gin.HandlerFunc

const (
	Get     Method = "GET"
	Post    Method = "POST"
	Delete  Method = "DELETE"
	Put     Method = "PUT"
	Patch   Method = "PATCH"
	Options Method = "OPTIONS"
)

type Handler struct {
	method  Method
	pattern Pattern
	handler GinHandler
}

type HandlerParams struct {
	fx.In

	Method  Method
	Pattern Pattern
	Handler GinHandler
}

type HandlerResult struct {
	fx.Out

	Handler Handler
}

func NewHandler(params HandlerParams) (HandlerResult, error) {
	return HandlerResult{
		Handler: Handler{
			method:  params.Method,
			pattern: params.Pattern,
			handler: params.Handler,
		},
	}, nil
}

// IHandler
func (c Handler) Method() Method {
	return c.method
}

// to implement Route interface of fx
func (c Handler) Pattern() Pattern {
	return c.pattern
}

func (c Handler) Handler() GinHandler {
	return c.handler
}
