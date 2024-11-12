package workflows

import (
	"boilerplate/internal/agents/tools"
	"boilerplate/internal/chain"
	"boilerplate/internal/utils"

	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/schema"
	Ltools "github.com/starmvp/langchaingo/tools"
)

type Options struct {
	tools.Options

	Tools                    []Ltools.Tool
	Memory                   *schema.Memory
	PromptPrefix             string
	PromptFormatInstructions string
	PromptSuffix             string
	PromptInputs             map[string]interface{}
}

func GetToolOptions(options Options) []tools.Option {
	return []tools.Option{
		tools.WithName(options.ToolName),
		tools.WithDescription(options.ToolDescription),
		tools.WithChain(options.Chain),
		tools.WithBuilder(options.Builder),
		tools.WithMemory(options.Memory),
		tools.WithVectorStore(options.VectorStore),
		tools.WithRetrieverNumDocuments(options.RetrieverNumDocuments),
		tools.WithRetriever(options.Retriever),
		tools.WithIO(options.IO),
		tools.WithCallbacksHandlers(options.CallbacksHandler),
	}
}

type Option func(*Options)

func WithTools(t []Ltools.Tool) Option {
	return func(opts *Options) {
		opts.Tools = t
	}
}

func WithTool(t Ltools.Tool) Option {
	return func(opts *Options) {
		opts.Tools = append(opts.Tools, t)
	}
}

func WithMemory(m *schema.Memory) Option {
	return func(opts *Options) {
		opts.Memory = m
	}
}

func WithPromptPrefix(prefix string) Option {
	return func(opts *Options) {
		opts.PromptPrefix = prefix
	}
}

func WithPromptFormatInstructions(instructions string) Option {
	return func(opts *Options) {
		opts.PromptFormatInstructions = instructions
	}
}

func WithPromptSuffix(suffix string) Option {
	return func(opts *Options) {
		opts.PromptSuffix = suffix
	}
}

func WithPromptInputs(inputs map[string]interface{}) Option {
	return func(opts *Options) {
		opts.PromptInputs = inputs
	}
}

func WithPromptInput(inputs map[string]interface{}) Option {
	return func(opts *Options) {
		for key, value := range inputs {
			opts.PromptInputs[key] = value
		}
	}
}

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
