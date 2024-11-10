package tools

import (
	"getidex_api/internal/chain"
	"getidex_api/internal/utils"

	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/schema"
)

type Options struct {
	ToolName        string
	ToolDescription string

	Chain   *chains.LLMChain
	Builder *chain.ChainBuilder
	Memory  *schema.Memory

	utils.IO

	CallbacksHandler []callbacks.Handler
}

type Option func(*Options)

func WithName(name string) Option {
	return func(o *Options) {
		o.ToolName = name
	}
}

func WithDescription(description string) Option {
	return func(o *Options) {
		o.ToolDescription = description
	}
}

func WithChain(c *chains.LLMChain) Option {
	return func(o *Options) {
		o.Chain = c
	}
}

func WithBuilder(b *chain.ChainBuilder) Option {
	return func(o *Options) {
		o.Builder = b
	}
}

func WithMemory(m *schema.Memory) Option {
	return func(o *Options) {
		o.Memory = m
	}
}

func WithIO(io utils.IO) Option {
	return func(o *Options) {
		o.IO = io
	}
}

func WithCallbacksHandler(h callbacks.Handler) Option {
	return func(o *Options) {
		o.CallbacksHandler = append(o.CallbacksHandler, h)
	}
}

func WithCallbacksHandlers(h []callbacks.Handler) Option {
	return func(o *Options) {
		o.CallbacksHandler = append(o.CallbacksHandler, h...)
	}
}
