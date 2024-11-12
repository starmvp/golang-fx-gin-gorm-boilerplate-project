package workflows

import (
	"boilerplate/internal/agents/tools"

	Ltools "github.com/starmvp/langchaingo/tools"
)

type Options struct {
	tools.Options

	ToolOptions []tools.Option // for the embeded tools.Options

	Tools                    []Ltools.Tool
	PromptPrefix             string
	PromptFormatInstructions string
	PromptSuffix             string
	PromptInputs             map[string]interface{}
}

func (options Options) GetToolOptions() tools.Options {
	return options.Options
}

type Option func(*Options)

func WithToolOptions(opts ...tools.Option) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(&o.Options)
		}
		o.ToolOptions = append(o.ToolOptions, opts...)
	}
}

func WithTools(ts ...Ltools.Tool) Option {
	return func(opts *Options) {
		opts.Tools = append(opts.Tools, ts...)
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
