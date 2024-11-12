package tools

import (
	"boilerplate/internal/chain"
	"boilerplate/internal/utils"

	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/schema"
	"github.com/starmvp/langchaingo/vectorstores"
)

type Options struct {
	ToolName        string
	ToolDescription string

	Chain   *chains.LLMChain
	Builder *chain.ChainBuilder
	Memory  schema.Memory

	VectorStore           vectorstores.VectorStore
	RetrieverNumDocuments int
	Retriever             schema.Retriever

	utils.IO

	CallbacksHandlers []callbacks.Handler
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

func WithMemory(m schema.Memory) Option {
	return func(o *Options) {
		o.Memory = m
	}
}

func WithIO(io utils.IO) Option {
	return func(o *Options) {
		o.IO = io
	}
}

func WithVectorStore(v vectorstores.VectorStore) Option {
	return func(o *Options) {
		o.VectorStore = v
	}
}

func WithRetrieverNumDocuments(n int) Option {
	return func(o *Options) {
		o.RetrieverNumDocuments = n
	}
}

func WithRetriever(r schema.Retriever) Option {
	return func(o *Options) {
		o.Retriever = r
	}
}

func WithCallbacksHandlers(hs ...callbacks.Handler) Option {
	return func(o *Options) {
		o.CallbacksHandlers = append(o.CallbacksHandlers, hs...)
	}
}
