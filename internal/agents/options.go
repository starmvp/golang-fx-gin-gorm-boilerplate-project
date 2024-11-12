package agents

import (
	"context"

	"boilerplate/internal/agents/workflows"
	"boilerplate/internal/chain"
	"boilerplate/internal/conversation"
	"boilerplate/internal/utils"

	Lagents "github.com/starmvp/langchaingo/agents"
	Lcallbacks "github.com/starmvp/langchaingo/callbacks"
	Lchains "github.com/starmvp/langchaingo/chains"
	Lllms "github.com/starmvp/langchaingo/llms"
	Lprompts "github.com/starmvp/langchaingo/prompts"
	Lschema "github.com/starmvp/langchaingo/schema"
	Ltools "github.com/starmvp/langchaingo/tools"
)

type Options struct {
	LangChainOptions []Lagents.Option

	LLM Lllms.Model

	LangChainAgent   Lagents.Agent
	UseStreamingMode bool

	// LangChainAgent options
	PromptTemplate           Lprompts.PromptTemplate
	PromptPrefix             string
	PromptSuffix             string
	PromptFormatInstructions string
	OutputKey                string

	Ctx               context.Context
	Builder           *chain.ChainBuilder
	Chains            []Lchains.Chain
	Tools             []Ltools.Tool
	Memory            Lschema.Memory
	CallbacksHandlers []Lcallbacks.Handler
	Conversation      conversation.Conversation

	utils.IO

	// ConversationalWorkflowAgent
	Workflows []workflows.Workflow
}

type Option func(*Options)

func WithLangChainOptions(os ...Lagents.Option) Option {
	return func(opts *Options) {
		opts.LangChainOptions = append(opts.LangChainOptions, os...)
	}
}

func WithLLM(llm Lllms.Model) Option {
	return func(o *Options) {
		o.LLM = llm
	}
}

func WithLangChainAgent(a Lagents.Agent) Option {
	return func(o *Options) {
		o.LangChainAgent = a
	}
}

// for (starmvp/langchaingo/agents/conversational.go)ConversationalAgent
func WithStreamingMode(use bool) Option {
	return func(o *Options) {
		o.UseStreamingMode = use
	}
}

func WithPromptTemplate(t Lprompts.PromptTemplate) Option {
	return func(o *Options) {
		o.PromptTemplate = t
		o.LangChainOptions = append(o.LangChainOptions, Lagents.WithPrompt(t))
	}
}

func WithPromptPrefix(prefix string) Option {
	return func(o *Options) {
		o.PromptPrefix = prefix
		o.LangChainOptions = append(o.LangChainOptions, Lagents.WithPromptPrefix(prefix))
	}
}

func WithFormatInstructions(instructions string) Option {
	return func(o *Options) {
		o.PromptFormatInstructions = instructions
		o.LangChainOptions = append(o.LangChainOptions, Lagents.WithPromptFormatInstructions(instructions))
	}
}

func WithPromptSuffix(suffix string) Option {
	return func(o *Options) {
		o.PromptSuffix = suffix
		o.LangChainOptions = append(o.LangChainOptions, Lagents.WithPromptSuffix(suffix))
	}
}

func WithOutputKey(key string) Option {
	return func(o *Options) {
		o.OutputKey = key
		o.LangChainOptions = append(o.LangChainOptions, Lagents.WithOutputKey(key))
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Ctx = ctx
	}
}

func WithChainBuilder(b *chain.ChainBuilder) Option {
	return func(o *Options) {
		o.Builder = b
	}
}

func WithChains(c []Lchains.Chain) Option {
	return func(o *Options) {
		o.Chains = append(o.Chains, c...)
	}
}

func WithChain(c Lchains.Chain) Option {
	return func(o *Options) {
		o.Chains = append(o.Chains, c)
	}
}

func WithTools(t []Ltools.Tool) Option {
	return func(o *Options) {
		o.Tools = append(o.Tools, t...)
	}
}

func WithTool(t Ltools.Tool) Option {
	return func(o *Options) {
		o.Tools = append(o.Tools, t)
	}
}

func WithMemory(m Lschema.Memory) Option {
	return func(o *Options) {
		o.Memory = m
	}
}

func WithCallbacksHandlers(hs ...Lcallbacks.Handler) Option {
	return func(o *Options) {
		o.CallbacksHandlers = append(o.CallbacksHandlers, hs...)
	}
}

func WithConversation(c conversation.Conversation) Option {
	return func(o *Options) {
		o.Conversation = c
	}
}

func WithIO(io utils.IO) Option {
	return func(o *Options) {
		o.IO = io
	}
}

func WithWorkflows(w []workflows.Workflow) Option {
	return func(o *Options) {
		o.Workflows = w
	}
}
